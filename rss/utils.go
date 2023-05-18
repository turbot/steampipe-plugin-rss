package rss

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func feedLink(_ context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.EqualsQuals
	fl := quals["feed_link"].GetStringValue()
	return fl, nil
}

func shouldIgnoreErrors(errors []string) plugin.ErrorPredicateWithContext {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
		errorStrings := []string{"403", "404"}
		for _, errStr := range errorStrings {
			if strings.Contains(err.Error(), errStr) {
				return true
			}
		}
		return false
	}
}

func handleFeedError(err error) bool {
	errorStrings := []string{"x509: certificate has expired or is not yet valid", "502"}
	for _, s := range errorStrings {
		if strings.Contains(err.Error(), s) {
			return true
		}
	}
	return false
}
