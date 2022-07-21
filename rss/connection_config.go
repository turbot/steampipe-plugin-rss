package rss

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/schema"
)

type rssConfig struct {
	FeedLinks []string `cty:"feed_links"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"feed_links": {
		Type: schema.TypeList,
		Elem: &schema.Attribute{Type: schema.TypeString},
	},
}

func ConfigInstance() interface{} {
	return &rssConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) rssConfig {
	if connection == nil || connection.Config == nil {
		return rssConfig{}
	}
	config, _ := connection.Config.(rssConfig)
	return config
}

func GetConfigFeedLink(ctx context.Context, d *plugin.QueryData) []string {
	config := GetConfig(d.Connection)
	if len(config.FeedLinks) > 0 {
		return config.FeedLinks
	}
	return nil
}
