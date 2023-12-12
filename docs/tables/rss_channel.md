---
title: "Steampipe Table: rss_channel - Query RSS Channel using SQL"
description: "Allows users to query RSS Channels, specifically providing details about the channel such as title, description, link, language, and copyright."
---

# Table: rss_channel - Query RSS Channel using SQL

RSS (Really Simple Syndication) Channel is a web feed that allows users to access updates to online content in a standardized, computer-readable format. These feeds can, for example, allow a user to keep track of many different websites in a single news aggregator. The news aggregator will automatically check the RSS feed for new content, allowing the content to be automatically passed from website to user.

## Table Usage Guide

The `rss_channel` table provides insights into RSS channels across various websites. As a content manager or a web developer, explore channel-specific details through this table, including title, description, link, language, and copyright. Utilize it to monitor the updates of different websites in an automated and efficient manner.

**Important Notes**
- It's not possible to list all feeds in the world, so this table requires a `feed_link` qualifier to be passed in the `where` or `join` clause for all queries.

## Examples

### Basic channel info
Discover the segments of a specific RSS feed by analyzing the title, link, and description. This could be beneficial for understanding the content and structure of the feed for further analysis or content curation.

```sql+postgres
select
  title,
  link,
  description
from
  rss_channel
where
  feed_link = 'https://steampipe.io/blog/feed.xml';
```

```sql+sqlite
select
  title,
  link,
  description
from
  rss_channel
where
  feed_link = 'https://steampipe.io/blog/feed.xml';
```

### Basic channel info for multiple feeds
Explore multiple RSS feeds to uncover their basic channel information. This can be useful for comparing content across different sources or tracking updates from preferred feeds.

```sql+postgres
select
  title,
  link,
  description
from
  rss_channel
where
  feed_link in ('https://steampipe.io/blog/feed.xml','https://www.podcastinsights.com/feed/');
```

```sql+sqlite
select
  title,
  link,
  description
from
  rss_channel
where
  feed_link in ('https://steampipe.io/blog/feed.xml','https://www.podcastinsights.com/feed/');
```

### Get the type of the channel
Explore the type and version of a specific RSS feed to understand its format and compatibility. This is particularly useful when integrating or troubleshooting RSS feed readers.

```sql+postgres
select
  title,
  feed_type,
  feed_version
from
  rss_channel
where
  feed_link = 'https://steampipe.io/blog/feed.xml';
```

```sql+sqlite
select
  title,
  feed_type,
  feed_version
from
  rss_channel
where
  feed_link = 'https://steampipe.io/blog/feed.xml';
```

### Get categories for the channel
Discover the segments that categorize a specific podcast channel, which can be useful for understanding the content scope and audience interest areas of the channel.

```sql+postgres
select
  title,
  categories
from
  rss_item
where
  feed_link = 'https://www.podcastinsights.com/feed/';
```

```sql+sqlite
select
  title,
  categories
from
  rss_item
where
  feed_link = 'https://www.podcastinsights.com/feed/';
```