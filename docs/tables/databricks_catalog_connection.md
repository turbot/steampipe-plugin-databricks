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

### Count the number of connections per connection type

```sql
select
  connection_type,
  count(*) as connection_count
from
  databricks_catalog_connection
group by
  connection_type;
```

### Calculate the total number of connections per owner, sorted by owner's total connections

```sql
select
  owner,
  count(*) as total_connections
from
  databricks_catalog_connection
group by
  owner
order by
  total_connections desc;
```

### List connections with properties that have the highest number of key-value pairs

```sql
select
  name,
  connection_type,
  jsonb_object_keys(properties_kvpairs) as keys
from
  databricks_catalog_connection
order by
  array_length(array(select keys), 1) desc
limit 10;
```

### Get details of the metastore associated to a particular connection

```sql
select
  c.name as connection_name,
  m.metastore_id,
  m.name as metastore_name,
  m.created_at as metastore_create_time,
  m.owner as metastore_owner,
  m.account_id as metastore_account_id
from
  databricks_catalog_connection as c
  left join databricks_catalog_metastore as m on c.metastore_id = m.metastore_id;
```