package rss

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func feedLink(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	fl := quals["feed_link"].GetStringValue()
	return fl, nil
}
