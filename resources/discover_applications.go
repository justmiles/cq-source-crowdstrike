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

func DiscoverApps() *schema.Table {
	return &schema.Table{
		Name:      "crowdstrike_falcon_discover_applications",
		Resolver:  fetchDiscoverApps,
		Transform: transformers.TransformWithStruct(&models.DomainDiscoverAPIApplication{}, transformers.WithPrimaryKeys("ID")),
	}
}

func getHostIds(ctx context.Context, meta schema.ClientMeta, filter string) <-chan []string {
	c := meta.(*client.Client)
	hostIds := make(chan []string)

	go func() {
		limit := int64(100)
		for offset := int64(0); ; {
			response, err := c.CrowdStrike.Discover.QueryHosts(&discover.QueryHostsParams{
				Context: ctx,
				Limit:   &limit,
				Offset:  &offset,
				Filter:  &filter,
			})
			if err != nil {
				panic(falcon.ErrorExplain(err))
			}
			if err = falcon.AssertNoError(response.Payload.Errors); err != nil {
				panic(err)
			}

			hosts := response.Payload.Resources
			if len(hosts) == 0 {
				break
			}
			hostIds <- hosts
			offset = offset + int64(len(hosts))
			if offset >= *response.Payload.Meta.Pagination.Total {
				break
			}
		}
		close(hostIds)
	}()
	return hostIds
}

func getAppIds(ctx context.Context, meta schema.ClientMeta, hostIds []string) <-chan []string {
	c := meta.(*client.Client)
	
	appIds := make(chan []string)

	for _, hostID := range hostIds {
		go func() {
			filter := "host.id:'" + hostID + "'"
			limit := int64(100)
			for offset := int64(0); ; {
				response, err := c.CrowdStrike.Discover.QueryApplications(&discover.QueryApplicationsParams{
					Context: ctx,
					Limit:   &limit,
					Offset:  &offset,
					Filter:  &filter,
				})
				if err != nil {
					panic(falcon.ErrorExplain(err))
				}
				if err = falcon.AssertNoError(response.Payload.Errors); err != nil {
					panic(err)
				}

				apps := response.Payload.Resources
				if len(apps) == 0 {
					break
				}
				appIds <- apps
				offset = offset + int64(len(apps))
				if offset >= *response.Payload.Meta.Pagination.Total {
					break
				}
			}
		}()
		close(appIds)
	}
	return appIds
}

func fetchDiscoverApps(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	hostFilter := "entity_type:'managed'"

	for hostIDbatch := range getHostIds(ctx, meta, hostFilter) {
		for appIDsbatch := range getAppIds(ctx, meta, hostIDbatch) {
			response, err := c.CrowdStrike.Discover.GetApplications(&discover.GetApplicationsParams{
				Context: ctx,
				Ids: appIDsbatch,
			})
			if err != nil {
				return fmt.Errorf("could not get Applications: %s", falcon.ErrorExplain(err))
			}
			if err = falcon.AssertNoError(response.Payload.Errors); err != nil {
				return fmt.Errorf("could not get Applications: %s", err.Error())
			}
		
			apps := response.Payload.Resources
			if len(apps) > 0 {
				for _, app := range apps {
					res <- app
				}
			}
		}
	}

	return nil
}
