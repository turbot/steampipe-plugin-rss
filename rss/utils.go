package rss

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func feedLink(_ context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	fl := quals["feed_link"].GetStringValue()
	return fl, nil
}
