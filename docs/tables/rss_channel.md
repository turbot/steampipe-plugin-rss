# Table: rss_channel

View RSS channel or Atom feed information including title, link, description and metadata.

Note: It's not possible to list all feeds in the world, so this table requires a
`feed_link` qualifier to be passed in the `where` or `join` clause for all queries.


## Examples

### Basic channel info

```sql
select
  title,
  link,
  description
from
  rss_channel
where
  feed_link = 'https://steampipe.io/blog/feed.xml'
```

### Basic channel info for multiple feeds

```sql
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

```sql
select
  title,
  feed_type,
  feed_version
from
  rss_channel
where
  feed_link = 'https://steampipe.io/blog/feed.xml'
```

### Get categories for the channel

```sql
select
  title,
  categories
from
  rss_item
where
  feed_link = 'https://www.podcastinsights.com/feed/'
```
