package main

import (
	"github.com/turbot/steampipe-plugin-rss/rss"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: rss.Plugin})
}
