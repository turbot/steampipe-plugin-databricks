---
title: "Steampipe Table: databricks_catalog_metastore - Query Databricks Catalog Metastores using SQL"
description: "Allows users to query Databricks Catalog Metastores, specifically the metadata of databases and tables, providing insights into the structure and properties of the data."
---

# Table: databricks_catalog_metastore - Query Databricks Catalog Metastores using SQL

A Databricks Catalog Metastore is a critical component within Databricks that stores metadata of databases and tables. It provides a centralized way to manage and access metadata, which includes information about databases, tables, views, functions, and more. Databricks Catalog Metastore helps users to understand the structure and properties of their data.

## Table Usage Guide

The `databricks_catalog_metastore` table provides insights into the Databricks Catalog Metastore within Databricks. As a Data Engineer or Data Scientist, explore metadata-specific details through this table, including database names, table names, and associated properties. Use it to uncover information about the organization and structure of your data, helping you to manage and utilize your data more effectively.

## Examples

### Basic info
Explore the key details of your Databricks metastores to gain insights into their creation time, ownership, and associated cloud platform. This can help you manage and track your data resources effectively.

```sql+postgres
select
  metastore_id,
  name,
  cloud,
  created_at,
  owner,
  account_id
from
  databricks_catalog_metastore;
```

```sql+sqlite
select
  metastore_id,
  name,
  cloud,
  created_at,
  owner,
  account_id
from
  databricks_catalog_metastore;
```

### List cloud provider configuration for the metastores
Analyze the settings to understand the configuration of your cloud provider for the metastores. This can be useful to identify the regions, storage roots, and associated credentials, helping you manage and optimize your cloud resources effectively.

```sql+postgres
select
  metastore_id,
  name,
  cloud,
  global_metastore_id,
  region,
  storage_root,
  storage_root_credential_id,
  storage_root_credential_name
from
  databricks_catalog_metastore;
```

```sql+sqlite
select
  metastore_id,
  name,
  cloud,
  global_metastore_id,
  region,
  storage_root,
  storage_root_credential_id,
  storage_root_credential_name
from
  databricks_catalog_metastore;
```

### List metastores that could be shared externally
Identify metastores that have the potential to be shared externally. This query is beneficial in managing data access and ensuring appropriate security measures are in place.

```sql+postgres
select
  metastore_id,
  name,
  cloud,
  global_metastore_id,
  owner,
  account_id
from
  databricks_catalog_metastore
where
  delta_sharing_scope = 'INTERNAL_AND_EXTERNAL';
```

```sql+sqlite
select
  metastore_id,
  name,
  cloud,
  global_metastore_id,
  owner,
  account_id
from
  databricks_catalog_metastore
where
  delta_sharing_scope = 'INTERNAL_AND_EXTERNAL';
```

### List metastores that were updated in the last 7 days
Identify recently updated metastores in your Databricks catalog to keep track of changes. This can be useful for auditing and maintaining up-to-date information.

```sql+postgres
select
  metastore_id,
  name,
  cloud,
  owner,
  updated_at,
  account_id
from
  databricks_catalog_metastore
where
  updated_at >= now() - interval '7 days';
```

```sql+sqlite
select
  metastore_id,
  name,
  cloud,
  owner,
  updated_at,
  account_id
from
  databricks_catalog_metastore
where
  updated_at >= datetime('now', '-7 days');
```

### Get effective permissions for each function
Analyze the permissions associated with each function in your Databricks catalog to understand who has what level of access. This can help in maintaining security and managing access control effectively.

```sql+postgres
select
  metastore_id,
  name,
  p ->> 'principal' as principal_name,
  p ->> 'privileges' as permissions
from
  databricks_catalog_metastore,
  jsonb_array_elements(metastore_effective_permissions) p;
```

```sql+sqlite
select
  metastore_id,
  name,
  json_extract(p.value, '$.principal') as principal_name,
  json_extract(p.value, '$.privileges') as permissions
from
  databricks_catalog_metastore,
  json_each(metastore_effective_permissions) as p;
```

### List metastores with the highest number of effective permissions
Discover the metastores that have the most effective permissions. This is useful for analyzing which metastores are potentially more exposed or have more complex access configurations.

```sql+postgres
select
  name,
  cloud,
  jsonb_array_length(metastore_effective_permissions) as permission_count
from
  databricks_catalog_metastore
order by
  permission_count desc
limit 10;
```

```sql+sqlite
select
  name,
  cloud,
  json_array_length(metastore_effective_permissions) as permission_count
from
  databricks_catalog_metastore
order by
  permission_count desc
limit 10;
```

### Find the most recently updated metastores
Explore which metastores have been updated most recently to keep track of the latest changes and ensure data integrity. This is particularly useful in managing large datasets across multiple cloud platforms.

```sql+postgres
select
  name,
  cloud,
  updated_at
from
  databricks_catalog_metastore
order by
  updated_at desc
limit 10;
```

```sql+sqlite
select
  name,
  cloud,
  updated_at
from
  databricks_catalog_metastore
order by
  updated_at desc
limit 10;
```

### Count the number of metastores per cloud
This query is used to analyze the distribution of metastores across different cloud platforms. It helps in understanding the spread of data storage and management systems, providing insights for strategic decision-making in cloud resource allocation.

```sql+postgres
select
  cloud,
  count(*) as metastore_count
from
  databricks_catalog_metastore
group by
  cloud;
```

```sql+sqlite
select
  cloud,
  count(*) as metastore_count
from
  databricks_catalog_metastore
group by
  cloud;
```