---
title: "Steampipe Table: databricks_sql_warehouse - Query Databricks SQL Warehouses using SQL"
description: "Allows users to query Databricks SQL Warehouses, providing insights into their configurations and operational status."
---

# Table: databricks_sql_warehouse - Query Databricks SQL Warehouses using SQL

Databricks SQL Warehouse is a resource within Databricks that offers a high-performance, fully managed, cloud-based SQL analytics service. It enables users to execute SQL queries directly on their data lake and view the results in a familiar, spreadsheet-like format. Databricks SQL Warehouse is designed to provide low-latency SQL analytics over large datasets.

## Table Usage Guide

The `databricks_sql_warehouse` table offers insights into SQL Warehouses within Databricks. As a data analyst or data engineer, explore warehouse-specific details through this table, including configurations, operational status, and associated metadata. Utilize it to monitor and manage your SQL warehouses, such as understanding their current state, capacity, and performance.

## Examples

### Basic info
Explore the size and creator of each Databricks SQL warehouse in your account to understand its usage and allocation. This can assist in resource management and identifying any unusual activities.

```sql+postgres
select
  id,
  name,
  cluster_size,
  creator_name,
  jdbc_url,
  account_id
from
  databricks_sql_warehouse;
```

```sql+sqlite
select
  id,
  name,
  cluster_size,
  creator_name,
  jdbc_url,
  account_id
from
  databricks_sql_warehouse;
```

### Get cluster configuration for each warehouse
Analyze the configuration of each warehouse to understand the size of the cluster, the serverless compute status, and the number of active sessions. This can be useful in optimizing resource allocation and managing workload distribution across your data warehouses.

```sql+postgres
select
  id,
  name,
  cluster_size,
  enable_serverless_compute,
  max_num_clusters,
  min_num_clusters,
  num_active_sessions,
  num_clusters,
  spot_instance_policy,
  account_id
from
  databricks_sql_warehouse;
```

```sql+sqlite
select
  id,
  name,
  cluster_size,
  enable_serverless_compute,
  max_num_clusters,
  min_num_clusters,
  num_active_sessions,
  num_clusters,
  spot_instance_policy,
  account_id
from
  databricks_sql_warehouse;
```

### List all stopped warehouse objects
Explore which warehouse objects have been halted to manage resources efficiently and optimize the performance of your Databricks SQL warehouse. This could be particularly useful in identifying potential cost savings or troubleshooting performance issues.

```sql+postgres
select
  id,
  name,
  cluster_size,
  creator_name,
  jdbc_url,
  account_id
from
  databricks_sql_warehouse
where
  state = 'STOPPED';
```

```sql+sqlite
select
  id,
  name,
  cluster_size,
  creator_name,
  jdbc_url,
  account_id
from
  databricks_sql_warehouse
where
  state = 'STOPPED';
```

### List warehouses that have serverless compute enabled
Explore which warehouses have the serverless compute feature enabled to optimize resource allocation and improve operational efficiency. This can help in identifying potential cost savings and performance improvements.

```sql+postgres
select
  id,
  name,
  cluster_size,
  creator_name,
  jdbc_url,
  account_id
from
  databricks_sql_warehouse
where
  enable_serverless_compute;
```

```sql+sqlite
select
  id,
  name,
  cluster_size,
  creator_name,
  jdbc_url,
  account_id
from
  databricks_sql_warehouse
where
  enable_serverless_compute = 1;
```

### List warehouses with multiple ctive sessions
Determine the areas in which multiple active sessions are occurring within your warehouses. This enables effective resource management and helps in identifying potential areas of congestion.

```sql+postgres
select
  id,
  name,
  cluster_size,
  num_clusters
  creator_name,
  jdbc_url,
  account_id
from
  databricks_sql_warehouse
where
  num_active_sessions > 1;
```

```sql+sqlite
select
  id,
  name,
  cluster_size,
  num_clusters,
  creator_name,
  jdbc_url,
  account_id
from
  databricks_sql_warehouse
where
  num_active_sessions > 1;
```

