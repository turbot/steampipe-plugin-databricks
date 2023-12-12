---
title: "Steampipe Table: databricks_serving_serving_endpoint - Query Databricks Serving Endpoints using SQL"
description: "Allows users to query Databricks Serving Endpoints, specifically providing details about each endpoint's configuration, status, and associated models."
---

# Table: databricks_serving_serving_endpoint - Query Databricks Serving Endpoints using SQL

Databricks Serving is a feature of Databricks that allows for the deployment and serving of machine learning models. Serving Endpoints, in particular, are the interfaces through which these models are accessed and utilized. They provide crucial information about the model's configuration, its current operational status, and any models associated with it.

## Table Usage Guide

The `databricks_serving_serving_endpoint` table offers insights into the Serving Endpoints within Databricks. As a Data Scientist or Machine Learning Engineer, explore endpoint-specific details through this table, including configuration, operational status, and associated models. Utilize it to monitor the deployment of your machine learning models, track their operational status, and manage their configurations.

## Examples

### Basic info
Explore which Databricks serving endpoints have been created, by whom, and when, to gain insights into the usage and permissions associated with your Databricks account. This allows for effective management and auditing of your resources.

```sql+postgres
select
  id,
  name,
  creation_timestamp,
  creator,
  permission_level,
  account_id
from
  databricks_serving_serving_endpoint;
```

```sql+sqlite
select
  id,
  name,
  creation_timestamp,
  creator,
  permission_level,
  account_id
from
  databricks_serving_serving_endpoint;
```

### List endpoints created in the last 7 days
Discover the recently created endpoints in your system to understand the changes made in the last week. This can help in monitoring activity, assessing resource allocation, and identifying any unexpected or unauthorized modifications.

```sql+postgres
select
  id,
  name,
  creation_timestamp,
  creator,
  permission_level,
  account_id
from
  databricks_serving_serving_endpoint
where
  creation_timestamp >= now() - interval '7' day;
```

```sql+sqlite
select
  id,
  name,
  creation_timestamp,
  creator,
  permission_level,
  account_id
from
  databricks_serving_serving_endpoint
where
  creation_timestamp >= datetime('now', '-7 day');
```

### List endpoints that have not been modified in the last 10 days
Discover the segments that have been static for the past 10 days by analyzing the settings of your endpoints. This could help in identifying potential areas for update or optimization.

```sql+postgres
select
  id,
  name,
  creation_timestamp,
  creator,
  permission_level,
  account_id
from
  databricks_serving_serving_endpoint
where
  last_updated_timestamp <= now() - interval '10' day;
```

```sql+sqlite
select
  id,
  name,
  creation_timestamp,
  creator,
  permission_level,
  account_id
from
  databricks_serving_serving_endpoint
where
  last_updated_timestamp <= datetime('now', '-10 day');
```

### List endpoints you can manage
Explore which endpoints you have management permissions for, to better understand your access levels and control over different resources. This can be useful for auditing purposes or for establishing more secure configurations.

```sql+postgres
select
  id,
  name,
  creation_timestamp,
  creator,
  permission_level,
  account_id
from
  databricks_serving_serving_endpoint
where
  permission_level = 'CAN_MANAGE';
```

```sql+sqlite
select
  id,
  name,
  creation_timestamp,
  creator,
  permission_level,
  account_id
from
  databricks_serving_serving_endpoint
where
  permission_level = 'CAN_MANAGE';
```

### List endpoints with failed config updates
Uncover the details of endpoints that have experienced failed configuration updates. This query is useful in identifying potential issues with your endpoints, allowing you to take corrective action promptly.

```sql+postgres
select
  id,
  name,
  pending_config ->> 'config_version' as pending_config_version,
  pending_config ->> 'served_models' as pending_served_models,
  pending_config ->> 'start_time' as update_start_time,
  pending_config ->> 'traffic_config' as pending_traffic_config
from
  databricks_serving_serving_endpoint
where
  state ->> 'config_update' = 'UPDATE_FAILED';
```

```sql+sqlite
select
  id,
  name,
  json_extract(pending_config, '$.config_version') as pending_config_version,
  json_extract(pending_config, '$.served_models') as pending_served_models,
  json_extract(pending_config, '$.start_time') as update_start_time,
  json_extract(pending_config, '$.traffic_config') as pending_traffic_config
from
  databricks_serving_serving_endpoint
where
  json_extract(state, '$.config_update') = 'UPDATE_FAILED';
```

### List the served models for the endpoint to serve
Discover the models being served by a particular endpoint, along with their versions and settings such as 'scale to zero' and 'workload size'. This can be useful for managing and optimizing model deployment in a machine learning environment.

```sql+postgres
select
  id,
  name,
  sm ->> 'model_name' as model_name,
  sm ->> 'model_version' as model_version,
  sm ->> 'scale_to_zero_enabled' as scale_to_zero_enabled,
  sm ->> 'workload_size' as workload_size,
  account_id
from
  databricks_serving_serving_endpoint,
  jsonb_array_elements(config -> 'served_models') as sm;
```

```sql+sqlite
select
  id,
  name,
  json_extract(sm.value, '$.model_name') as model_name,
  json_extract(sm.value, '$.model_version') as model_version,
  json_extract(sm.value, '$.scale_to_zero_enabled') as scale_to_zero_enabled,
  json_extract(sm.value, '$.workload_size') as workload_size,
  account_id
from
  databricks_serving_serving_endpoint,
  json_each(config, '$.served_models') as sm;
```

### Get the traffic configuration associated with the serving endpoint
Explore the distribution of traffic across different model-serving endpoints in your Databricks environment. This helps in understanding how your processing power is allocated and can be useful in optimizing resource utilization.

```sql+postgres
select
  id,
  name,
  r ->> 'served_model_name' as served_model_name,
  r ->> 'traffic_percentage' as traffic_percentage,
  account_id
from
  databricks_serving_serving_endpoint,
  jsonb_array_elements(config -> 'traffic_config' -> 'routes') as r
where
  config -> 'traffic_config' is not null;
```

```sql+sqlite
select
  id,
  name,
  json_extract(r.value, '$.served_model_name') as served_model_name,
  json_extract(r.value, '$.traffic_percentage') as traffic_percentage,
  account_id
from
  databricks_serving_serving_endpoint,
  json_each(json_extract(config, '$.traffic_config.routes')) as r
where
  json_extract(config, '$.traffic_config') is not null;
```

### Get the permissions associated to each endpoint
Explore the access levels for various users and groups across different endpoints. This can help in managing data security and ensuring appropriate access rights are in place.

```sql+postgres
select
  id,
  name,
  acl ->> 'user_name' as principal_user_name,
  acl ->> 'group_name' as principal_group_name,
  acl ->> 'all_permissions' as permission_level
from
  databricks_serving_serving_endpoint,
  jsonb_array_elements(serving_endpoint_permissions -> 'access_control_list') as acl;
```

```sql+sqlite
select
  id,
  name,
  json_extract(acl.value, '$.user_name') as principal_user_name,
  json_extract(acl.value, '$.group_name') as principal_group_name,
  json_extract(acl.value, '$.all_permissions') as permission_level
from
  databricks_serving_serving_endpoint,
  json_each(serving_endpoint_permissions, '$.access_control_list') as acl;
```