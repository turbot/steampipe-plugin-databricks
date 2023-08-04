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

### List view only data sources

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

### List details of the associated warehouse for a particular data source

```sql
select
  d.id as data_source_id,
  d.name as data_source_name,
  w.id as warehouse_id,
  w.name as warehouse_name,
  w.cluster_size warehouse_cluster_size,
  w.creator_name as warehouse_creator_name,
  w.jdbc_url as warehouse_jdbc_url,
  w.account_id
from
  databricks_sql_data_source as d
  left join databricks_sql_warehouse as w on d.warehouse_id = w.id;
```