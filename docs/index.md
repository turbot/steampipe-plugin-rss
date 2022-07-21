---
organization: Turbot
category: ["media"]
icon_url: "/images/plugins/turbot/rss.svg"
brand_color: "#FFA500"
display_name: RSS
name: rss
description: Steampipe plugin to query RSS channels & Atom feeds
og_description: "Query RSS with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/rss-social-graphic.png"
---

# RSS + Steampipe

[RSS Channel](https://en.wikipedia.org/wiki/RSS) and [Atom Feed](https://en.wikipedia.org/wiki/Atom_(Web_standard)) content including descriptions, items, links and metadata.

Data for the different feed types is normalized into a common set of columns.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query data using SQL.

For example:

```sql
select
  title,
  link,
  description
from
  rss_item
where
  feed_link = 'https://steampipe.io/blog/feed.xml';
```

```sh
+----------------+---------------------------+----------------------------------------+
| title          | link                      | description                            |
+----------------+---------------------------+----------------------------------------+
| Steampipe Blog | https://steampipe.io/blog | Blog and Resource Center for Steampipe |
+----------------+---------------------------+----------------------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/rss/tables)**

## Installation

Download and install the latest rss plugin:

```bash
steampipe plugin install rss
```

### Credentials

No credentials are required.

### Configuration

Installing the latest rss plugin will create a config file (`~/.steampipe/config/rss.spc`) with a single connection named `rss`:

```hcl
connection "rss" {
  plugin = "rss"

  # Specify the default list of feed links for all the tables
  # Feed link must be provided either in where clause or should be configured as default feed links in the config file.
  # feed_links = ["https://steampipe.io/blog/feed.xml","https://www.podcastinsights.com/feed/"]
}
```

- `feed_links` - A list of default feed links. Feed link must be provided either in where clause or should be configured as default feed links in the config file.

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-rss
- Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)