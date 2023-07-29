# Table: databricks_sql_query_history

Query definitions include the target SQL warehouse, query text, name, description, tags, parameters, and visualizations. This table contains the run history of all queries in the workspace.

## Examples

### Basic info

```sql
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

```sql
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

```sql
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

### List query history by order of duration

```sql
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

```sql
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

