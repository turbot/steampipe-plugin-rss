package rss

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-rss",
		DefaultTransform: transform.FromGo().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"rss_channel": tableRSSChannel(ctx),
			"rss_item":    tableRSSItem(ctx),
		},
		DefaultShouldIgnoreError: isNotFoundError([]string{"403","404"}),
	}
	return p
}
