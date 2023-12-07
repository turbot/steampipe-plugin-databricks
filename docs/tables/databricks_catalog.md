---
title: "Steampipe Table: databricks_catalog - Query Databricks Catalogs using SQL"
description: "Allows users to query Databricks Catalogs, specifically the details of all Databricks databases and their corresponding tables, providing insights into data organization and structure."
---

# Table: databricks_catalog - Query Databricks Catalogs using SQL

Databricks Catalog is a feature within Databricks that provides a unified view of all your Databricks data objects. It allows you to organize, manage, and access Databricks databases, tables, and functions. The Catalog helps in managing data objects across all workspaces and simplifies data discovery and access.

## Table Usage Guide

The `databricks_catalog` table provides insights into Databricks Catalogs within Databricks. As a Data Engineer or Data Analyst, you can explore catalog-specific details through this table, including the databases and their corresponding tables. Utilize it to uncover information about your Databricks data objects, such as their organization, structure, and accessibility across workspaces.

## Examples

### Basic info
Explore the basic information related to your Databricks catalog, such as who created it and when, to better understand its origin and usage. This can be useful in auditing or managing your Databricks resources.

```sql+postgres
select
  name,
  catalog_type,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog;
```

```sql+sqlite
select
  name,
  catalog_type,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog;
```

### List catalogs modified in the last 7 days
Explore recent modifications to catalogs in the past week. This is useful for keeping track of changes and updates to your data resources.

```sql+postgres
select
  name,
  catalog_type,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog
where
  updated_at >= now() - interval '7 days';
```

```sql+sqlite
select
  name,
  catalog_type,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog
where
  updated_at >= datetime('now','-7 days');
```

### List catalogs with auto maintenance enabled
Explore which catalogs have automatic maintenance enabled in Databricks to better manage and optimize your data storage and processing. This can help in ensuring data integrity and improved performance.

```sql+postgres
select
  name,
  catalog_type,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog
where
  enable_auto_maintenance = 'ENABLE';
```

```sql+sqlite
select
  name,
  catalog_type,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog
where
  enable_auto_maintenance = 'ENABLE';
```

### List catalogs that are not isolated to the workspace
Explore which catalogs are not isolated within your workspace. This is useful for identifying potential security risks or data sharing opportunities across different areas of your organization.

```sql+postgres
select
  name,
  catalog_type,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog
where
  isolation_mode = 'OPEN';
```

```sql+sqlite
select
  name,
  catalog_type,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog
where
  isolation_mode = 'OPEN';
```

### Get effective permissions for each catalog
Explore the various permissions assigned to each catalog to understand the access levels of different principals. This can be useful in managing and auditing access control within your Databricks environment.

```sql+postgres
select
  name,
  p ->> 'principal' as principal_name,
  p ->> 'privileges' as permissions
from
  databricks_catalog,
  jsonb_array_elements(catalog_effective_permissions) p;
```

```sql+sqlite
select
  name,
  json_extract(p.value, '$.principal') as principal_name,
  json_extract(p.value, '$.privileges') as permissions
from
  databricks_catalog,
  json_each(catalog_effective_permissions) as p;
```

### Get total catalogs of each type
Analyze the distribution of different catalog types in your Databricks environment. This can be useful to understand the variety of data sources and their prevalence, aiding in data management and strategic planning.

```sql+postgres
select
  catalog_type,
  count(*) as total_catalogs
from
  databricks_catalog
group by
  catalog_type;
```

```sql+sqlite
select
  catalog_type,
  count(*) as total_catalogs
from
  databricks_catalog
group by
  catalog_type;
```

### List the most recently updated catalog
Explore the most recent changes in your catalog. This query is useful in monitoring updates and ensuring the latest modifications are as expected.

```sql+postgres
select
  name,
  catalog_type,
  updated_at
from
  databricks_catalog
order by
  updated_at desc
limit 1;
```

```sql+sqlite
select
  name,
  catalog_type,
  updated_at
from
  databricks_catalog
order by
  updated_at desc
limit 1;
```

### Count the number of catalogs created by each user, including a percentage of their ownership
Discover the segments that have been created by each user and understand their relative contribution in terms of percentage ownership. This can provide valuable insights into user activity and resource utilization.

```sql+postgres
select
  owner,
  count(*) as total_catalogs,
  (count(*) * 100.0 / sum(count(*)) over ()) as ownership_percentage
from
  databricks_catalog
group by
  owner;
```

```sql+sqlite
select
  owner,
  count(*) as total_catalogs,
  (count(*) * 100.0 / (select count(*) from databricks_catalog)) as ownership_percentage
from
  databricks_catalog
group by
  owner;
```