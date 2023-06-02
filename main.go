package main

import (
	"github.com/justmiles/cq-source-crowdstrike/plugin"

	"github.com/cloudquery/plugin-sdk/v3/serve"
)

func main() {
	serve.Source(plugin.Plugin())
}
