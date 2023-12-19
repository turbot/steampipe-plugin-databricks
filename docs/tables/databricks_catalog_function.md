---
title: "Steampipe Table: databricks_catalog_function - Query Databricks Catalog Functions using SQL"
description: "Allows users to query Databricks Catalog Functions, specifically providing insights into function names, databases, descriptions, and class names."
---

# Table: databricks_catalog_function - Query Databricks Catalog Functions using SQL

Databricks Catalog Functions are a part of Databricks' Unified Data Service, which provides a unified, collaborative workspace for data teams to build data pipelines, explore data, and perform machine learning tasks. Catalog Functions allow users to create, manage, and invoke functions that can be used in SQL expressions. These functions are stored in databases, which can be shared across multiple workspaces.

## Table Usage Guide

The `databricks_catalog_function` table provides insights into Catalog Functions within Databricks Unified Data Service. As a data engineer or data scientist, explore function-specific details through this table, including function names, associated databases, descriptions, and class names. Utilize it to uncover information about functions, such as their usage in SQL expressions, their storage in databases, and their shareability across multiple workspaces.

## Examples

### Basic info
Explore which functions have been created in a specific catalog and schema in Databricks, along with their details such as the creator and creation date. This can be useful for auditing, tracking changes, or understanding the structure of your data.

```sql+postgres
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

```sql+sqlite
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
Explore which functions have been updated in the past week. This can be useful in tracking recent changes and maintaining an understanding of ongoing modifications to your system.

```sql+postgres
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

```sql+sqlite
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
  updated_at >= datetime('now', '-7 days')
  and catalog_name = 'catalog'
  and schema_name = 'schema';
```

### List all scalar functions
Gain insights into the scalar functions present in your Databricks catalog to understand their creation, modification, and the users involved. This aids in managing and auditing the functions in your catalog.

```sql+postgres
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

```sql+sqlite
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
Explore which functions within your Databricks catalog are deterministic, allowing you to understand the functions that will always produce the same results given the same input values. This is useful in maintaining consistency and predictability in your data processing and analysis tasks.

```sql+postgres
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

```sql+sqlite
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
Explore the range of SQL functions within a specific catalog and schema. This is useful for understanding what functions are available and who created or updated them, providing a clearer view of your database's functionality and history.

```sql+postgres
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

```sql+sqlite
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

### List all external functions
Discover the segments that utilize external functions within your Databricks catalog. This can be useful in understanding the dependencies and interactions within your data schema, aiding in efficient data management and optimization.

```sql+postgres
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

```sql+sqlite
select
  function_id,
  name,
  routine_body,
  routine_definition,
  json_extract(routine_dependencies, '$.function') as routine_dependency_function,
  json_extract(routine_dependencies, '$.table') as routine_dependency_table
from
  databricks_catalog_function
where
  routine_body = 'EXTERNAL'
  and catalog_name = 'catalog'
  and schema_name = 'schema';
```

### List all functions that reads sql data
Explore which functions within a databricks catalog have access to read SQL data. This can be especially useful to identify potential data access points and maintain data security.

```sql+postgres
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

```sql+sqlite
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
Assess the elements within your Databricks catalog to gain insights into the effective permissions allocated to each function. This is useful to ensure correct access rights and maintain security within your system.

```sql+postgres
select
  name,
  p ->> 'principal' as principal_name,
  p ->> 'privileges' as permissions
from
  databricks_catalog_function,
  jsonb_array_elements(function_effective_permissions) p
where
  catalog_name = 'catalog'
  and schema_name = 'schema';
```

```sql+sqlite
select
  name,
  json_extract(p.value, '$.principal') as principal_name,
  json_extract(p.value, '$.privileges') as permissions
from
  databricks_catalog_function,
  json_each(function_effective_permissions) as p
where
  catalog_name = 'catalog'
  and schema_name = 'schema';
```

### Get details for a specific function
Explore the specifics of a particular function in your Databricks catalog, including its creator and creation date. This is useful for auditing purposes or understanding the history and purpose of a function in your data pipeline.

```sql+postgres
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

```sql+sqlite
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