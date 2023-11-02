package client

import (
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/crowdstrike/gofalcon/falcon/client"
	"github.com/rs/zerolog"
)

type Client struct {
	Logger  zerolog.Logger
	Spec    Spec
	Backend state.Client

	CrowdStrike *client.CrowdStrikeAPISpecification
}

func (c *Client) ID() string {
	return "CrowdStrike"
}

func New(logger zerolog.Logger, spec Spec, services *client.CrowdStrikeAPISpecification, bk state.Client) *Client {
	return &Client{
		Logger:      logger,
		Spec:        spec,
		Backend:     bk,
		CrowdStrike: services,
	}
}
