# Table: databricks_sql_dashboard

Users can write SQL queries to analyze data and create charts, tables, and other visualizations directly from the query results. Databricks SQL Dashboards support various chart types like line charts, bar charts, pie charts, and more, which can be customized and interactively explored.

## Examples

### Basic info

```sql
select
  id,
  name,
  created_at,
  slug,
  account_id
from
  databricks_sql_dashboard;
```

### List dashboards created in the last 7 days

```sql
select
  id,
  name,
  created_at,
  slug,
  account_id
from
  databricks_sql_dashboard
where
  created_at >= now() - interval '7' day;
```

### List dashboards that are editable

```sql
select
  id,
  name,
  created_at,
  slug,
  account_id
from
  databricks_sql_dashboard
where
  can_edit;
```

### List dashboards that are archived

```sql
select
  id,
  name,
  created_at,
  slug,
  account_id
from
  databricks_sql_dashboard
where
  is_archived;
```

### List dashboards that are marked as favourite

```sql
select
  id,
  name,
  created_at,
  slug,
  account_id
from
  databricks_sql_dashboard
where
  is_favorite;
```

### List dashboards that have filters enabled

```sql
select
  id,
  name,
  created_at,
  slug,
  account_id
from
  databricks_sql_dashboard
where
  dashboard_filters_enabled;
```

### List dashboards that are in draft

```sql
select
  id,
  name,
  created_at,
  slug,
  account_id
from
  databricks_sql_dashboard
where
  is_draft;
```

### List dashboards modified in the past 7 days

```sql
select
  id,
  name,
  updated_at,
  slug,
  account_id
from
  databricks_sql_dashboard
where
  updated_at > now() - interval '7' day;
```

### List dashboards you can manage

```sql
select
  id,
  name,
  created_at,
  slug,
  account_id
from
  databricks_sql_dashboard
where
  permission_tier = 'CAN_MANAGE';
```

### Get widgets associated to each dashboard

```sql
select
  id,
  name,
  w ->> 'id' as widget_id,
  w -> 'options' ->> 'created_at' as widget_created_at,
  w -> 'options' ->> 'is_hidden' as widget_is_hidden,
  w -> 'options' ->> 'position' as widget_position,
  w -> 'options' ->> 'updated_at' as widget_updated_at,
  account_id
from
  databricks_sql_dashboard,
  jsonb_array_elements(widgets) as w;
```