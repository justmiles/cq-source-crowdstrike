package resources

import (
	"context"
	"fmt"

	"github.com/justmiles/cq-source-crowdstrike/client"

	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/crowdstrike/gofalcon/falcon/client/incidents"
	"github.com/crowdstrike/gofalcon/falcon/models"
)

func Incidents() *schema.Table {
	return &schema.Table{
		Name:      "crowdstrike_falcon_incidents",
		Resolver:  fetchIncidents,
		Transform: transformers.TransformWithStruct(&models.DomainIncident{}, transformers.WithPrimaryKeys("IncidentID")),
	}
}

func fetchIncidents(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	queryIncidentsOK, err := c.CrowdStrike.Incidents.QueryIncidents(&incidents.QueryIncidentsParams{
		Context: context.Background(),
	})
	if err != nil {
		return fmt.Errorf("could not query incident: %s", err.Error())
	}
	queryResponse := queryIncidentsOK.GetPayload()
	var ids []string
	for _, id := range queryResponse.Resources {
		if idstr, ok := id.(string); ok {
			ids = append(ids, string(idstr))
		}
	}

	getIncidentsOK, err := c.CrowdStrike.Incidents.GetIncidents(&incidents.GetIncidentsParams{
		Context: context.Background(),
		Body: &models.MsaIdsRequest{
			Ids: ids,
		},
	})
	if err != nil {
		return fmt.Errorf("could not get incidents: %s", err.Error())
	}

	incidentResponse := getIncidentsOK.GetPayload()

	for _, incident := range incidentResponse.Resources {
		res <- incident
	}

	return nil
}
