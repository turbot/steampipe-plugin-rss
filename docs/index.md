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
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# RSS

[RSS Channel](https://en.wikipedia.org/wiki/RSS) and [Atom Feed](https://en.wikipedia.org/wiki/Atom_(Web_standard)) content including descriptions, items, links and metadata.

Data for the different feed types is normalized into a common set of columns.


## Installation

Download and install the latest rss plugin:

```bash
steampipe plugin install rss
```

## Run a query

```bash
$ steampipe query
Welcome to Steampipe v0.2.2
For more information, type .help
> select title, link, description from rss_channel where feed_link = 'https://steampipe.io/blog/feed.xml'
+----------------+---------------------------+----------------------------------------+
| title          | link                      | description                            |
+----------------+---------------------------+----------------------------------------+
| Steampipe Blog | https://steampipe.io/blog | Blog and Resource Center for Steampipe |
+----------------+---------------------------+----------------------------------------+
```
