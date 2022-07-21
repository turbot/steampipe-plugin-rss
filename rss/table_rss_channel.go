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

func tableRSSChannel(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "rss_channel",
		Description: "Information about the RSS channel or Atom feed.",
		List: &plugin.ListConfig{
			Hydrate: listChannel,
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
			{Name: "title", Type: proto.ColumnType_STRING, Description: "The name of the channel. It's how people refer to your service. If you have an HTML website that contains the same information as your RSS file, the title of your channel should be the same as the title of your website."},
			{Name: "link", Type: proto.ColumnType_STRING, Description: "The URL to the HTML website corresponding to the channel."},

			// Other columns
			{Name: "author_email", Type: proto.ColumnType_STRING, Transform: transform.FromField("Author.Email"), Description: "Email of the author."},
			{Name: "author_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Author.Name"), Description: "Name of the author."},
			{Name: "categories", Type: proto.ColumnType_JSON, Description: "Specify one or more categories that the channel belongs to."},
			{Name: "copyright", Type: proto.ColumnType_STRING, Description: "Copyright notice for content in the channel."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Phrase or sentence describing the channel."},
			{Name: "extensions", Type: proto.ColumnType_JSON, Description: "Extension data (e.g. Dublin Core, ITunes) for the channel."},
			{Name: "feed_link", Type: proto.ColumnType_STRING, Description: "URL of the feed itself."},
			{Name: "feed_type", Type: proto.ColumnType_STRING, Description: "Type of the feed: rss, atom or json."},
			{Name: "feed_version", Type: proto.ColumnType_STRING, Description: "Version of the feed type, e.g. 2.0"},
			{Name: "generator", Type: proto.ColumnType_STRING, Description: "A string indicating the program used to generate the channel."},
			{Name: "image_title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Image.Title"), Description: "The title of a GIF, JPEG or PNG image that can be displayed with the channel."},
			{Name: "image_url", Type: proto.ColumnType_STRING, Transform: transform.FromField("Image.URL"), Description: "The URL of a GIF, JPEG or PNG image that can be displayed with the channel."},
			{Name: "language", Type: proto.ColumnType_STRING, Description: "The language the channel is written in. This allows aggregators to group all Italian language sites, for example, on a single page. A list of allowable values for this element, as provided by Netscape, is here. You may also use values defined by the W3C."},
			{Name: "published", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("PublishedParsed"), Description: "Timestamp when the feed was published."},
			{Name: "updated", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("UpdatedParsed"), Description: "Timestamp when the feed was updated."},
		},
	}
}

type Channel struct {
	gofeed.Feed
	FeedLink string
}

func listChannel(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	feedLink := d.KeyColumnQuals["feed_link"].GetStringValue()
	configFeedLinks := GetConfigFeedLink(ctx, d)
	if feedLink == "" && len(configFeedLinks) == 0 {
		return nil, errors.New("feed link must be provided either in where clause or should be configured as default feed links in the config file")
	}
	fp := gofeed.NewParser()
	if feedLink != "" {
		feed, err := fp.ParseURLWithContext(feedLink, ctx)
		if err != nil {
			plugin.Logger(ctx).Error("listChannel", "Error", err)
			return nil, err
		}
		d.StreamListItem(ctx, Channel{*feed, feedLink})

		return nil, nil
	}
	var wg sync.WaitGroup
	errorCh := make(chan error, len(configFeedLinks))
	for _, link := range configFeedLinks {
		wg.Add(1)
		go getChannelAsync(ctx, d, link, fp, &wg, errorCh)
	}

	// wait for all feed links to be processed
	wg.Wait()

	// NOTE: close channel before ranging over results
	close(errorCh)

	for err := range errorCh {
		// return the first error
		return nil, err
	}

	return nil, nil
}

func getChannelAsync(ctx context.Context, d *plugin.QueryData, link string, fp *gofeed.Parser, wg *sync.WaitGroup, errorCh chan error) {
	defer wg.Done()

	err := getChannelDetails(ctx, d, link, fp)
	if err != nil {
		errorCh <- err
	}
}

func getChannelDetails(ctx context.Context, d *plugin.QueryData, link string, fp *gofeed.Parser) error {
	feed, err := fp.ParseURLWithContext(link, ctx)
	if err != nil {
		plugin.Logger(ctx).Error("getChannelDetails", "Error", err)
		return err
	}
	d.StreamListItem(ctx, Channel{*feed, link})
	return nil
}
