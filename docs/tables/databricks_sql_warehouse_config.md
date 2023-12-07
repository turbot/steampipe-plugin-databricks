---
title: "Steampipe Table: databricks_sql_warehouse_config - Query Databricks SQL Warehouse Configurations using SQL"
description: "Allows users to query Databricks SQL Warehouse Configurations, providing detailed insights into the configuration parameters of the SQL warehouses."
---

# Table: databricks_sql_warehouse_config - Query Databricks SQL Warehouse Configurations using SQL

Databricks SQL Warehouse Configuration is a set of parameters that define the behavior and capacity of a SQL warehouse in Databricks. It determines the size, type, and other attributes of the warehouse, influencing its performance and cost. These configurations can be adjusted to optimize the warehouse for specific workloads or usage patterns.

## Table Usage Guide

The `databricks_sql_warehouse_config` table provides insights into the configuration parameters of SQL warehouses within Databricks. As a data engineer or DBA, explore the details of these configurations through this table, including size, type, and other attributes. Utilize it to monitor and optimize the performance and cost of your SQL warehouses, based on specific workloads or usage patterns.

## Examples

### Get warehouse sql configuration
Explore the security policies and configuration parameters of your SQL warehouse in Databricks. This is useful for assessing the current configuration and identifying any potential areas for optimization or security enhancement.

```sql+postgres
select
  security_policy,
  cp ->> 'key' as config_parameter_key,
  cp ->> 'value' as config_parameter_value,
  account_id
from
  databricks_sql_warehouse_config,
  jsonb_array_elements(sql_configuration_parameters -> 'configuration_pairs') as cp;
```

```sql+sqlite
select
  security_policy,
  json_extract(cp.value, '$.key') as config_parameter_key,
  json_extract(cp.value, '$.value') as config_parameter_value,
  account_id
from
  databricks_sql_warehouse_config,
  json_each(sql_configuration_parameters, '$.configuration_pairs') as cp;
```

### Check if warehouse config uses a security policy
Analyze the settings to understand if the warehouse configuration employs a security policy. This information is crucial to ensure the security and integrity of your data.

```sql+postgres
select
  security_policy,
  google_service_account,
  instance_profile_arn,
  account_id
from
  databricks_sql_warehouse_config
where
  security_policy <> 'NONE';
```

```sql+sqlite
select
  security_policy,
  google_service_account,
  instance_profile_arn,
  account_id
from
  databricks_sql_warehouse_config
where
  security_policy <> 'NONE';
```

### Get data acces configuration
Analyze the settings to understand the security policy and data access configurations within your Databricks SQL warehouse. This can be beneficial to maintain and review your data security standards across the account.

```sql+postgres
select
  security_policy,
  ac ->> 'key' as config_parameter_key,
  ac ->> 'value' as config_parameter_value,
  account_id
from
  databricks_sql_warehouse_config,
  jsonb_array_elements(data_access_config) as ac;
```

```sql+sqlite
select
  security_policy,
  json_extract(ac.value, '$.key') as config_parameter_key,
  json_extract(ac.value, '$.value') as config_parameter_value,
  account_id
from
  databricks_sql_warehouse_config,
  json_each(data_access_config) as ac;
```

### Get all enabled warehouse types for the workspace
Explore which warehouse types are currently active within a workspace. This information aids in understanding the resources available for data processing and analytics tasks.

```sql+postgres
select
  security_policy,
  wt ->> 'warehouse_type' as warehouse_type,
  wt ->> 'enabled' as enabled,
  account_id
from
  databricks_sql_warehouse_config,
  jsonb_array_elements(enabled_warehouse_types) as wt
where
  wt ->> 'enabled' = 'true';
```

```sql+sqlite
select
  security_policy,
  json_extract(wt.value, '$.warehouse_type') as warehouse_type,
  json_extract(wt.value, '$.enabled') as enabled,
  account_id
from
  databricks_sql_warehouse_config,
  json_each(enabled_warehouse_types) as wt
where
  json_extract(wt.value, '$.enabled') = 'true';
```

### Get details of thew instance profile used to pass IAM role to the cluster
Analyze the settings to understand the linkage between the Databricks SQL warehouse configuration and the compute instance profile. This can be particularly useful in assessing how IAM roles are passed to clusters for managing permissions and access control.

```sql+postgres
select
  c.google_service_account,
  i.iam_role_arn,
  i.is_meta_instance_profile,
  i.account_id
from
  databricks_sql_warehouse_config as c
  left join databricks_compute_instance_profile as i on c.instance_profile_arn = i.instance_profile_arn;
```

```sql+sqlite
select
  c.google_service_account,
  i.iam_role_arn,
  i.is_meta_instance_profile,
  i.account_id
from
  databricks_sql_warehouse_config as c
  left join databricks_compute_instance_profile as i on c.instance_profile_arn = i.instance_profile_arn;
```