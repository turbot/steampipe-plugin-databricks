---
title: "Steampipe Table: databricks_catalog_connection - Query Databricks Catalog Connections using SQL"
description: "Allows users to query Databricks Catalog Connections, providing details about the connections available in the Databricks catalog."
---

# Table: databricks_catalog_connection - Query Databricks Catalog Connections using SQL

Databricks Catalog Connections is a feature within Databricks that allows users to manage and utilize data connections within their Databricks workspace. It provides a centralized way to set up and manage connections to various data sources, including databases, data lakes, and other data storage systems. Databricks Catalog Connections helps users to streamline their data workflows and improve data accessibility and utilization.

## Table Usage Guide

The `databricks_catalog_connection` table provides insights into the catalog connections within Databricks. As a data engineer, explore connection-specific details through this table, including connection properties, type, and associated metadata. Utilize it to manage and monitor data connections, such as those with specific properties, the types of connections, and the verification of connection status.

## Examples

### Basic info
Explore the various connections within your Databricks catalog to understand their types, creators, and creation dates. This can be useful for auditing purposes or to assess the overall configuration of your Databricks environment.

```sql+postgres
select
  name,
  connection_id,
  comment,
  connection_type,
  created_at,
  created_by,
  full_name,
  metastore_id,
  account_id
from
  databricks_catalog_connection;
```

```sql+sqlite
select
  name,
  connection_id,
  comment,
  connection_type,
  created_at,
  created_by,
  full_name,
  metastore_id,
  account_id
from
  databricks_catalog_connection;
```

### List connections modified in the last 7 days
Explore which connections have been updated in the past week. This could be useful for monitoring recent changes and ensuring they align with your organization's data policies.

```sql+postgres
select
  name,
  connection_id,
  comment,
  connection_type,
  created_at,
  created_by,
  full_name,
  metastore_id,
  account_id
from
  databricks_catalog_connection
where
  updated_at >= now() - interval '7 days';
```

```sql+sqlite
select
  name,
  connection_id,
  comment,
  connection_type,
  created_at,
  created_by,
  full_name,
  metastore_id,
  account_id
from
  databricks_catalog_connection
where
  updated_at >= datetime('now', '-7 days');
```

### List read only connections
Discover the segments that have read-only access in your databricks catalog to better manage data security and access permissions. This can be particularly useful when auditing user access or refining access control policies.

```sql+postgres
select
  name,
  connection_id,
  comment,
  connection_type,
  created_at,
  created_by,
  full_name,
  metastore_id,
  account_id
from
  databricks_catalog_connection
where
  read_only;
```

```sql+sqlite
select
  name,
  connection_id,
  comment,
  connection_type,
  created_at,
  created_by,
  full_name,
  metastore_id,
  account_id
from
  databricks_catalog_connection
where
  read_only;
```

### List all postgres connections
Explore which Databricks catalog connections are specifically linked to PostgreSQL. This is useful for understanding and managing your PostgreSQL integrations within Databricks.

```sql+postgres
select
  name,
  connection_id,
  comment,
  connection_type,
  created_at,
  created_by,
  full_name,
  metastore_id,
  account_id
from
  databricks_catalog_connection
where
  connection_type = 'POSTGRESQL';
```

```sql+sqlite
select
  name,
  connection_id,
  comment,
  connection_type,
  created_at,
  created_by,
  full_name,
  metastore_id,
  account_id
from
  databricks_catalog_connection
where
  connection_type = 'POSTGRESQL';
```

### Count the number of connections per connection type
Analyze your databricks catalog to understand the distribution of different connection types. This can help optimize resource allocation by identifying which connection types are used most frequently.

```sql+postgres
select
  connection_type,
  count(*) as connection_count
from
  databricks_catalog_connection
group by
  connection_type;
```

```sql+sqlite
select
  connection_type,
  count(*) as connection_count
from
  databricks_catalog_connection
group by
  connection_type;
```

### Calculate the total number of connections per owner, sorted by owner's total connections
Assess the elements within your databricks catalog to understand the distribution of connections per owner. This is particularly useful in identifying which owners are utilizing the most connections, thereby aiding in resource allocation and management.

```sql+postgres
select
  owner,
  count(*) as total_connections
from
  databricks_catalog_connection
group by
  owner
order by
  total_connections desc;
```

```sql+sqlite
select
  owner,
  count(*) as total_connections
from
  databricks_catalog_connection
group by
  owner
order by
  total_connections desc;
```

### List connections with properties that have the highest number of key-value pairs
Analyze your connections to understand which ones have the most complex properties, helping you focus on the ones that may require more maintenance or are more likely to experience issues due to their complexity.

```sql+postgres
select
  name,
  connection_type,
  jsonb_object_keys(properties_kvpairs) as keys
from
  databricks_catalog_connection
order by
  array_length(array(select keys), 1) desc
limit 10;
```

```sql+sqlite
Error: SQLite does not support array functions like array_length and array.
```

### Get details of the metastore associated to a particular connection
Explore the specifics of a particular connection by identifying its associated metastore details. This can be beneficial in understanding the ownership, creation time, and account details of the metastore, providing a comprehensive view of the connection's metadata storage.

```sql+postgres
select
  c.name as connection_name,
  m.metastore_id,
  m.name as metastore_name,
  m.created_at as metastore_create_time,
  m.owner as metastore_owner,
  m.account_id as metastore_account_id
from
  databricks_catalog_connection as c
  left join databricks_catalog_metastore as m on c.metastore_id = m.metastore_id;
```

```sql+sqlite
select
  c.name as connection_name,
  m.metastore_id,
  m.name as metastore_name,
  m.created_at as metastore_create_time,
  m.owner as metastore_owner,
  m.account_id as metastore_account_id
from
  databricks_catalog_connection as c
  left join databricks_catalog_metastore as m on c.metastore_id = m.metastore_id;
```