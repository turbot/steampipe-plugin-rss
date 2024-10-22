## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#55](https://github.com/turbot/steampipe-plugin-rss/pull/55))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#55](https://github.com/turbot/steampipe-plugin-rss/pull/55))

## v0.7.0 [2024-03-20]

_Enhancements_

- Recompiled plugin with [mmcdole/gofeed v1.3.0](https://github.com/mmcdole/gofeed). ([#50](https://github.com/turbot/steampipe-plugin-rss/pull/50))

## v0.6.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#47](https://github.com/turbot/steampipe-plugin-rss/pull/47))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#47](https://github.com/turbot/steampipe-plugin-rss/pull/47))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-rss/blob/main/docs/LICENSE). ([#47](https://github.com/turbot/steampipe-plugin-rss/pull/47))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#46](https://github.com/turbot/steampipe-plugin-rss/pull/46))

## v0.5.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#39](https://github.com/turbot/steampipe-plugin-rss/pull/39))

## v0.5.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#37](https://github.com/turbot/steampipe-plugin-rss/pull/37))
- Recompiled plugin with Go version `1.21`. ([#37](https://github.com/turbot/steampipe-plugin-rss/pull/37))

## v0.4.0 [2023-03-23]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#28](https://github.com/turbot/steampipe-plugin-rss/pull/28))

## v0.3.0 [2022-09-27]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements. ([#26](https://github.com/turbot/steampipe-plugin-rss/pull/26))
- Recompiled plugin with Go version `1.19`. ([#26](https://github.com/turbot/steampipe-plugin-rss/pull/26))

## v0.2.1 [2022-05-23]

_Bug fixes_

- Fixed the Slack community links in README and docs/index.md files. ([#20](https://github.com/turbot/steampipe-plugin-rss/pull/20))

## v0.2.0 [2022-04-27]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds. ([#17](https://github.com/turbot/steampipe-plugin-rss/pull/17))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#16](https://github.com/turbot/steampipe-plugin-rss/pull/16))

## v0.1.0 [2021-11-23]

_Enhancements_

- Recompiled plugin with Go version 1.17 ([#12](https://github.com/turbot/steampipe-plugin-rss/pull/12))
- Recompiled plugin with [steampipe-plugin-sdk v1.8.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v182--2021-11-22) ([#11](https://github.com/turbot/steampipe-plugin-rss/pull/11))

## v0.0.3 [2021-10-06]

_Enhancements_

- Updated the README file to include GitHub clone link and license information ([#8](https://github.com/turbot/steampipe-plugin-rss/pull/8))
- Updated plugin license to Apache 2.0 per [turbot/steampipe#488](https://github.com/turbot/steampipe/issues/488) ([#5](https://github.com/turbot/steampipe-plugin-rss/pull/5))

## v0.0.2 [2021-03-18]

_Enhancements_

- Updated examples for `rss_channel` and `rss_item` tables ([#3](https://github.com/turbot/steampipe-plugin-rss/pull/3))
- Recompiled plugin with [steampipe-plugin-sdk v0.2.4](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v024-2021-03-16)

## v0.0.1 [2021-03-04]

_What's new?_

- Initial release with tables

  - rss_channel
  - rss_item
