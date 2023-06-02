package resources

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"

	"github.com/crowdstrike/gofalcon/falcon"
	"github.com/crowdstrike/gofalcon/falcon/client/incidents"
	"github.com/crowdstrike/gofalcon/falcon/models"
)

func Incidents() *schema.Table {
	return &schema.Table{
		Name:      "crowdstrike_falcon_incidents",
		Resolver:  fetchIncidents,
		Transform: transformers.TransformWithStruct(&models.DomainIncident{}),
	}
}

func fetchIncidents(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
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

	queryIncidentsOK, err := client.Incidents.QueryIncidents(&incidents.QueryIncidentsParams{
		Context: context.Background(),
	})
	if err != nil {
		return fmt.Errorf("could not get incident: %s", err.Error())
	}
	queryResponse := queryIncidentsOK.GetPayload()
	var ids []string
	for _, id := range queryResponse.Resources {
		if idstr, ok := id.(string); ok {
			ids = append(ids, string(idstr))
		}
	}

	getIncidentsOK, err := client.Incidents.GetIncidents(&incidents.GetIncidentsParams{
		Context: context.Background(),
		Body: &models.MsaIdsRequest{
			Ids: ids,
		},
	})

	incidentResponse := getIncidentsOK.GetPayload()

	for _, incident := range incidentResponse.Resources {
		res <- incident
	}

	return nil
}

// func Run() error {

// 	// falconClientID := os.Getenv("FALCON_CLIENT_ID")
// 	// falconSecret := os.Getenv("FALCON_SECRET")

// 	falconClientID := "deb028d3f4c64d789b3160e372cf8bb5"
// 	falconSecret := "4XpnAH8bDSPo5JzxfC631NiQGM0V9gcurBUOFq72"

// 	client, err := falcon.NewClient(&falcon.ApiConfig{
// 		ClientId:     falconClientID,
// 		ClientSecret: falconSecret,
// 		Context:      context.Background(),
// 	})
// 	if err != nil {
// 		return fmt.Errorf("could not auth: %s", err.Error())
// 	}

// 	res, err := client.Incidents.QueryIncidents(&incidents.QueryIncidentsParams{
// 		Context: context.Background(),
// 	})
// 	if err != nil {
// 		return fmt.Errorf("could not get incident: %s", err.Error())
// 	}
// 	queryResponse := res.GetPayload()
// 	var ids []string
// 	for _, id := range queryResponse.Resources {
// 		if idstr, ok := id.(string); ok {
// 			ids = append(ids, string(idstr))
// 		}
// 	}

// 	resI, err := client.Incidents.GetIncidents(&incidents.GetIncidentsParams{
// 		Context: context.Background(),
// 		Body: &models.MsaIdsRequest{
// 			Ids: ids,
// 		},
// 	})

// 	incidentResponse := resI.GetPayload()

// 	for _, incident := range incidentResponse.Resources {
// 		spew.Dump(incident)
// 	}

// 	return nil
// }

// func GetIncident(id string) error {
// 	falconClientID := "deb028d3f4c64d789b3160e372cf8bb5"
// 	falconSecret := "4XpnAH8bDSPo5JzxfC631NiQGM0V9gcurBUOFq72"

// 	client, err := falcon.NewClient(&falcon.ApiConfig{
// 		ClientId:     falconClientID,
// 		ClientSecret: falconSecret,
// 		Context:      context.Background(),
// 	})

// }
