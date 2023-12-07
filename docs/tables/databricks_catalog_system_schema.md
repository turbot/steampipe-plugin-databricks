---
title: "Steampipe Table: databricks_catalog_system_schema - Query Databricks Catalog System Schemas using SQL"
description: "Allows users to query Databricks Catalog System Schemas, specifically the system-level details and metadata of the Databricks catalog."
---

# Table: databricks_catalog_system_schema - Query Databricks Catalog System Schemas using SQL

A Databricks Catalog System Schema is a feature within Databricks that provides a unified interface for managing all data objects (tables, views, functions, etc.) and their respective schemas. It allows users to organize their data into databases for better management and querying. The system schema provides system-level details and metadata of the Databricks catalog.

## Table Usage Guide

The `databricks_catalog_system_schema` table provides insights into the system-level details and metadata of the Databricks catalog. As a Data Engineer or Data Scientist, explore specific details through this table, such as the database name, description, location URI, owner name, and more. Utilize it to uncover information about the structure and organization of your Databricks data objects, and to manage and query your data more effectively.

## Examples

### Basic info
Analyze the status of your Databricks system schemas to understand which ones are active within your account. This can help in effectively managing your Databricks resources and ensuring optimal performance.

```sql+postgres
select
  metastore_id,
  schema,
  state,
  account_id
from
  databricks_catalog_system_schema;
```

```sql+sqlite
select
  metastore_id,
  schema,
  state,
  account_id
from
  databricks_catalog_system_schema;
```

### List all system schemas that are unavailable
Uncover the details of system schemas that are currently unavailable in your Databricks catalog. This can help you identify potential issues with your data management and ensure smooth operation of your systems.

```sql+postgres
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

```sql+sqlite
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
Gain insights into the specific parent storage system associated with a particular schema in your Databricks catalog. This is useful in understanding the origin and ownership of data, aiding in data governance and management.

```sql+postgres
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

```sql+sqlite
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
Explore which account has the highest number of schemas, allowing you to identify the account with the most extensive system schema usage. This information can be particularly useful for resource allocation and system optimization efforts.

```sql+postgres
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

```sql+sqlite
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