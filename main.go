package main

import (
	"github.com/turbot/steampipe-plugin-databricks/databricks"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: databricks.Plugin})
}
