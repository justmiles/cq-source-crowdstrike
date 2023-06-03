package resources

import (
	"context"
	"fmt"

	"github.com/justmiles/cq-source-crowdstrike/client"

	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/crowdstrike/gofalcon/falcon"
	"github.com/crowdstrike/gofalcon/falcon/client/spotlight_vulnerabilities"
	"github.com/crowdstrike/gofalcon/falcon/models"
)

func Vulnerabilities() *schema.Table {
	return &schema.Table{
		Name:      "crowdstrike_falcon_vulnerabilities",
		Resolver:  fetchVulnerabilities,
		Transform: transformers.TransformWithStruct(&models.DomainBaseAPIVulnerabilityV2{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchVulnerabilities(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	lastSeen := (*string)(nil)
	for {

		queryRespOK, err := c.CrowdStrike.SpotlightVulnerabilities.CombinedQueryVulnerabilities(&spotlight_vulnerabilities.CombinedQueryVulnerabilitiesParams{
			Context: ctx,
			Facet:   []string{"cve", "host_info", "remediation"},
			After:   lastSeen,
			Filter:  "status:'open'",
		})
		if err != nil {
			return fmt.Errorf("could not get incident: %s", err.Error())
		}
		if err = falcon.AssertNoError(queryRespOK.Payload.Errors); err != nil {
			return fmt.Errorf(falcon.ErrorExplain(err))
		}

		queryResponse := queryRespOK.GetPayload()
		for _, msaAPIError := range queryResponse.Errors {
			if msaAPIError.Message != nil {
				return fmt.Errorf("error QueryDevicesByFilter: %s", *msaAPIError.Message)
			}
		}

		for _, vulnerability := range queryResponse.Resources {
			res <- vulnerability
		}

		if queryRespOK.Payload.Meta == nil && queryRespOK.Payload.Meta.Pagination == nil && queryRespOK.Payload.Meta.Pagination.Limit == nil {
			return fmt.Errorf("Cannot paginate Vulnerabilities, pagination information missing")
		}
		if *queryRespOK.Payload.Meta.Pagination.Limit > int32(len(queryResponse.Resources)) {
			break
		} else {
			lastSeen = queryRespOK.Payload.Meta.Pagination.After
		}
	}

	return nil
}
