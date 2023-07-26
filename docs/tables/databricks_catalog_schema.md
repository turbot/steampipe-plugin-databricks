# Table: databricks_catalog_schema

A schema (also called a database) is the second layer of Unity Catalog’s three-level namespace. A schema organizes tables, views and functions. 

## Examples

### Basic info

```sql
select
  full_name,
  name,
  catalog_name,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog_schema;
```

### List schemas modified in the last 7 days

```sql
select
  full_name,
  name,
  catalog_name,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from  
  databricks_catalog_schema
where
  updated_at >= now() - interval '7 days';
```

### List system created schemas

```sql
select
  full_name,
  name,
  catalog_name,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog_schema
where
  owner = 'System user';
```

### List schemas having auto maintenance enabled

```sql
select
  full_name,
  name,
  catalog_name,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog_schema
where
  enable_auto_maintenance;
```

### Get effective permissions for each external location

```sql
select
  name,
  p ->> 'principal' as principal_name,
  p ->> 'privileges' as permissions
from
  databricks_catalog_schema,
  jsonb_array_elements(schema_effective_permissions) p;
```
