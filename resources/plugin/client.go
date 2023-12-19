package plugin

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/cloudquery/plugin-sdk/v4/docs"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/crowdstrike/gofalcon/falcon"
	"github.com/justmiles/cq-source-crowdstrike/client"
	"github.com/justmiles/cq-source-crowdstrike/resources/services"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	maxMsgSize = 100 * 1024 * 1024 // 100 MiB
)

type Client struct {
	logger      zerolog.Logger
	config      client.Spec
	tables      schema.Tables
	options     plugin.NewClientOptions
	scheduler   *scheduler.Scheduler
	backendConn *grpc.ClientConn

	plugin.UnimplementedDestination
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) Sync(ctx context.Context, options plugin.SyncOptions, res chan<- message.SyncMessage) error {
	tt, err := c.tables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return err
	}

	var stateClient state.Client
	if options.BackendOptions == nil {
		c.logger.Info().Msg("No backend options provided, using no state backend")
		stateClient = &state.NoOpClient{}
		c.backendConn = nil
	} else {
		c.backendConn, err = grpc.DialContext(ctx, options.BackendOptions.Connection,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultCallOptions(
				grpc.MaxCallRecvMsgSize(maxMsgSize),
				grpc.MaxCallSendMsgSize(maxMsgSize),
			),
		)
		if err != nil {
			return fmt.Errorf("failed to dial grpc source plugin at %s: %w", options.BackendOptions.Connection, err)
		}
		stateClient, err = state.NewClient(ctx, c.backendConn, options.BackendOptions.TableName)
		if err != nil {
			return fmt.Errorf("failed to create state client: %w", err)
		}
		c.logger.Info().Str("table_name", options.BackendOptions.TableName).Msg("Connected to state backend")
	}

	falconClientID := os.Getenv("FALCON_CLIENT_ID")
	falconSecret := os.Getenv("FALCON_SECRET")

	fc, err := falcon.NewClient(&falcon.ApiConfig{
		ClientId:     falconClientID,
		ClientSecret: falconSecret,
		Debug:        false,
		Context:      ctx,
	})
	if err != nil {
		return fmt.Errorf("could not auth: %w", err)
	}

	schedulerClient := client.New(c.logger, c.config, fc, stateClient)
	err = c.scheduler.Sync(ctx, schedulerClient, tt, res, scheduler.WithSyncDeterministicCQID(options.DeterministicCQID))
	if err != nil {
		return fmt.Errorf("failed to sync: %w", err)
	}
	return stateClient.Flush(ctx)
}

func (c *Client) Tables(_ context.Context, options plugin.TableOptions) (schema.Tables, error) {
	tt, err := c.tables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return nil, err
	}
	return tt, nil
}

func (c *Client) Close(_ context.Context) error {
	if c.backendConn != nil {
		return c.backendConn.Close()
	}
	return nil
}

func Configure(_ context.Context, logger zerolog.Logger, specBytes []byte, opts plugin.NewClientOptions) (plugin.Client, error) {
	if opts.NoConnection {
		return &Client{
			logger:  logger.With().Str("module", "crowdstrike").Logger(),
			options: opts,
			tables:  getTables(),
		}, nil
	}

	config := client.Spec{}
	if err := json.Unmarshal(specBytes, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	config.SetDefaults()
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate spec: %w", err)
	}

	return &Client{
		logger:  logger.With().Str("module", "crowdstrike").Logger(),
		options: opts,
		config:  config,
		scheduler: scheduler.NewScheduler(
			scheduler.WithLogger(logger),
			scheduler.WithConcurrency(config.Concurrency),
		),
		tables: getTables(),
	}, nil
}

func getTables() schema.Tables {
	tables := []*schema.Table{
		services.Incidents(),
		services.Detections(),
		services.Hosts(),
		services.Vulnerabilities(),
		services.DiscoverApps(),
		services.DiscoverHosts(),
		services.ZTAs(),
	}
	if err := transformers.TransformTables(tables); err != nil {
		panic(err)
	}
	if err := transformers.Apply(tables, func(t *schema.Table) error {
		t.Title = docs.DefaultTitleTransformer(t)
		return nil
	}); err != nil {
		panic(err)
	}
	for _, t := range tables {
		schema.AddCqIDs(t)
	}
	return tables
}
