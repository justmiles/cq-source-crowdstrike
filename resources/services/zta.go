package services

import (
	"context"
	"fmt"

	"github.com/justmiles/cq-source-crowdstrike/client"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/crowdstrike/gofalcon/falcon"
	"github.com/crowdstrike/gofalcon/falcon/client/hosts"
	"github.com/crowdstrike/gofalcon/falcon/client/zero_trust_assessment"
	"github.com/crowdstrike/gofalcon/falcon/models"
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

		queryRespOK, err := c.CrowdStrike.Hosts.QueryDevicesByFilter(&hosts.QueryDevicesByFilterParams{
			Context: ctx,
			Limit:   &limit,
			Offset:  &offset,
		})
		if err != nil {
			return fmt.Errorf("could not query hosts: %s", err.Error())
		}

		queryResponse := queryRespOK.GetPayload()
		if err = falcon.AssertNoError(queryResponse.Errors); err != nil {
			return fmt.Errorf("Error querying hosts: %s", falcon.ErrorExplain(err))
		}

		for _, hostId := range queryResponse.Resources {
			params := zero_trust_assessment.NewGetAssessmentV1ParamsWithContext(ctx).WithIds([]string{hostId})
			ztaOK, err := c.CrowdStrike.ZeroTrustAssessment.GetAssessmentV1(params)
			if err != nil {
				// If one of the host id is not enrolled to zta, whole thing fails
				// By checking one host id at a time, we are able to get ALL zta assesments
				fmt.Printf("could not get ZTA: %s\n", err.Error())
				continue
			}

			ztaResponse := ztaOK.GetPayload()
			if err = falcon.AssertNoError(ztaResponse.Errors); err != nil {
				return fmt.Errorf("Error querying ZTA: %s", falcon.ErrorExplain(err))
			}

			for _, zta := range ztaResponse.Resources {
				res <- zta
			}
		}

		offset += limit

		if *queryResponse.Meta.Pagination.Total == int64(*queryResponse.Meta.Pagination.Offset) {
			break
		}
	}

	return nil
}
