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

func isNotFoundError(notFoundErrors []string) plugin.ErrorPredicate {
	return func(err error) bool {

		for _, pattern := range notFoundErrors {
			if strings.Contains(err.Error(), pattern) {
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
