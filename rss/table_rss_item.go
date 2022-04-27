package rss

import (
	"context"

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
			KeyColumns: plugin.SingleColumn("feed_link"),
			Hydrate:    listItem,
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
			{Name: "feed_link", Type: proto.ColumnType_STRING, Hydrate: feedLink, Transform: transform.FromValue(), Description: "URL of the feed itself."},
			{Name: "guid", Type: proto.ColumnType_STRING, Description: "A string that uniquely identifies the item."},
			{Name: "image_title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Image.Title"), Description: ""},
			{Name: "image_url", Type: proto.ColumnType_STRING, Transform: transform.FromField("Image.URL"), Description: ""},
			{Name: "published", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("PublishedParsed"), Description: "Timestamp when the feed was published."},
			{Name: "updated", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("UpdatedParsed"), Description: "Timestamp when the feed was updated."},
		},
	}
}

func listItem(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	fl := quals["feed_link"].GetStringValue()

	fp := gofeed.NewParser()
	feed, err := fp.ParseURLWithContext(fl, ctx)
	if err != nil {
		plugin.Logger(ctx).Error("listItem", "Error", err)
		return nil, err
	}

	for _, i := range feed.Items {
		d.StreamListItem(ctx, i)
	}

	return nil, nil
}
