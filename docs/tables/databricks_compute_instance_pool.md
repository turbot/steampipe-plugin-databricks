# Table: databricks_compute_instance_pool

Instance Pools API are used to create, edit, delete and list instance pools by using ready-to-use cloud instances which reduces a cluster start and auto-scaling times.

## Examples

### Basic info

```sql
select
  instance_pool_id,
  instance_pool_name,
  node_type_id,
  state,
  account_id
from
  databricks_compute_instance_pool;
```

### Get instance pool configuration

```sql
select
  instance_pool_id,
  instance_pool_name,
  idle_instance_autotermination_minutes,
  max_capacity,
  min_idle_instances,
  node_type_id
from
  databricks_compute_instance_pool;
```

### List instance pools that are stopped

```sql
select
  instance_pool_id,
  instance_pool_name,
  node_type_id,
  state,
  account_id
from
  databricks_compute_instance_pool
where
  state = 'STOPPED';
```

### Get AWS configurations instance pools deployed in AWS

```sql
select
  instance_pool_id,
  instance_pool_name,
  aws_attributes ->> 'availability' as availability,
  aws_attributes ->> 'spot_bid_price_percent' as spot_bid_price_percent,
  aws_attributes ->> 'zone_id' as aws_zone_id,
  account_id
from
  databricks_compute_instance_pool
where
  aws_attributes is not null;
```

### Get Azure configurations instance pools deployed in AWS

```sql
select
  instance_pool_id,
  instance_pool_name,
  azure_attributes ->> 'availability' as availability,
  azure_attributes ->> 'spot_bid_max_price' as spot_bid_max_price,
  account_id
from
  databricks_compute_instance_pool
where
  azure_attributes is not null;
```

### Get GCP configurations instance pools deployed in AWS

```sql
select
  instance_pool_id,
  instance_pool_name,
  gcp_attributes ->> 'gcp_availability' as gcp_availability,
  gcp_attributes ->> 'local_ssd_count' as local_ssd_count,
  account_id
from
  databricks_compute_instance_pool
where
  gcp_attributes is not null;
```

### Get disc specifications for each instance pool

```sql
select
  instance_pool_id,
  instance_pool_name,
  disk_spec ->> 'disk_count' as disk_count,
  disk_spec ->> 'disk_iops' as disk_iops,
  disk_spec ->> 'disk_size' as disk_size,
  disk_spec ->> 'disk_throughput' as disk_throughput,
  disk_spec ->> 'disk_type' as disk_type,
  account_id
from
  databricks_compute_instance_pool;
```

### Get fleet related settings to power the instance pool

```sql
select
  instance_pool_id,
  instance_pool_name,
  fleet_attributes ->> 'fleet_on_demand_option' as fleet_on_demand_option,
  fleet_attributes ->> 'fleet_spot_option' as fleet_spot_option,
  fleet_attributes ->> 'launch_template_overrides' as launch_template_overrides,
  account_id
from
  databricks_compute_instance_pool;
```

### Get all preloaded docker images for each instance pool

```sql
select
  instance_pool_id,
  instance_pool_name,
  p ->> 'basic_auth' as docker_image_basic_auth,
  p ->> 'url' as docker_image_url,
  account_id
from
  databricks_compute_instance_pool
  jsonb_array_elements(preloaded_docker_images) as p;
```

### Get stats for each instance pool

```sql
select
  instance_pool_id,
  instance_pool_name,
  stats ->> 'idle_count' as idle_count,
  stats ->> 'pending_idle_count' as pending_idle_count,
  stats ->> 'pending_used_count' as pending_used_count,
  stats ->> 'used_count' as used_count,
  account_id
from
  databricks_compute_instance_pool;
```

### Get the permissions associated to each instance pool

```sql
select
  instance_pool_id,
  instance_pool_name,
  acl ->> 'user_name' as principal_user_name,
  acl ->> 'group_name' as principal_group_name,
  acl ->> 'all_permissions' as permission_level
from
  databricks_compute_instance_pool,
  jsonb_array_elements(instance_pool_permission -> 'access_control_list') as acl;
