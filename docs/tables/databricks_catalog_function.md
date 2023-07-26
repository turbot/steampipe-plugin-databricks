# Table: databricks_catalog_function

Functions implement User-Defined Functions (UDFs) in Unity Catalog. The function implementation can be any SQL expression or Query, and it can be invoked wherever a table reference is allowed in a query.

The `databricks_catalog_function` table can be used to query information about any function, and **you must specify the catalog name and schema name** in the where or join clause using the `catalog_name` and `schema_name` columns.

## Examples

### Basic info

```sql
select
  function_id,
  name,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog_function
where
  catalog_name = 'catalog'
  and schema_name = 'schema';
```

### List functions modified in the last 7 days

```sql
select
  function_id,
  name,
  comment,
  created_at,
  created_by,
  updated_by,
  account_id
from
  databricks_catalog_function
where
  updated_at >= now() - interval '7 days'
  and catalog_name = 'catalog'
  and schema_name = 'schema';
```

### List all scalar functions

```sql
select
  function_id,
  name,
  comment,
  created_at,
  created_by,
  updated_by,
  data_type
from
  databricks_catalog_function
where
  data_type is not null
  and catalog_name = 'catalog'
  and schema_name = 'schema';
```

### List all deterministic functions

```sql
select
  function_id,
  name,
  comment,
  created_at,
  created_by,
  updated_by,
  account_id
from
  databricks_catalog_function
where
  is_deterministic
  and catalog_name = 'catalog'
  and schema_name = 'schema';
```

### List all SQL functions

```sql
select
  function_id,
  name,
  comment,
  created_at,
  created_by,
  updated_by,
  account_id
from
  databricks_catalog_function
where
  routine_body = 'SQL'
  and catalog_name = 'catalog'
  and schema_name = 'schema';
```

### List all EXTERNAL functions

```sql
select
  function_id,
  name,
  routine_body,
  routine_definition,
  routine_dependencies ->> 'function' as routine_dependency_function,
  routine_dependencies ->> 'table' as routine_dependency_table
from
  databricks_catalog_function
where
  routine_body = 'EXTERNAL'
  and catalog_name = 'catalog'
  and schema_name = 'schema';
```

### List all functions that reads sql data
  
```sql
select
  function_id,
  name,
  sql_data_access,
  sql_path
from
  databricks_catalog_function
where
  sql_data_access = 'READS_SQL_DATA'
  and catalog_name = 'catalog'
  and schema_name = 'schema';
```

### Get effective permissions for each function

```sql
select
  name,
  p ->> 'principal' as principal_name,
  p ->> 'privileges' as permissions
from
  databricks_catalog_function,
  jsonb_array_elements(function_effective_permissions) p
where
  and catalog_name = 'catalog'
  and schema_name = 'schema';
```

### Get details for a specific function

```sql
select
  function_id,
  name,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog_function
where
  full_name = '__catalog_name__.__schema_name__.__table_name__';
```