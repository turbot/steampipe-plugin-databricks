---
title: "Steampipe Table: databricks_compute_instance_pool - Query Databricks Compute Instance Pools using SQL"
description: "Allows users to query Databricks Compute Instance Pools, providing details about the instance pool configurations and statuses."
---

# Table: databricks_compute_instance_pool - Query Databricks Compute Instance Pools using SQL

A Databricks Compute Instance Pool is a group of pre-created instances that are ready to use for job execution or notebook execution, reducing the latency of waiting for instances to be created. This resource is useful for managing costs and performance in Databricks workspaces. It provides a way to manage the lifecycle of instances and control the costs associated with idle instances.

## Table Usage Guide

The `databricks_compute_instance_pool` table provides insights into Compute Instance Pools within Databricks. As a DevOps engineer, you can explore details about each instance pool, such as its configuration and status, through this table. Use it to manage and monitor the usage of instances, ensuring optimal cost and performance in your Databricks workspace.

## Examples

### Basic info
Explore which Databricks compute instance pools are currently active, along with their associated account and node type IDs. This can be used to manage resources and track usage within your Databricks environment.

```sql+postgres
select
  instance_pool_id,
  instance_pool_name,
  node_type_id,
  state,
  account_id
from
  databricks_compute_instance_pool;
```

```sql+sqlite
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
Explore the configuration of your instance pool to understand its capacity and idle settings. This can help optimize resource usage and manage costs effectively in your Databricks environment.

```sql+postgres
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

```sql+sqlite
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
Discover the segments that have halted instance pools in Databricks Compute to understand potential resource underutilization or cost savings opportunities.

```sql+postgres
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

```sql+sqlite
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
Uncover the details of the deployed instance pools in AWS and analyze the configurations to understand their availability, bid price percentage for spot instances, and the specific AWS zone they are located in. This could be useful in managing resources and optimizing cost in a cloud environment.

```sql+postgres
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

```sql+sqlite
select
  instance_pool_id,
  instance_pool_name,
  json_extract(aws_attributes, '$.availability') as availability,
  json_extract(aws_attributes, '$.spot_bid_price_percent') as spot_bid_price_percent,
  json_extract(aws_attributes, '$.zone_id') as aws_zone_id,
  account_id
from
  databricks_compute_instance_pool
where
  aws_attributes is not null;
```

### Get Azure configurations instance pools deployed in AWS
Explore which Azure configurations instance pools are deployed in AWS to assess their availability and maximum spot bid price. This is beneficial for understanding resource allocation and cost management within your cloud infrastructure.

```sql+postgres
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

```sql+sqlite
select
  instance_pool_id,
  instance_pool_name,
  json_extract(azure_attributes, '$.availability') as availability,
  json_extract(azure_attributes, '$.spot_bid_max_price') as spot_bid_max_price,
  account_id
from
  databricks_compute_instance_pool
where
  azure_attributes is not null;
```

### Get GCP configurations instance pools deployed in AWS
Explore which instance pools are deployed in Google Cloud Platform (GCP) through Databricks Compute. This is useful for assessing the distribution of resources and identifying areas for potential reallocation or optimization.

```sql+postgres
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

```sql+sqlite
select
  instance_pool_id,
  instance_pool_name,
  json_extract(gcp_attributes, '$.gcp_availability') as gcp_availability,
  json_extract(gcp_attributes, '$.local_ssd_count') as local_ssd_count,
  account_id
from
  databricks_compute_instance_pool
where
  gcp_attributes is not null;
```

### Get disc specifications for each instance pool
Explore the specifications of each disk within an instance pool to understand their performance and capacity. This can help in assessing the adequacy of your current resources and planning for future capacity requirements.

```sql+postgres
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

```sql+sqlite
select
  instance_pool_id,
  instance_pool_name,
  json_extract(disk_spec, '$.disk_count') as disk_count,
  json_extract(disk_spec, '$.disk_iops') as disk_iops,
  json_extract(disk_spec, '$.disk_size') as disk_size,
  json_extract(disk_spec, '$.disk_throughput') as disk_throughput,
  json_extract(disk_spec, '$.disk_type') as disk_type,
  account_id
from
  databricks_compute_instance_pool;
```

### Get fleet related settings to power the instance pool
Analyze the settings to understand the configuration of your instance pool, specifically focusing on fleet-related settings. This is beneficial to manage your resources effectively and optimize your instance pool usage.

```sql+postgres
select
  instance_pool_id,
  instance_pool_name,
  instance_pool_fleet_attributes ->> 'fleet_on_demand_option' as fleet_on_demand_option,
  instance_pool_fleet_attributes ->> 'fleet_spot_option' as fleet_spot_option,
  instance_pool_fleet_attributes ->> 'launch_template_overrides' as launch_template_overrides,
  account_id
