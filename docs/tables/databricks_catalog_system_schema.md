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
