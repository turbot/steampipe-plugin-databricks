# Table: databricks_sql_data_source

When creating a query object, you may optionally specify a data_source_id for the SQL warehouse against which it will run.

## Examples

### Basic info

```sql
select
  id,
  name,
  syntax,
  type,
  warehouse_id,
  account_id
from
  databricks_sql_data_source;
```

### List all view only data sources

```sql
select
  id,
  name,
  syntax,
  type,
  warehouse_id,
  account_id
from
  databricks_sql_data_source
where
  view_only;
```

### List all paused data sources

```sql
select
  id,
  name,
  syntax,
  pause_reason,
  warehouse_id,
  account_id
from
  databricks_sql_data_source
where
  paused;
```

### List all data sources that support auto limit

```sql
select
  id,
  name,
  syntax,
  type,
  warehouse_id,
  account_id
from
  databricks_sql_data_source
where
  supports_auto_limit;
```