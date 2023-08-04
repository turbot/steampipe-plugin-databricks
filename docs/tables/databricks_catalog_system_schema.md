# Table: databricks_catalog_system_schema

A system schema is a schema that lives within the system catalog. A system schema may contain information about customer usage of Unity Catalog such as audit-logs, billing-logs, lineage information, etc.

## Examples

### Basic info

```sql
select
  metastore_id,
  schema,
  state,
  account_id
from
  databricks_catalog_system_schema;
```

### List all system schemas that are unavailable

```sql
select
  metastore_id,
  schema,
  state,
  account_id
from
  databricks_catalog_system_schema
where
  state = 'UNAVAILABLE';
```

### Give details of the parent merastore associated to a particular schema

```sql
select
  s.title as system_schema_name,
  m.metastore_id,
  m.name as metastore_name,
  m.created_at as metastore_create_time,
  m.owner as metastore_owner,
  m.account_id as metastore_account_id
from
  databricks_catalog_system_schema as s
  left join databricks_catalog_metastore as m on s.metastore_id = m.metastore_id
where
  s.title = 'operational_data';
```

### Find the account with the most schemas

```sql
select
  account_id,
  count(*) as schema_count
from
  databricks_catalog_system_schema
group by
  account_id
order by
  schema_count desc
limit 1;
```