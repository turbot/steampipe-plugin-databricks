---
title: "Steampipe Table: databricks_sql_dashboard - Query Databricks SQL Dashboards using SQL"
description: "Allows users to query SQL Dashboards in Databricks, providing insights into the configuration, status, and metadata of each dashboard."
---

# Table: databricks_sql_dashboard - Query Databricks SQL Dashboards using SQL

Databricks SQL Dashboards are a feature of Databricks that allows users to create, view, and share interactive dashboards with others. These dashboards can include visualizations, tables, and other elements that are based on SQL queries. They are a powerful tool for data exploration and sharing insights with others.

## Table Usage Guide

The `databricks_sql_dashboard` table provides insights into SQL Dashboards within Databricks. As a data analyst or data scientist, you can explore dashboard-specific details through this table, including configuration, status, and metadata. Use it to uncover information about dashboards, such as their configuration details, the number of views they have received, and the last time they were updated.

## Examples

### Basic info
Explore which Databricks SQL dashboards were created under specific accounts and when, to gain insights into usage patterns and account activity. This can help in understanding the distribution and management of resources across different accounts.

```sql+postgres
select
  id,
  name,
  created_at,
  slug,
  account_id
from
  databricks_sql_dashboard;
```

```sql+sqlite
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
Discover the dashboards that were created within the past week to stay updated on the latest data visualizations and insights. This is useful to track recent activity and progress in your data analysis tasks.

```sql+postgres
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

```sql+sqlite
select
  id,
  name,
  created_at,
  slug,
  account_id
from
  databricks_sql_dashboard
where
  created_at >= datetime('now', '-7 day');
```

### List dashboards that are editable
Explore which dashboards are editable to understand your account's configuration, helping you manage and modify dashboards as needed.

```sql+postgres
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

```sql+sqlite
select
  id,
  name,
  created_at,
  slug,
  account_id
from
  databricks_sql_dashboard
where
  can_edit = 1;
```

### List dashboards that are archived
Determine the areas in which dashboards have been archived to review their status and manage resources effectively within your Databricks account.

```sql+postgres
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

```sql+sqlite
select
  id,
  name,
  created_at,
  slug,
  account_id
from
  databricks_sql_dashboard
where
  is_archived = 1;
```

### List dashboards that are marked as favourite
Explore which dashboards have been marked as favorite, helping users quickly access their most frequently used or important data visualizations. This can streamline data analysis and improve efficiency.

```sql+postgres
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

```sql+sqlite
select
  id,
  name,
  created_at,
  slug,
  account_id
from
  databricks_sql_dashboard
where
  is_favorite = 1;
```

### List dashboards that have filters enabled
Explore which dashboards have enabled filters to enhance data visualization and analysis, thereby improving decision-making processes.

```sql+postgres
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

```sql+sqlite
select
  id,
  name,
  created_at,
  slug,
  account_id
from
  databricks_sql_dashboard
where
  dashboard_filters_enabled = 1;
```

### List dashboards that are in draft
Discover the segments that are still in draft within your Databricks SQL dashboard. This can be useful for identifying incomplete or under-review segments that need further development or approval.

```sql+postgres
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

```sql+sqlite
select
  id,
  name,
  created_at,
  slug,
  account_id
from
  databricks_sql_dashboard
where
  is_draft = 1;
```

### List dashboards modified in the past 7 days
Explore which dashboards have been updated in the past week. This is useful for keeping track of recent modifications and ensuring all changes are accounted for.

```sql+postgres
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

```sql+sqlite
select
  id,
  name,
  updated_at,
  slug,
  account_id
from
  databricks_sql_dashboard
where
  updated_at > datetime('now', '-7 day');
```

### List dashboards you can manage
Explore which dashboards you have management permissions for, allowing you to understand your level of control and access within the Databricks environment. This can be useful in assessing your ability to make changes or updates to these dashboards.

```sql+postgres
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

```sql+sqlite
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
Determine the areas in which widgets are linked to various dashboards. This helps in understanding how different widgets are utilized across different dashboards, providing insights into widget usage and placement trends.

```sql+postgres
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

```sql+sqlite
select
  id,
  name,
  json_extract(w.value, '$.id') as widget_id,
  json_extract(w.value, '$.options.created_at') as widget_created_at,
  json_extract(w.value, '$.options.is_hidden') as widget_is_hidden,
  json_extract(w.value, '$.options.position') as widget_position,
  json_extract(w.value, '$.options.updated_at') as widget_updated_at,
  account_id
from
  databricks_sql_dashboard,
  json_each(widgets) as w;
```