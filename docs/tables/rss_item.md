---
title: "Steampipe Table: rss_item - Query RSS Feed Items using SQL"
description: "Allows users to query RSS Feed Items, specifically the details of each item, providing insights into the content and metadata of RSS feeds."
---

# Table: rss_item - Query RSS Feed Items using SQL

An RSS Feed Item is an individual content piece from an RSS feed. These items contain the actual content of the feed, including titles, descriptions, and links to the original content. Each item also carries metadata such as publishing dates, authorship, and categories.

## Table Usage Guide

The `rss_item` table provides insights into individual items within an RSS feed. As a content manager or data analyst, explore item-specific details through this table, including content, metadata, and associated links. Utilize it to uncover information about items, such as their publishing timeline, authorship details, and categorization.

**Important Notes**
- It's not possible to list all feeds in the world, so this table requires a `feed_link` qualifier to be passed in the `where` or `join` clause for all queries.

## Examples

### Query items from a channel, newest first
Explore the latest items from a specific RSS feed to stay updated with the most recent content. This is especially useful for keeping track of the latest updates on frequently updated websites or blogs.

```sql
select
  title,
  published,
  link
from
  rss_item
where
  feed_link = 'https://www.hardcorehumanism.com/feed/'
order by
  published desc;
```

### Count items by category
Explore the distribution of podcast topics from a specific feed. This query helps you understand the most frequent themes or categories in the chosen podcast, providing insights into its content focus.

```sql
select
  category,
  count(*)
from
  rss_item,
  jsonb_array_elements_text(categories) as category
where
  feed_link = 'https://www.podcastinsights.com/feed/'
group by
  category
order by
  count desc;
```