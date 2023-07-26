# Table: databricks_catalog_connection

A connection is an abstraction of an external data source that can be connected from Databricks Compute.

## Examples

### Basic info

```sql
select
  name,
  connection_id,
  comment,
  connection_type,
  created_at,
  created_by,
  full_name,
  metastore_id,
  account_id
from
  databricks_catalog_connection;
```

### List connections modified in the last 7 days

```sql
select
  name,
  connection_id,
  comment,
  connection_type,
  created_at,
  created_by,
  full_name,
  metastore_id,
  account_id
from  
  databricks_catalog_connection
where
  updated_at >= now() - interval '7 days';
```

### List read only connections

```sql
select
  name,
  connection_id,
  comment,
  connection_type,
  created_at,
  created_by,
  full_name,
  metastore_id,
  account_id
from  
  databricks_catalog_connection
where
  read_only;
```

### List all postgres connections

```sql
select
  name,
  connection_id,
  comment,
  connection_type,
  created_at,
  created_by,
  full_name,
  metastore_id,
  account_id
from
  databricks_catalog_connection
where
  connection_type = 'POSTGRESQL';
```
