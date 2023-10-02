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
