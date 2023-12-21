package services

import (
	"context"
	"encoding/json"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/crowdstrike/gofalcon/falcon"
	"github.com/crowdstrike/gofalcon/falcon/client/hosts"
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
	c := meta.(*client.Client)

	var offset int64 = 0
	var limit int64 = 100

	for {
		queryOk, err := c.CrowdStrike.Hosts.QueryDevicesByFilter(&hosts.QueryDevicesByFilterParams{
			Context: ctx,
			Limit:   &limit,
			Offset:  &offset,
		})
		if err != nil {
			return err
		}

		queryResponse := queryOk.GetPayload()
		if err := falcon.AssertNoError(queryResponse.Errors); err != nil {
			return err
		}

		ztas, err := resolveZTAs(ctx, c, queryResponse.Resources)
		if err != nil {
			return err
		}
		for _, zta := range ztas.Resources {
			res <- zta
		}

		offset += limit

		if *queryResponse.Meta.Pagination.Total == int64(*queryResponse.Meta.Pagination.Offset) {
			break
		}
	}

	return nil
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
