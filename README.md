![image](https://hub.steampipe.io/images/plugins/turbot/rss-social-graphic.png)

# RSS Plugin for Steampipe

Use SQL to query RSS channels and Atom Feeds.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/rss)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/rss/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-rss/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install rss
```

Run a query:

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

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-rss.git
cd steampipe-plugin-rss
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/rss.spc
```

Try it!

```
steampipe query
> .inspect rss
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). Contributions to the plugin are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-rss/blob/main/LICENSE). Contributions to the plugin documentation are subject to the [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-rss/blob/main/docs/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [RSS Plugin](https://github.com/turbot/steampipe-plugin-rss/labels/help%20wanted)
