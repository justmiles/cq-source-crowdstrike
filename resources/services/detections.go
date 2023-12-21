package services

import (
	"context"
	"fmt"

	"github.com/justmiles/cq-source-crowdstrike/client"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/crowdstrike/gofalcon/falcon/client/detects"
	"github.com/crowdstrike/gofalcon/falcon/models"
)

func Detections() *schema.Table {
	return &schema.Table{
		Name:      "crowdstrike_falcon_detections",
		Resolver:  fetchDetections,
		Transform: transformers.TransformWithStruct(&models.DomainAPIDetectionDocument{}, transformers.WithPrimaryKeys("DetectionID")),
	}
}

func fetchDetections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	queryOK, err := c.CrowdStrike.Detects.QueryDetects(&detects.QueryDetectsParams{
		Context: ctx,
	})
	if err != nil {
		return fmt.Errorf("could not query detects: %s", err.Error())
	}
	queryResponse := queryOK.GetPayload()

	detectionOK, err := c.CrowdStrike.Detects.GetDetectSummaries(&detects.GetDetectSummariesParams{
		Context: ctx,
		Body: &models.MsaIdsRequest{
			Ids: queryResponse.Resources,
		},
	})
	if err != nil {
		return fmt.Errorf("could not get detects: %s", err.Error())
	}

	detectionResponse := detectionOK.GetPayload()

	for _, detect := range detectionResponse.Resources {
		res <- detect
	}

	return nil
}
