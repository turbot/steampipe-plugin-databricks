---
title: "Steampipe Table: databricks_sql_data_source - Query Databricks SQL Data Sources using SQL"
description: "Allows users to query Databricks SQL Data Sources, providing insights into the metadata and configurations of the data sources used in Databricks SQL analytics."
---

# Table: databricks_sql_data_source - Query Databricks SQL Data Sources using SQL

Databricks SQL Data Source is a component within Databricks that allows you to connect and interact with various data sources for SQL analytics. It provides a way to configure and manage data sources, enabling seamless integration with databases, data warehouses, and other data platforms. Databricks SQL Data Source helps you stay informed about the configurations and metadata of your data sources, ensuring efficient data analysis and management.

## Table Usage Guide

The `databricks_sql_data_source` table provides insights into SQL data sources within Databricks. As a Data Analyst or Data Engineer, explore data source-specific details through this table, including configurations, metadata, and associated properties. Utilize it to uncover information about data sources, such as their types, options, and the databases they are associated with.

## Examples

### Basic info
Explore which data sources are connected to your Databricks SQL account, allowing you to understand the structure and organization of your data. This can be useful in managing data access and optimizing data exploration.

```sql+postgres
select
  id,
  name,
  syntax,
  type,
  warehouse_id,
  account_id
from
  databricks_sql_data_source;
```

```sql+sqlite
select
  id,
  name,
  syntax,
  type,
  warehouse_id,
  account_id
from
  databricks_sql_data_source;
```

### List view only data sources
Explore which data sources in Databricks are set to 'view-only'. This can be useful for understanding access limitations and managing permissions within your data warehouse.

```sql+postgres
select
  id,
  name,
  syntax,
  type,
  warehouse_id,
  account_id
from
  databricks_sql_data_source
where
  view_only;
```

```sql+sqlite
select
  id,
  name,
  syntax,
  type,
  warehouse_id,
  account_id
from
  databricks_sql_data_source
where
  view_only = 1;
```

### List all paused data sources
Explore which data sources are currently paused in Databricks, providing a way to assess the elements within your system that may need attention or troubleshooting.

```sql+postgres
select
  id,
  name,
  syntax,
  pause_reason,
  warehouse_id,
  account_id
from
  databricks_sql_data_source
where
  paused;
```

```sql+sqlite
select
  id,
  name,
  syntax,
  pause_reason,
  warehouse_id,
  account_id
from
  databricks_sql_data_source
where
  paused = 1;
```

### List all data sources that support auto limit
Analyze the settings to understand which data sources automatically limit data retrieval, helping to optimize data management and prevent overwhelming the system with large data sets. This is useful in maintaining efficient system performance and data integrity.

```sql+postgres
select
  id,
  name,
  syntax,
  type,
  warehouse_id,
  account_id
from
  databricks_sql_data_source
where
  supports_auto_limit;
```

```sql+sqlite
select
  id,
  name,
  syntax,
  type,
  warehouse_id,
  account_id
from
  databricks_sql_data_source
where
  supports_auto_limit = 1;
```

### List details of the associated warehouse for a particular data source
This example helps you understand the relationship between a specific data source and its associated warehouse in Databricks SQL. It's useful for gaining insights into warehouse details like its size, creator, and JDBC URL, which can aid in resource management and optimization.

```sql+postgres
select
  d.id as data_source_id,
  d.name as data_source_name,
  w.id as warehouse_id,
  w.name as warehouse_name,
  w.cluster_size warehouse_cluster_size,
  w.creator_name as warehouse_creator_name,
  w.jdbc_url as warehouse_jdbc_url,
  w.account_id
from
  databricks_sql_data_source as d
  left join databricks_sql_warehouse as w on d.warehouse_id = w.id;
```

```sql+sqlite
select
  d.id as data_source_id,
  d.name as data_source_name,
  w.id as warehouse_id,
  w.name as warehouse_name,
  w.cluster_size warehouse_cluster_size,
  w.creator_name as warehouse_creator_name,
  w.jdbc_url as warehouse_jdbc_url,
  w.account_id
from
  databricks_sql_data_source as d
  left join databricks_sql_warehouse as w on d.warehouse_id = w.id;
```