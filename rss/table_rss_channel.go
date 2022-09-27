package rss

import (
	"context"

	"github.com/mmcdole/gofeed"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableRSSChannel(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "rss_channel",
		Description: "Information about the RSS channel or Atom feed.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("feed_link"),
			Hydrate:    listChannel,
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

func listChannel(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	fl := quals["feed_link"].GetStringValue()
	fp := gofeed.NewParser()
	feed, err := fp.ParseURLWithContext(fl, ctx)
	if err != nil {
		return nil, err
	}
	d.StreamListItem(ctx, feed)
	return nil, nil
}
