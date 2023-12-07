---
title: "Steampipe Table: databricks_catalog_table - Query Databricks Catalog Tables using SQL"
description: "Allows users to query Databricks Catalog Tables, providing detailed information about the tables in your Databricks workspaces."
---

# Table: databricks_catalog_table - Query Databricks Catalog Tables using SQL

Databricks Catalog is a feature in Databricks that allows users to discover, organize, and manage data across all Databricks workspaces. It provides a unified view of all tables, views, and functions in your Databricks environment. Databricks Catalog Tables are the specific tables within the catalog that hold your data.

## Table Usage Guide

The `databricks_catalog_table` table provides insights into the tables within your Databricks Catalog. As a Data Engineer or Data Scientist, explore table-specific details through this table, including table properties, associated databases, and table types. Utilize it to manage and organize your data effectively, ensuring optimal data discovery and usage within your Databricks environment.

## Examples

### Basic info
Explore the basic information of your Databricks catalog tables to understand their creation details and associated accounts. This can be useful for auditing, tracking changes, and managing access.

```sql+postgres
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

```sql+sqlite
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
Explore which tables have been updated in the past week. This is useful for keeping track of recent changes and modifications in your database.

```sql+postgres
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

```sql+sqlite
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
  updated_at >= datetime('now', '-7 days')
  and catalog_name = 'catalog'
  and schema_name = 'schema';
```

### List all view type tables
Determine the areas in which view type tables are used within your Databricks catalog. This is useful for understanding dependencies and managing your data architecture effectively.

```sql+postgres
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

```sql+sqlite
select
  table_id,
  name,
  view_definition,
  json_extract(view_dependencies, '$.function') as view_dependency_function,
  json_extract(view_dependencies, '$.table') as view_dependency_table,
  account_id
from
  databricks_catalog_table
where
  table_type = 'VIEW'
  and catalog_name = 'catalog'
  and schema_name = 'schema';
```

### List all tables with source as CSV
Explore which tables in your Databricks catalog are sourced from CSV files. This can be beneficial in instances where you need to identify and manage your data sources effectively.

```sql+postgres
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

```sql+sqlite
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
Discover the segments that are storing data externally in your Databricks catalog. This is useful for understanding where your data is located and how it's being accessed, especially for security and data management purposes.

```sql+postgres
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

```sql+sqlite
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
Discover the segments that have specific constraints applied to them. This is useful for assessing the integrity and structure of your data, ensuring that relationships between tables are properly enforced.

```sql+postgres
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

```sql+sqlite
select
  table_id,
  name,
  json_extract(c.value, '$.foreign_key_constraint') as foreign_key_constraint,
  json_extract(c.value, '$.primary_key_constraint') as primary_key_constraint,
  json_extract(c.value, '$.named_table_constraint') as named_table_constraint,
  account_id
from
  databricks_catalog_table,
  json_each(databricks_catalog_table.table_constraints, '$.table_constraints') as c
where
  catalog_name = 'catalog'
  and schema_name = 'schema';
```

### Get effective permissions for each table
Explore which users have specific permissions for each table in a Databricks catalog, providing a comprehensive view of access rights within the system. This can be particularly useful in managing user access and maintaining security protocols.

```sql+postgres
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

```sql+sqlite
select
  name,
  json_extract(p.value, '$.principal') as principal_name,
  json_extract(p.value, '$.privileges') as permissions
from
  databricks_catalog_table,
  json_each(table_effective_permissions) as p
where
  catalog_name = 'catalog'
  and schema_name = 'schema';
```

### Get details for a specific table
Explore the specifics of a particular table to understand its origin, authorship, and associated comments. This can be useful for auditing purposes or to gain insights into the table's history and purpose.

```sql+postgres
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

```sql+sqlite
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
This query helps in identifying the source catalog details for a specific table in Databricks. It's useful for understanding the origin and ownership of data, thereby aiding in data governance and accountability.

```sql+postgres
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
  left join databricks_catalog as c on t.catalog_name = c.name
where
  full_name = '__catalog_name__.__schema_name__.__table_name__';
```

```sql+sqlite
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
  left join databricks_catalog as c on t.catalog_name = c.name
where
  full_name = '__catalog_name__.__schema_name__.__table_name__';
```