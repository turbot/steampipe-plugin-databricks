---
title: "Steampipe Table: databricks_catalog_schema - Query Databricks Catalog Schemas using SQL"
description: "Allows users to query Databricks Catalog Schemas, specifically the database name, schema name, and schema owner, providing insights into the organization and ownership of data within a Databricks workspace."
---

# Table: databricks_catalog_schema - Query Databricks Catalog Schemas using SQL

Databricks Catalog is a feature within Databricks that organizes data into databases and tables. It provides a unified view of all data in Databricks and allows users to manage, discover, and utilize data effectively. A Databricks Catalog Schema is a logical grouping of tables within a database, providing a way to organize and manage data within a Databricks workspace.

## Table Usage Guide

The `databricks_catalog_schema` table provides insights into the organization and ownership of data within a Databricks workspace. As a data engineer or data scientist, explore schema-specific details through this table, including the database name, schema name, and schema owner. Utilize it to understand the structure of your data, discover who owns which schemas, and manage your data more effectively.

## Examples

### Basic info
Explore the basic details of your Databricks catalog schemas, such as who created them and when, to gain insights into their usage and management. This could be particularly useful for auditing purposes or for understanding the distribution of responsibility within your team.

```sql+postgres
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

```sql+sqlite
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
Gain insights into recent schema modifications by identifying those that have been updated in the past week. This can be useful for tracking changes, auditing purposes, or troubleshooting recent issues.

```sql+postgres
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

```sql+sqlite
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
  updated_at >= datetime('now', '-7 days');
```

### List system created schemas
Explore which schemas have been created by the system to gain insights into the organization and management of your data. This can be particularly useful for understanding the structure of your data and identifying areas for optimization.

```sql+postgres
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

```sql+sqlite
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
Explore which schemas have automatic maintenance enabled to streamline management and ensure optimal performance. This can be useful in identifying areas for potential optimization and troubleshooting.

```sql+postgres
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

```sql+sqlite
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
  enable_auto_maintenance = 1;
```

### Get effective permissions for each external location
Analyze the settings to understand the effective permissions assigned to each external location. This can help in managing access control and maintaining security protocols within your system.

```sql+postgres
select
  name,
  p ->> 'principal' as principal_name,
  p ->> 'privileges' as permissions
from
  databricks_catalog_schema,
  jsonb_array_elements(schema_effective_permissions) p;
```

```sql+sqlite
select
  name,
  json_extract(p.value, '$.principal') as principal_name,
  json_extract(p.value, '$.privileges') as permissions
from
  databricks_catalog_schema,
  json_each(schema_effective_permissions) as p;
```

### List catalog types and the average number of schemas per catalog
Explore the different types of catalogs and understand the average number of schemas each type typically contains. This can help in managing and optimizing the distribution of schemas across various catalogs.

```sql+postgres
select
  catalog_schema_counts.catalog_type,
  avg(catalog_schema_counts.schema_count) as avg_schemas_per_catalog
from (
  select
    c.catalog_type,
    count(s.full_name) as schema_count
  from
    databricks_catalog as c
    left join databricks_catalog_schema as s on c.name = s.catalog_name
  group by
    c.catalog_type
) as catalog_schema_counts
group by
  catalog_schema_counts.catalog_type;
```

```sql+sqlite
select
  catalog_schema_counts.catalog_type,
  avg(catalog_schema_counts.schema_count) as avg_schemas_per_catalog
from (
  select
    c.catalog_type,
    count(s.full_name) as schema_count
  from
    databricks_catalog as c
    left join databricks_catalog_schema as s on c.name = s.catalog_name
  group by
    c.catalog_type
) as catalog_schema_counts
group by
  catalog_schema_counts.catalog_type;
```