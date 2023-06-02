package resources

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"

	"github.com/crowdstrike/gofalcon/falcon"
	"github.com/crowdstrike/gofalcon/falcon/client/hosts"
	"github.com/crowdstrike/gofalcon/falcon/models"
)

func Hosts() *schema.Table {
	return &schema.Table{
		Name:      "crowdstrike_falcon_hosts",
		Resolver:  fetchHosts,
		Transform: transformers.TransformWithStruct(&models.DeviceapiDeviceSwagger{}),
	}
}

func fetchHosts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
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

	var more = true
	var offset int64 = 0
	var limit int64 = 100
	for more {

		queryRespOK, err := client.Hosts.QueryDevicesByFilter(&hosts.QueryDevicesByFilterParams{
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
				return fmt.Errorf("error QueryDevicesByFilter: %s", *msaAPIError.Message)
			}
		}

		time.Sleep(1 * time.Second)
		offset += limit
		if *queryResponse.Meta.Pagination.Total == int64(*queryResponse.Meta.Pagination.Offset) {
			more = false
		}
		detailsOK, err := client.Hosts.GetDeviceDetailsV2(&hosts.GetDeviceDetailsV2Params{
			Context: ctx,
			Ids:     queryResponse.Resources,
		})

		detailsResponse := detailsOK.GetPayload()

		for _, host := range detailsResponse.Resources {
			res <- host
		}

	}

	return nil
}