from
  databricks_compute_instance_pool;
```

```sql+sqlite
select
  instance_pool_id,
  instance_pool_name,
  json_extract(instance_pool_fleet_attributes, '$.fleet_on_demand_option') as fleet_on_demand_option,
  json_extract(instance_pool_fleet_attributes, '$.fleet_spot_option') as fleet_spot_option,
  json_extract(instance_pool_fleet_attributes, '$.launch_template_overrides') as launch_template_overrides,
  account_id
from
  databricks_compute_instance_pool;
```

### Get all preloaded docker images for each instance pool
Explore which Docker images are preloaded for each instance pool in your Databricks compute environment. This can help you manage your resources more effectively and ensure the necessary tools are readily available for your data processing tasks.

```sql+postgres
select
  instance_pool_id,
  instance_pool_name,
  p ->> 'basic_auth' as docker_image_basic_auth,
  p ->> 'url' as docker_image_url,
  account_id
from
  databricks_compute_instance_pool,
  jsonb_array_elements(preloaded_docker_images) as p;
```

```sql+sqlite
select
  instance_pool_id,
  instance_pool_name,
  json_extract(p.value, '$.basic_auth') as docker_image_basic_auth,
  json_extract(p.value, '$.url') as docker_image_url,
  account_id
from
  databricks_compute_instance_pool,
  json_each(preloaded_docker_images) as p;
```

### Get stats for each instance pool
Explore the statistics of each instance pool to understand its usage pattern. This can help in optimizing resource allocation and identifying any potential bottlenecks or underutilized resources.

```sql+postgres
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

```sql+sqlite
select
  instance_pool_id,
  instance_pool_name,
  json_extract(stats, '$.idle_count') as idle_count,
  json_extract(stats, '$.pending_idle_count') as pending_idle_count,
  json_extract(stats, '$.pending_used_count') as pending_used_count,
  json_extract(stats, '$.used_count') as used_count,
  account_id
from
  databricks_compute_instance_pool;
```

### Get the permissions associated to each instance pool
Explore the level of access granted to each user and group within your instance pool. This can be beneficial in managing security and access control, ensuring only appropriate permissions are granted.

```sql+postgres
select
  instance_pool_id,
  instance_pool_name,
  acl ->> 'user_name' as principal_user_name,
  acl ->> 'group_name' as principal_group_name,
  acl ->> 'all_permissions' as permission_level
from
  databricks_compute_instance_pool,
  jsonb_array_elements(instance_pool_permission -> 'access_control_list') as acl;
```

```sql+sqlite
select
  instance_pool_id,
  instance_pool_name,
  json_extract(acl.value, '$.user_name') as principal_user_name,
  json_extract(acl.value, '$.group_name') as principal_group_name,
  json_extract(acl.value, '$.all_permissions') as permission_level
from
  databricks_compute_instance_pool,
  json_each(instance_pool_permission, '$.access_control_list') as acl;
```

### List instance pools capable of autoscaling local storage
Determine the areas in which instance pools are capable of autoscaling local storage. This helps in understanding the flexibility and scalability of storage resources in your compute environment.

```sql+postgres
select
  instance_pool_id,
  instance_pool_name,
  idle_instance_autotermination_minutes,
  max_capacity,
  min_idle_instances,
  node_type_id
from
  databricks_compute_instance_pool
where
  enable_elastic_disk;
```

```sql+sqlite
select
  instance_pool_id,
  instance_pool_name,
  idle_instance_autotermination_minutes,
  max_capacity,
  min_idle_instances,
  node_type_id
from
  databricks_compute_instance_pool
where
  enable_elastic_disk = 1;
```

### List instance pools having no pending instance errors
Explore the instance pools that are functioning smoothly without any pending error instances. This is beneficial for monitoring the health and efficiency of your databricks compute infrastructure.

```sql+postgres
select
  instance_pool_id,
  instance_pool_name,
  idle_instance_autotermination_minutes,
  max_capacity,
  min_idle_instances,
  node_type_id
from
  databricks_compute_instance_pool
where
  jsonb_array_length(pending_instance_errors) = 0;
```

```sql+sqlite
select
  instance_pool_id,
  instance_pool_name,
  idle_instance_autotermination_minutes,
  max_capacity,
  min_idle_instances,
  node_type_id
from
  databricks_compute_instance_pool
where
  json_array_length(pending_instance_errors) = 0;
```