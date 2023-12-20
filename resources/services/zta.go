package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/crowdstrike/gofalcon/falcon"
	"github.com/crowdstrike/gofalcon/falcon/client/zero_trust_assessment"
	"github.com/crowdstrike/gofalcon/falcon/models"
	"github.com/justmiles/cq-source-crowdstrike/client"
)

func ZTAs() *schema.Table {
	return &schema.Table{
		Name:      "crowdstrike_falcon_zta",
		Resolver:  fetchZTAs,
		Transform: transformers.TransformWithStruct(&models.DomainSignalProperties{}, transformers.WithPrimaryKeys("Aid")),
	}
}

func fetchZTAs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {

	hostsChan := make(chan any)
	defer close(hostsChan)

	errChan := make(chan error)
	defer close(errChan)
	go func() {
		errChan <- fetchHosts(ctx, meta, parent, hostsChan)
	}()

	c := meta.(*client.Client)
	for batch := range batchedHosts(hostsChan, 100, 5*time.Second) {
		result, err := resolveZTAs(ctx, c, batch)
		if err != nil {
			return err
		}
		for _, zta := range result.Resources {
			res <- zta
		}
	}

	return <-errChan
}

func batchedHosts(hosts <-chan any, maxBatchSize int, maxTimeout time.Duration) chan []string {
	batches := make(chan []string)

	go func() {
		defer close(batches)

		for keepGoing := true; keepGoing; {
			var batch []string
			expire := time.After(maxTimeout)
			for {
				select {
				case host, ok := <-hosts:
					if !ok {
						keepGoing = false
						goto done
					}
					batch = append(batch, *host.(*models.DeviceapiDeviceSwagger).DeviceID)
					if len(batch) == maxBatchSize {
						goto done
					}
				case <-expire:
					return
				}
			}
		done:
			if len(batch) > 0 {
				batches <- batch
			}
		}
	}()

	return batches
}

func resolveZTAs(ctx context.Context, c *client.Client, hostIds []string) (*models.DomainAssessmentsResponse, error) {
	ztaOK, err := c.CrowdStrike.ZeroTrustAssessment.GetAssessmentV1(&zero_trust_assessment.GetAssessmentV1Params{
		Ids:     hostIds,
		Context: ctx,
	})

	switch t := err.(type) {
	case *zero_trust_assessment.GetAssessmentV1NotFound:
		c.Logger.Warn().Msgf("No ZTA found = %s", t.Payload.Errors)
		foundResults := &models.DomainAssessmentsResponse{}
		rawData, err := falcon.ErrorExtractPayload(err).MarshalBinary()
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(rawData, foundResults); err != nil {
			return nil, err
		}
		return foundResults, nil
	case nil:
		return ztaOK.GetPayload(), nil
	default:
		return nil, t
	}
}
