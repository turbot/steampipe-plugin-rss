# Table: rss_item

Query items from the RSS channel or Atom feed. Item information includes title, link, content, categories and other metadata.

Note: It's not possible to list all feeds in the world, so this table requires a
`feed_link` qualifier to be passed in the where clause for all queries.


## Examples

### Query items from a channel

```sql
select
  title,
  published,
  link
from
  rss_item
where
  feed_link = 'https://www.hardcorehumanism.com/feed/'
```

### Count items by category

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
  count desc
```
