package resources

import (
	"context"
	"fmt"

	"github.com/justmiles/cq-source-crowdstrike/client"

	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/crowdstrike/gofalcon/falcon/client/hosts"
	"github.com/crowdstrike/gofalcon/falcon/models"
)

func Hosts() *schema.Table {
	return &schema.Table{
		Name:      "crowdstrike_falcon_hosts",
		Resolver:  fetchHosts,
		Transform: transformers.TransformWithStruct(&models.DeviceapiDeviceSwagger{}, transformers.WithPrimaryKeys("DeviceID")),
	}
}

func fetchHosts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
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
			return fmt.Errorf("could not get incident: %s", err.Error())
		}
		queryResponse := queryRespOK.GetPayload()
		for _, msaAPIError := range queryResponse.Errors {
			if msaAPIError.Message != nil {
				return fmt.Errorf("could not query hosts: %s", *msaAPIError.Message)
			}
		}

		detailsOK, err := c.CrowdStrike.Hosts.GetDeviceDetailsV2(&hosts.GetDeviceDetailsV2Params{
			Context: ctx,
			Ids:     queryResponse.Resources,
		})
		if err != nil {
			return fmt.Errorf("could not get hosts: %s", err.Error())
		}

		detailsResponse := detailsOK.GetPayload()

		for _, host := range detailsResponse.Resources {
			res <- host
		}

		offset += limit

		if *queryResponse.Meta.Pagination.Total == int64(*queryResponse.Meta.Pagination.Offset) {
			break
		}
	}

	return nil
}