### List warehouses that use photon optimized clusters
Explore which warehouses are utilizing photon-optimized clusters to better manage resources and improve overall performance. This can be particularly useful in identifying areas for potential optimization and cost-efficiency.

```sql+postgres
select
  id,
  name,
  cluster_size,
  creator_name,
  jdbc_url,
  account_id
from
  databricks_sql_warehouse
where
  enable_photon;
```

```sql+sqlite
select
  id,
  name,
  cluster_size,
  creator_name,
  jdbc_url,
  account_id
from
  databricks_sql_warehouse
where
  enable_photon = 1;
```

### List unhealthy warehouse objects
Discover the segments that are experiencing issues within your data warehouse. This query is useful for quickly identifying problematic areas for immediate attention and resolution.

```sql+postgres
select
  id,
  name,
  health ->> 'details' as health_details,
  health ->> 'failure_reason' as failure_reason,
  health ->> 'status' as health_status,
  health ->> 'summary' as health_summary,
  account_id
from
  databricks_sql_warehouse
where
  health ->> 'status' in ('FAILED', 'DEGRADED');
```

```sql+sqlite
select
  id,
  name,
  json_extract(health, '$.details') as health_details,
  json_extract(health, '$.failure_reason') as failure_reason,
  json_extract(health, '$.status') as health_status,
  json_extract(health, '$.summary') as health_summary,
  account_id
from
  databricks_sql_warehouse
where
  json_extract(health, '$.status') in ('FAILED', 'DEGRADED');
```

### Get count of warehouse objects by type
Discover the segments that have varying counts of warehouse objects, enabling you to understand the distribution and frequency of different types of warehouse objects within your Databricks SQL Warehouse. This could be useful for managing resources and optimizing warehouse operations.

```sql+postgres
select
  warehouse_type,
  count(*) as count
from
  databricks_sql_warehouse
group by
  warehouse_type;
```

```sql+sqlite
select
  warehouse_type,
  count(*) as count
from
  databricks_sql_warehouse
group by
  warehouse_type;
```

### Get warehouse odbc parameters
Explore which warehouses have specific ODBC parameters to help streamline your data management and ensure optimal database performance. This is useful in identifying any inconsistencies or issues in your warehouse configurations.

```sql+postgres
select
  id,
  name,
  odbc_params ->> 'hostname' as odbc_hostname,
  odbc_params ->> 'port' as odbc_port,
  odbc_params ->> 'path' as odbc_path,
  odbc_params ->> 'protocol' as odbc_protocol,
  account_id
from
  databricks_sql_warehouse;
```

```sql+sqlite
select
  id,
  name,
  json_extract(odbc_params, '$.hostname') as odbc_hostname,
  json_extract(odbc_params, '$.port') as odbc_port,
  json_extract(odbc_params, '$.path') as odbc_path,
  json_extract(odbc_params, '$.protocol') as odbc_protocol,
  account_id
from
  databricks_sql_warehouse;
```

### Get the permissions associated to each warehouse
Gain insights into the level of permissions assigned to each warehouse, identifying the principal user and group associated with each. This can help in managing access control and ensuring appropriate permissions are allocated.

```sql+postgres
select
  id,
  name,
  acl ->> 'user_name' as principal_user_name,
  acl ->> 'group_name' as principal_group_name,
  acl ->> 'all_permissions' as permission_level
from
  databricks_sql_warehouse,
  jsonb_array_elements(warehouse_permissions -> 'access_control_list') as acl;
```

```sql+sqlite
select
  id,
  name,
  json_extract(acl.value, '$.user_name') as principal_user_name,
  json_extract(acl.value, '$.group_name') as principal_group_name,
  json_extract(acl.value, '$.all_permissions') as permission_level
from
  databricks_sql_warehouse,
  json_each(warehouse_permissions, '$.access_control_list') as acl;
```