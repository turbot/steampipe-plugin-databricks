---
title: "Steampipe Table: databricks_sql_query - Query Databricks SQL Queries using SQL"
description: "Allows users to query Databricks SQL Queries, specifically the query text, execution status, and associated metadata, providing insights into SQL operations and potential issues."
---

# Table: databricks_sql_query - Query Databricks SQL Queries using SQL

Databricks SQL is a service within Databricks that provides a powerful, collaborative, and integrated environment for data exploration and visualization. It allows users to run SQL queries on their data in Databricks, and visualize the results. Databricks SQL provides a centralized way to manage and execute SQL queries, offering insights into query performance and data exploration.

## Table Usage Guide

The `databricks_sql_query` table provides insights into SQL queries executed in Databricks. As a data scientist or data engineer, explore query-specific details through this table, including query text, execution status, and associated metadata. Utilize it to uncover information about queries, such as those that are running slowly, the ones that have failed, and to verify the efficiency of your SQL operations.

## Examples

### Basic info
Explore the basic details of your Databricks SQL queries, such as when they were created and their descriptions, to better understand the queries you have in your account. This can be useful for managing and optimizing your database queries.

```sql+postgres
select
  id,
  name,
  created_at,
  description,
  query,
  account_id
from
  databricks_sql_query;
```

```sql+sqlite
select
  id,
  name,
  created_at,
  description,
  query,
  account_id
from
  databricks_sql_query;
```

### List queries modified in the past 7 days
Explore which queries have been updated recently to gain insights into changes and modifications made within the last week. This is particularly useful for monitoring activity and keeping track of alterations made to your queries.

```sql+postgres
select
  id,
  name,
  created_at,
  description,
  last_modified_by,
  updated_at,
  query,
  account_id
from
  databricks_sql_query
where
  updated_at > now() - interval '7' day;
```

```sql+sqlite
select
  id,
  name,
  created_at,
  description,
  last_modified_by,
  updated_at,
  query,
  account_id
from
  databricks_sql_query
where
  updated_at > datetime('now', '-7 day');
```

### List archived queries
Discover the segments that have archived queries to better manage and organize your databricks data. This is useful for maintaining a clean workspace and keeping track of old queries for potential future reference.

```sql+postgres
select
  id,
  name,
  created_at,
  description,
  query,
  account_id
from
  databricks_sql_query
where
  is_archived;
```

```sql+sqlite
select
  id,
  name,
  created_at,
  description,
  query,
  account_id
from
  databricks_sql_query
where
  is_archived;
```

### List queries marked as favourite
Explore which queries have been marked as favorite. This can help you quickly access frequently used or important queries, enhancing your efficiency and productivity.

```sql+postgres
select
  id,
  name,
  created_at,
  description,
  query,
  account_id
from
  databricks_sql_query
where
  is_favorite;
```

```sql+sqlite
select
  id,
  name,
  created_at,
  description,
  query,
  account_id
from
  databricks_sql_query
where
  is_favorite = 1;
```

### List queries that are in draft
Explore which queries are still in draft mode. This can help to manage and prioritize work by identifying incomplete queries that may still require attention or completion.

```sql+postgres
select
  id,
  name,
  created_at,
  description,
  query,
  account_id
from
  databricks_sql_query
where
  is_draft;
```

```sql+sqlite
select
  id,
  name,
  created_at,
  description,
  query,
  account_id
from
  databricks_sql_query
where
  is_draft;
```

### List queries that are safe from SQL injection
Explore which queries are safeguarded against SQL injection, allowing you to maintain a secure database environment. This is crucial in preventing unauthorized access or potential data breaches.

```sql+postgres
select
  id,
  name,
  created_at,
  description,
  query,
  account_id
from
  databricks_sql_query
where
  is_safe;
```

```sql+sqlite
select
  id,
  name,
  created_at,
  description,
  query,
  account_id
from
  databricks_sql_query
where
  is_safe;
```

### List queries that can be managed by you
Uncover the details of SQL queries that you have management access to. This can be useful in understanding and controlling the queries that you are responsible for within the Databricks environment.

```sql+postgres
select
  id,
  name,
  created_at,
  description,
  query,
  account_id
from
  databricks_sql_query
where
  permission_tier = 'CAN_MANAGE';
```

```sql+sqlite
select
  id,
  name,
  created_at,
  description,
  query,
  account_id
from
  databricks_sql_query
where
  permission_tier = 'CAN_MANAGE';
```

### List parameters associated with each query
Determine the various parameters linked with each database query to gain insights into their characteristics and values. This is useful for understanding the details of each query and its associated parameters, enhancing data management and query optimization.

```sql+postgres
select
  id,
  name,
  created_at,
  description,
  query,
  p ->> 'name' as parameter_name,
  p ->> 'type' as parameter_type,
  p ->> 'value' as parameter_value,
  p ->> 'title' as parameter_title,
  account_id
from
  databricks_sql_query,
  jsonb_array_elements(options -> 'parameters') as p;
```

```sql+sqlite
select
  id,
  name,
  created_at,
  description,
  query,
  json_extract(p.value, '$.name') as parameter_name,
  json_extract(p.value, '$.type') as parameter_type,
  json_extract(p.value, '$.value') as parameter_value,
  json_extract(p.value, '$.title') as parameter_title,
  account_id
from
  databricks_sql_query,
  json_each(options, '$.parameters') as p;
```

### List all queries that are not editable
Discover the segments that consist of queries which are unmodifiable. This is particularly useful in maintaining data integrity and preventing unauthorized changes.

```sql+postgres
select
  id,
  name,
  created_at,
  description,
  query,
  account_id
from
  databricks_sql_query
where
  not can_edit;
```

```sql+sqlite
select
  id,
  name,
  created_at,
  description,
  query,
  account_id
from
  databricks_sql_query
where
  can_edit = 0;
```

### List visualizations associated to the queries
This query is useful to gain insights into the relationship between queries and associated visualizations in your Databricks account. It helps identify which visualizations are linked to certain queries, providing a better understanding of data usage and representation.

```sql+postgres
select
  id,
  name,
  created_at,
  query,
  account_id,
  visualizations ->> 'CreatedAt' as visualization_create_time,
  visualizations ->> 'Id' as visualization_id,
  visualizations ->> 'Name' as visualization_name,
  visualizations ->> 'Type' as visualization_type,
  visualizations ->> 'UpdatedAt' as visualization_update_time
from
  databricks_sql_query
where
  visualizations is not null;
```

```sql+sqlite
select
  id,
  name,
  created_at,
  query,
  account_id,
  json_extract(visualizations, '$.CreatedAt') as visualization_create_time,
  json_extract(visualizations, '$.Id') as visualization_id,
  json_extract(visualizations, '$.Name') as visualization_name,
  json_extract(visualizations, '$.Type') as visualization_type,
  json_extract(visualizations, '$.UpdatedAt') as visualization_update_time
from
  databricks_sql_query
where
  visualizations is not null;
```