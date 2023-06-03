package client

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/crowdstrike/gofalcon/falcon"
	"github.com/crowdstrike/gofalcon/falcon/client"
	"github.com/rs/zerolog"
)

type Client struct {
	Logger zerolog.Logger

	CrowdStrike *client.CrowdStrikeAPISpecification
}

func (c *Client) ID() string {
	return "CrowdStrike"
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source, opts source.Options) (schema.ClientMeta, error) {
	var pluginSpec Spec

	if err := s.UnmarshalSpec(&pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal plugin spec: %w", err)
	}

	falconClientID := os.Getenv("FALCON_CLIENT_ID")
	falconSecret := os.Getenv("FALCON_SECRET")

	c, err := falcon.NewClient(&falcon.ApiConfig{
		ClientId:     falconClientID,
		ClientSecret: falconSecret,
		Context:      context.Background(),
	})
	if err != nil {
		return nil, fmt.Errorf("could not auth: %s", err.Error())
	}

	return &Client{
		Logger:      logger,
		CrowdStrike: c,
	}, nil
}
