---
title: "Steampipe Table: databricks_sql_query_history - Query Databricks SQL Query History using SQL"
description: "Allows users to query Databricks SQL Query History, providing insights into past SQL queries executed in a Databricks workspace."
---

# Table: databricks_sql_query_history - Query Databricks SQL Query History using SQL

Databricks SQL Query History is a feature of Databricks that keeps a record of the SQL queries executed in a Databricks workspace. It provides details about each query such as the query text, execution time, user who ran the query, and more. This information can be useful for auditing, performance tuning, and understanding the usage patterns of the Databricks workspace.

## Table Usage Guide

The `databricks_sql_query_history` table provides insights into the SQL queries executed within a Databricks workspace. As a data analyst or data engineer, you can explore the details of past queries through this table, including the query text, execution time, user who ran the query, and more. Utilize it to audit the usage of the Databricks workspace, tune the performance of your SQL queries, and understand the usage patterns of your team.

## Examples

### Basic info
Analyze your Databricks SQL query history to gain insights into the user activity and performance. This can help you identify trends, optimize resource usage, and improve overall system efficiency.

```sql+postgres
select
  query_id,
  warehouse_id,
  executed_as_user_name,
  query_text,
  rows_produced,
  account_id
from
  databricks_sql_query_history;
```

```sql+sqlite
select
  query_id,
  warehouse_id,
  executed_as_user_name,
  query_text,
  rows_produced,
  account_id
from
  databricks_sql_query_history;
```

### List all DML queries
Explore which data modifications have been made in your Databricks warehouse. This is useful for tracking changes and understanding the impact of various insertions, updates, and deletions within your data.

```sql+postgres
select
  query_id,
  warehouse_id,
  executed_as_user_name,
  query_text,
  rows_produced,
  statement_type,
  account_id
from
  databricks_sql_query_history
where
  statement_type in ('INSERT', 'UPDATE', 'DELETE');
```

```sql+sqlite
select
  query_id,
  warehouse_id,
  executed_as_user_name,
  query_text,
  rows_produced,
  statement_type,
  account_id
from
  databricks_sql_query_history
where
  statement_type in ('INSERT', 'UPDATE', 'DELETE');
```

### List all failed queries
Analyze your databricks history to pinpoint specific instances where queries have failed. This can help you identify common issues and improve the overall efficiency of your database operations.

```sql+postgres
select
  query_id,
  query_text,
  error_message,
  account_id
from
  databricks_sql_query_history
where
  status = 'FAILED';
```

```sql+sqlite
select
  query_id,
  query_text,
  error_message,
  account_id
from
  databricks_sql_query_history
where
  status = 'FAILED';
```

### List queries with no more expected updates
Explore the list of queries that have completed their updates to get insights into any potential errors or issues. This is useful for identifying and troubleshooting problematic queries within your Databricks account.

```sql+postgres
select
  query_id,
  query_text,
  error_message,
  account_id
from
  databricks_sql_query_history
where
  is_final;
```

```sql+sqlite
select
  query_id,
  query_text,
  error_message,
  account_id
from
  databricks_sql_query_history
where
  is_final;
```

### List queries that have existing plans
Discover the segments that have existing plans in your Databricks SQL query history. This can be useful for identifying potential errors or inefficiencies in your queries.

```sql+postgres
select
  query_id,
  query_text,
  error_message,
  account_id
from
  databricks_sql_query_history
where
  plans_state = 'EXISTS';
```

```sql+sqlite
select
  query_id,
  query_text,
  error_message,
  account_id
from
  databricks_sql_query_history
where
  plans_state = 'EXISTS';
```

### List query history by order of duration
Analyze the history of executed queries to understand which ones took the most time to complete. This can help in pinpointing inefficient queries that may need optimization for better performance.

```sql+postgres
select
  query_id,
  rows_produced,
  duration,
  executed_as_user_name,
  query_text,
  account_id
from
  databricks_sql_query_history
order by
  duration desc;
```

```sql+sqlite
select
  query_id,
  rows_produced,
  duration,
  executed_as_user_name,
  query_text,
  account_id
from
  databricks_sql_query_history
order by
  duration desc;
```

### List metrics for each query execution
Determine the performance of each executed query by assessing factors such as compilation time, execution time, and total time taken. This can help in identifying inefficient queries and optimizing them for better performance.

```sql+postgres
select
  query_id,
  metrics ->> 'compilation_time_ms' as compilation_time_ms,
  metrics ->> 'execution_time_ms' as execution_time_ms,
  metrics ->> 'network_sent_bytes' as network_sent_bytes,
  metrics ->> 'read_bytes' as read_bytes,
  metrics ->> 'result_fetch_time_ms' as result_fetch_time_ms,
  metrics ->> 'result_from_cache' as result_from_cache,
  metrics ->> 'rows_read_count' as rows_read_count,
  metrics ->> 'total_time_ms' as total_time_ms
from
  databricks_sql_query_history;
```

```sql+sqlite
select
  query_id,
  json_extract(metrics, '$.compilation_time_ms') as compilation_time_ms,
  json_extract(metrics, '$.execution_time_ms') as execution_time_ms,
  json_extract(metrics, '$.network_sent_bytes') as network_sent_bytes,
  json_extract(metrics, '$.read_bytes') as read_bytes,
  json_extract(metrics, '$.result_fetch_time_ms') as result_fetch_time_ms,
  json_extract(metrics, '$.result_from_cache') as result_from_cache,
  json_extract(metrics, '$.rows_read_count') as rows_read_count,
  json_extract(metrics, '$.total_time_ms') as total_time_ms
from
  databricks_sql_query_history;
```