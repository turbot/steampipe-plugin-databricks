---
title: "Steampipe Table: databricks_catalog_volume - Query Databricks Catalog Volumes using SQL"
description: "Allows users to query Databricks Catalog Volumes, providing detailed information about the volumes in the Databricks catalog."
---

# Table: databricks_catalog_volume - Query Databricks Catalog Volumes using SQL

A Databricks Catalog Volume is a logical storage unit within Databricks, a unified analytics platform. These volumes are used to organize and manage data within the Databricks environment. They can contain a variety of data types and formats, including tables, views, and functions.

## Table Usage Guide

The `databricks_catalog_volume` table provides insights into the Catalog Volumes within Databricks. As a data engineer or data analyst, you can explore volume-specific details through this table, including volume type, properties, and associated metadata. Utilize it to uncover information about volumes, such as their organization, the types of data they contain, and their properties.

## Examples

### Basic info
Gain insights into specific volumes within your Databricks catalog, pinpointing details such as their creation date and creator. This information can be useful for auditing purposes or to understand the history and ownership of different volumes.

```sql+postgres
select
  volume_id,
  name,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog_volume
where
  catalog_name = 'catalog'
  and schema_name = 'schema';
```

```sql+sqlite
The provided PostgreSQL query does not contain any PostgreSQL-specific functions or data types that need to be converted to SQLite. Therefore, the SQLite equivalent of the provided query is exactly the same as the original query.

select
  volume_id,
  name,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog_volume
where
  catalog_name = 'catalog'
  and schema_name = 'schema';
```

### List volumes not modified in the last 90 days
Determine the volumes that have remained unmodified for the past 90 days. This query can help in identifying potential areas for data cleanup or archiving, thus optimizing your data storage and management.

```sql+postgres
select
  volume_id,
  name,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog_volume
where
  updated_at <= now() - interval '90 days'
  and catalog_name = 'catalog'
  and schema_name = 'schema';
```

```sql+sqlite
select
  volume_id,
  name,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog_volume
where
  updated_at <= datetime('now', '-90 days')
  and catalog_name = 'catalog'
  and schema_name = 'schema';
```

### List volumes created in the last 7 days
Explore volumes that have been created within the past week. This is useful for tracking recent additions and understanding the growth of your data storage.

```sql+postgres
select
  volume_id,
  name,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog_volume
where
  created_at >= now() - interval '7 days'
  and catalog_name = 'catalog'
  and schema_name = 'schema';
```

```sql+sqlite
select
  volume_id,
  name,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog_volume
where
  created_at >= datetime('now', '-7 days')
  and catalog_name = 'catalog'
  and schema_name = 'schema';
```

### List all external volumes
Discover the segments that utilize external volumes in your Databricks catalog. This can be beneficial in understanding your data storage and management strategy.

```sql+postgres
select
  volume_id,
  name,
  volume_type,
  storage_location,
  account_id
from
  databricks_catalog_volume
where
  volume_type = 'EXTERNAL'
  and catalog_name = 'catalog'
  and schema_name = 'schema';
```

```sql+sqlite
select
  volume_id,
  name,
  volume_type,
  storage_location,
  account_id
from
  databricks_catalog_volume
where
  volume_type = 'EXTERNAL'
  and catalog_name = 'catalog'
  and schema_name = 'schema';
```

### Get details for a specific volume
Explore specific volume details to understand its creation timeline, associated account, and metadata store. This is useful for auditing and tracking changes to the volume over time.

```sql+postgres
select
  volume_id,
  name,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog_volume
where
  full_name = '__catalog_name__.__schema_name__.__volume_name__';
```

```sql+sqlite
select
  volume_id,
  name,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog_volume
where
  full_name = '__catalog_name__.__schema_name__.__volume_name__';
```

### Count the number of volumes in a particular catalog
Explore the quantity of volumes within a specific catalog to understand its size and complexity. This is beneficial in managing and optimizing data storage and accessibility.

```sql+postgres
select
  catalog_name,
  count(*) as volume_count
from
  databricks_catalog_volume
where
  catalog_name = 'catalog'
  and schema_name = 'schema'
group by
  catalog_name;
```

```sql+sqlite
select
  catalog_name,
  count(*) as volume_count
from
  databricks_catalog_volume
where
  catalog_name = 'catalog'
  and schema_name = 'schema'
group by
  catalog_name;
```