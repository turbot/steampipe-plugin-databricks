# Table: databricks_serving_serving_endpoint

You can use a serving endpoint to serve models from the Databricks Model Registry or from Unity Catalog. Endpoints expose the underlying models as scalable REST API endpoints using serverless compute.

## Examples

### Basic info

```sql
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

### List endpoints modified in the last 7 days

```sql
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
  last_updated_timestamp > now() - interval '7' day;
```

### List endpoints you can manage
  
```sql
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

```sql
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

### List the served models for the endpoint to serve

```sql
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

### Get the traffic configuration associated with the serving endpoint

```sql
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

### Get the permissions associated to each endpoint

```sql
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