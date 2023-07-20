package resources

import (
	"context"
	"fmt"

	"github.com/justmiles/cq-source-crowdstrike/client"

	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/crowdstrike/gofalcon/falcon"
	"github.com/crowdstrike/gofalcon/falcon/client/discover"
	"github.com/crowdstrike/gofalcon/falcon/models"
)

func DiscoverHosts() *schema.Table {
	return &schema.Table{
		Name:      "crowdstrike_falcon_discover_hosts",
		Resolver:  fetchDiscoverHosts,
		Transform: transformers.TransformWithStruct(&models.DomainDiscoverAPIHost{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchDiscoverHosts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	limit := int64(100)
	filter := ""

	for offset := int64(0); ; {
		response, err := c.CrowdStrike.Discover.QueryHosts(&discover.QueryHostsParams{
			Context: ctx,
			Limit:   &limit,
			Offset:  &offset,
			Filter:  &filter,
		})
		if err != nil {
			return fmt.Errorf("could not get hostIDs: %s", falcon.ErrorExplain(err))
		}
		if err = falcon.AssertNoError(response.Payload.Errors); err != nil {
			return fmt.Errorf("could not get hostIDs: %s", err.Error())
		}

		hosts := response.Payload.Resources
		if len(hosts) == 0 {
			break
		}

		hostDetails, err := c.CrowdStrike.Discover.GetHosts(&discover.GetHostsParams{
			Ids:     hosts,
			Context: ctx,
		})
		if err != nil {
			return fmt.Errorf("could not get hosts: %s", falcon.ErrorExplain(err))
		}
		if err = falcon.AssertNoError(response.Payload.Errors); err != nil {
			return fmt.Errorf("could not get hosts: %s", err.Error())
		}

		for _, host := range hostDetails.Payload.Resources {
			res <- host
		}
		
		offset = offset + int64(len(hosts))
		if offset >= *response.Payload.Meta.Pagination.Total {
			break
		}
	}
	
	return nil
}
