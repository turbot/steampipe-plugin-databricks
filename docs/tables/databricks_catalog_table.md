# Table: databricks_catalog_table

A table resides in the third layer of Unity Catalog’s three-level namespace. It contains rows of data.

The `databricks_catalog_table` table can be used to query information about any table, and **you must specify the catalog name and schema name** in the where or join clause using the `catalog_name` and `schema_name` columns.

**Note** To query a table, users must have the **SELECT** permission on the table, and they must have the **USE_CATALOG** permission on its parent catalog and the **USE_SCHEMA** permission on its parent schema.

## Examples

### Basic info

```sql
select
  table_id,
  name,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog_table
where
  catalog_name = 'catalog'
  and schema_name = 'schema';
```

### List tables modified in the last 7 days

```sql
select
  table_id,
  name,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog_table
where
  updated_at >= now() - interval '7 days'
  and catalog_name = 'catalog'
  and schema_name = 'schema';
```

### List all view type tables

```sql
select
  table_id,
  name,
  view_definition,
  view_dependencies ->> 'function' as view_dependency_function,
  view_dependencies ->> 'table' as view_dependency_table,
  account_id
from
  databricks_catalog_table
where
  table_type = 'VIEW'
  and catalog_name = 'catalog'
  and schema_name = 'schema';
```

### List all tables with source as CSV

```sql
select
  table_id,
  name,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog_table
where
  data_source_format = 'CSV'
  and catalog_name = 'catalog'
  and schema_name = 'schema';
```

### List all external tables

```sql
select
  table_id,
  name,
  table_type,
  storage_location,
  storage_credential_name
from
  databricks_catalog_table
where
  table_type = 'EXTERNAL'
  and catalog_name = 'catalog'
  and schema_name = 'schema';
```

### Get all table constraints

```sql
select
  table_id,
  name,
  c ->> 'foreign_key_constraint' as foreign_key_constraint,
  c ->> 'primary_key_constraint' as primary_key_constraint,
  c ->> 'named_table_constraint' as named_table_constraint,
  account_id
from
  databricks_catalog_table,
  jsonb_array_elements(table_constraints -> 'table_constraints') as c
where
  catalog_name = 'catalog'
  and schema_name = 'schema';
```

### Get effective permissions for each table

```sql
select
  name,
  p ->> 'principal' as principal_name,
  p ->> 'privileges' as permissions
from
  databricks_catalog_table,
  jsonb_array_elements(table_effective_permissions) p
where
  catalog_name = 'catalog'
  and schema_name = 'schema';
```

### Get details for a specific table

```sql
select
  table_id,
  name,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog_table
where
  full_name = '__catalog_name__.__schema_name__.__table_name__';
```

### List details of the parent catalog for a particular table

```sql
select
  t.name as table_name,
  c.name as catalog_name,
  c.catalog_type,
  c.created_at as catalog_create_time,
  c.created_by as catalog_created_by,
  c.metastore_id,
  c.account_id
from
  databricks_catalog_table as t
  left join databricks_catalog_catalog as c on t.catalog_name = c.name
where
  full_name = '__catalog_name__.__schema_name__.__table_name__';
```