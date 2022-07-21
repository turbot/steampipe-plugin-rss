package rss

import (
	"context"
	"errors"
	"sync"

	"github.com/mmcdole/gofeed"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableRSSItem(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "rss_item",
		Description: "An item may represent a story - much like a story in a newspaper or magazine; if so its description is a synopsis of the story, and the link points to the full story.",
		List: &plugin.ListConfig{
			Hydrate: listItem,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:       "feed_link",
					Require:    plugin.Optional,
					CacheMatch: "exact",
				},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "title", Type: proto.ColumnType_STRING, Description: "The title of the item."},
			{Name: "link", Type: proto.ColumnType_STRING, Description: "The URL of the item."},

			// Other columns
			{Name: "author_email", Type: proto.ColumnType_STRING, Transform: transform.FromField("Author.Email"), Description: "Email address of the author of the item."},
			{Name: "author_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Author.Name"), Description: "Name of the author of the item."},
			{Name: "categories", Type: proto.ColumnType_JSON, Description: "Specify one or more categories that the item belongs to."},
			{Name: "content", Type: proto.ColumnType_STRING, Description: "Content of the item."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "The item synopsis."},
			{Name: "enclosures", Type: proto.ColumnType_JSON, Description: "Media objects attached to the item."},
			{Name: "extensions", Type: proto.ColumnType_JSON, Description: "Extension data (e.g. Dublin Core, ITunes) for the item."},
			{Name: "feed_link", Type: proto.ColumnType_STRING, Description: "URL of the feed itself."},
			{Name: "guid", Type: proto.ColumnType_STRING, Description: "A string that uniquely identifies the item."},
			{Name: "image_title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Image.Title"), Description: ""},
			{Name: "image_url", Type: proto.ColumnType_STRING, Transform: transform.FromField("Image.URL"), Description: ""},
			{Name: "published", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("PublishedParsed"), Description: "Timestamp when the feed was published."},
			{Name: "updated", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("UpdatedParsed"), Description: "Timestamp when the feed was updated."},
		},
	}
}

type Item struct {
	gofeed.Item
	FeedLink string
}

func listItem(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	feedLink := d.KeyColumnQuals["feed_link"].GetStringValue()
	configFeedLinks := GetConfigFeedLink(ctx, d)
	if feedLink == "" && len(configFeedLinks) == 0 {
		return nil, errors.New("feed link must be provided either in where clause or should be configured as default feed links in the config file")
	}
	fp := gofeed.NewParser()
	if feedLink != "" {
		feed, err := fp.ParseURLWithContext(feedLink, ctx)
		if err != nil {
			plugin.Logger(ctx).Error("listItem", "Error", err)
			return nil, err
		}
		for _, item := range feed.Items {
			d.StreamListItem(ctx, Item{*item, feedLink})
		}

		return nil, nil
	}
	var wg sync.WaitGroup
	errorCh := make(chan error, len(configFeedLinks))
	for _, link := range configFeedLinks {
		wg.Add(1)
		go getItemAsync(ctx, d, link, fp, &wg, errorCh)
	}

	// wait for all feed links to be processed
	wg.Wait()

	// NOTE: close Item before ranging over results
	close(errorCh)

	for err := range errorCh {
		// return the first error
		return nil, err
	}

	return nil, nil
}

func getItemAsync(ctx context.Context, d *plugin.QueryData, link string, fp *gofeed.Parser, wg *sync.WaitGroup, errorCh chan error) {
	defer wg.Done()

	err := getItemDetails(ctx, d, link, fp)
	if err != nil {
		errorCh <- err
	}
}

func getItemDetails(ctx context.Context, d *plugin.QueryData, link string, fp *gofeed.Parser) error {
	feed, err := fp.ParseURLWithContext(link, ctx)
	if err != nil {
		plugin.Logger(ctx).Error("getItemDetails", "Error", err)
		return err
	}
	for _, item := range feed.Items {
		d.StreamListItem(ctx, Item{*item, link})
	}

	return nil
}
