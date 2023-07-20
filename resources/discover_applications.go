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

func fetchDiscoverApps(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	response, err := c.CrowdStrike.Discover.GetApplications(&discover.GetApplicationsParams{
		Context: ctx,
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

	return nil
}
