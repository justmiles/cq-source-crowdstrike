package resources

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"

	"github.com/crowdstrike/gofalcon/falcon"
	"github.com/crowdstrike/gofalcon/falcon/client/detects"
	"github.com/crowdstrike/gofalcon/falcon/models"
)

func Detections() *schema.Table {
	return &schema.Table{
		Name:      "crowdstrike_falcon_detections",
		Resolver:  fetchDetections,
		Transform: transformers.TransformWithStruct(&models.DomainAPIDetectionDocument{}),
	}
}

func fetchDetections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	// client := meta.(*Client)

	falconClientID := os.Getenv("FALCON_CLIENT_ID")
	falconSecret := os.Getenv("FALCON_SECRET")

	client, err := falcon.NewClient(&falcon.ApiConfig{
		ClientId:     falconClientID,
		ClientSecret: falconSecret,
		Context:      context.Background(),
	})
	if err != nil {
		return fmt.Errorf("could not auth: %s", err.Error())
	}

	queryOK, err := client.Detects.QueryDetects(&detects.QueryDetectsParams{
		Context: context.Background(),
	})
	if err != nil {
		return fmt.Errorf("could not get incident: %s", err.Error())
	}
	queryResponse := queryOK.GetPayload()

	detectionOK, err := client.Detects.GetDetectSummaries(&detects.GetDetectSummariesParams{
		Context: context.Background(),
		Body: &models.MsaIdsRequest{
			Ids: queryResponse.Resources,
		},
	})

	detectionResponse := detectionOK.GetPayload()

	for _, detect := range detectionResponse.Resources {
		res <- detect
	}

	return nil
}
