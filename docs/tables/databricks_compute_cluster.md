---
title: "Steampipe Table: databricks_compute_cluster - Query Databricks Compute Clusters using SQL"
description: "Allows users to query Databricks Compute Clusters, providing detailed information about the configuration and status of each cluster."
---

# Table: databricks_compute_cluster - Query Databricks Compute Clusters using SQL

Databricks Compute Clusters are a resource within the Databricks service that allows users to run data analytics workloads on scalable, optimized hardware. They provide a flexible and powerful environment for running a wide range of analytics tasks, from data processing to machine learning. Databricks Compute Clusters can be easily scaled up or down, depending on the computational requirements of the workload.

## Table Usage Guide

The `databricks_compute_cluster` table provides insights into Compute Clusters within Databricks. As a data engineer or data scientist, explore cluster-specific details through this table, including configuration, status, and associated metadata. Utilize it to uncover information about clusters, such as their current state, the hardware configuration, and the version of Databricks Runtime they are running.

## Examples

### Basic info
Discover the segments that allow you to identify and analyze the state and start time of compute clusters in Databricks, along with the user who created them. This can be beneficial in understanding usage patterns and managing resources effectively.

```sql+postgres
select
  cluster_id,
  cluster_name,
  creator_user_name,
  start_time,
  state,
  account_id
from
  databricks_compute_cluster;
```

```sql+sqlite
select
  cluster_id,
  cluster_name,
  creator_user_name,
  start_time,
  state,
  account_id
from
  databricks_compute_cluster;
```

### List clusters having local disk encryption disabled
Explore which clusters lack local disk encryption, a crucial security feature. This query is useful for identifying potential security vulnerabilities within your system.

```sql+postgres
select
  cluster_id,
  cluster_name,
  creator_user_name,
  start_time,
  state,
  account_id
from
  databricks_compute_cluster
where
  not enable_local_disk_encryption;
```

```sql+sqlite
select
  cluster_id,
  cluster_name,
  creator_user_name,
  start_time,
  state,
  account_id
from
  databricks_compute_cluster
where
  enable_local_disk_encryption = 0;
```

### List clusters that are running
Discover the segments that are actively running within your Databricks clusters. This can assist in managing resources and identifying potential areas for optimization or troubleshooting.

```sql+postgres
select
  cluster_id,
  cluster_name,
  creator_user_name,
  start_time,
  state,
  account_id
from
  databricks_compute_cluster
where
  state = 'RUNNING';
```

```sql+sqlite
select
  cluster_id,
  cluster_name,
  creator_user_name,
  start_time,
  state,
  account_id
from
  databricks_compute_cluster
where
  state = 'RUNNING';
```

### Get autoscaling configuration for each cluster
Explore the autoscaling configuration of each computing cluster to understand the minimum and maximum number of workers assigned. This helps in assessing the scalability of your compute resources.

```sql+postgres
select
  cluster_id,
  cluster_name,
  creator_user_name,
  autoscale ->> 'min_workers' as min_workers,
  autoscale ->> 'max_workers' as max_workers
from
  databricks_compute_cluster;
```

```sql+sqlite
select
  cluster_id,
  cluster_name,
  creator_user_name,
  json_extract(autoscale, '$.min_workers') as min_workers,
  json_extract(autoscale, '$.max_workers') as max_workers
from
  databricks_compute_cluster;
```

### Get AWS attributes associated with each cluster
Explore which AWS attributes are associated with each Databricks compute cluster to optimize resource allocation and improve system availability. This can provide insights into how your clusters are configured and help identify potential areas for cost savings or performance improvements.

```sql+postgres
select
  cluster_id,
  cluster_name,
  creator_user_name,
  aws_attributes ->> 'availability' as availability,
  aws_attributes ->> 'ebs_volume_type' as ebs_volume_type,
  aws_attributes ->> 'ebs_volume_count' as ebs_volume_count,
  aws_attributes ->> 'ebs_volume_iops' as ebs_volume_iops,
  aws_attributes ->> 'ebs_volume_size' as ebs_volume_size,
  aws_attributes ->> 'ebs_volume_throughput' as ebs_volume_throughput,
  aws_attributes ->> 'first_on_demand' as first_on_demand,
  aws_attributes ->> 'instance_profile_arn' as instance_profile_arn,
  aws_attributes ->> 'spot_bid_price_percent' as spot_bid_price_percent,
  aws_attributes ->> 'zone_id' as zone_id
from
  databricks_compute_cluster
where
  aws_attributes is not null;
```

```sql+sqlite
select
  cluster_id,
  cluster_name,
  creator_user_name,
  json_extract(aws_attributes, '$.availability') as availability,
  json_extract(aws_attributes, '$.ebs_volume_type') as ebs_volume_type,
  json_extract(aws_attributes, '$.ebs_volume_count') as ebs_volume_count,
  json_extract(aws_attributes, '$.ebs_volume_iops') as ebs_volume_iops,
  json_extract(aws_attributes, '$.ebs_volume_size') as ebs_volume_size,
  json_extract(aws_attributes, '$.ebs_volume_throughput') as ebs_volume_throughput,
  json_extract(aws_attributes, '$.first_on_demand') as first_on_demand,
  json_extract(aws_attributes, '$.instance_profile_arn') as instance_profile_arn,
  json_extract(aws_attributes, '$.spot_bid_price_percent') as spot_bid_price_percent,
  json_extract(aws_attributes, '$.zone_id') as zone_id
from
  databricks_compute_cluster
where
  aws_attributes is not null;
```

### Get Azure attributes associated with each cluster
This query is useful for gaining insights into the Azure attributes linked with each cluster in your Databricks compute environment. It helps in understanding the availability, on-demand status, log analytics information, and maximum bid price for spot instances, assisting in better resource management and cost optimization.

```sql+postgres
select
  cluster_id,
  cluster_name,
  creator_user_name,
  azure_attributes ->> 'availability' as availability,
  azure_attributes ->> 'first_on_demand' as first_on_demand,
  azure_attributes -> 'log_analytics_info' ->> 'log_analytics_primary_key' as log_analytics_primary_key,
  azure_attributes -> 'log_analytics_info' ->> 'log_analytics_workspace_id' as log_analytics_workspace_id,
  azure_attributes ->> 'spot_bid_max_price' as spot_bid_max_price
from
  databricks_compute_cluster
where
  azure_attributes is not null;
```

```sql+sqlite
select
  cluster_id,
  cluster_name,
  creator_user_name,
  json_extract(azure_attributes, '$.availability') as availability,
  json_extract(azure_attributes, '$.first_on_demand') as first_on_demand,
  json_extract(azure_attributes, '$.log_analytics_info.log_analytics_primary_key') as log_analytics_primary_key,
  json_extract(azure_attributes, '$.log_analytics_info.log_analytics_workspace_id') as log_analytics_workspace_id,
  json_extract(azure_attributes, '$.spot_bid_max_price') as spot_bid_max_price
from
  databricks_compute_cluster
where
  azure_attributes is not null;
```

### Get GCP attributes associated with each cluster
Explore the specifics of each cluster in your Google Cloud Platform (GCP) to understand its availability, boot disk size, associated Google service account, and local SSD count. This is useful for managing resources and ensuring optimal configuration for your cloud operations.

```sql+postgres
select
  cluster_id,
  cluster_name,
  creator_user_name,
  gcp_attributes ->> 'availability' as availability,
  gcp_attributes ->> 'boot_disk_size' as boot_disk_size,
  gcp_attributes -> 'google_service_account' as google_service_account,
  gcp_attributes -> 'local_ssd_count' as local_ssd_count
from
  databricks_compute_cluster
where
  gcp_attributes is not null;
```

```sql+sqlite
select
  cluster_id,
  cluster_name,
  creator_user_name,
  json_extract(gcp_attributes, '$.availability') as availability,
  json_extract(gcp_attributes, '$.boot_disk_size') as boot_disk_size,
  json_extract(gcp_attributes, '$.google_service_account') as google_service_account,
  json_extract(gcp_attributes, '$.local_ssd_count') as local_ssd_count
from
  databricks_compute_cluster
where
  gcp_attributes is not null;
```

### List clusters terminated due to inactivity
Determine the areas in which clusters have been terminated due to inactivity. This is useful to manage resources efficiently and avoid unnecessary costs associated with idle clusters.

```sql+postgres
select
  cluster_id,
  cluster_name,
  creator_user_name,
  start_time,
  state,
  account_id
from
  databricks_compute_cluster
where
  state = 'TERMINATED'
  and termination_reason ->> 'code' = 'INACTIVITY';
```

```sql+sqlite
select
  cluster_id,
  cluster_name,
  creator_user_name,
  start_time,
  state,
  account_id
from
  databricks_compute_cluster
where
  state = 'TERMINATED'
  and json_extract(termination_reason, '$.code') = 'INACTIVITY';
```

### Get the permissions associated to each cluster
Explore the access levels of different users and groups across each computing cluster. This can be useful for auditing security measures and ensuring appropriate access rights are in place.

```sql+postgres
select
  cluster_id,
  cluster_name,
  acl ->> 'user_name' as principal_user_name,
  acl ->> 'group_name' as principal_group_name,
  acl ->> 'all_permissions' as permission_level
from
  databricks_compute_cluster,
  jsonb_array_elements(cluster_permissions -> 'access_control_list') as acl;
```

```sql+sqlite
select
  cluster_id,
  cluster_name,
  json_extract(acl.value, '$.user_name') as principal_user_name,
  json_extract(acl.value, '$.group_name') as principal_group_name,
  json_extract(acl.value, '$.all_permissions') as permission_level
from
  databricks_compute_cluster,
  json_each(cluster_permissions, '$.access_control_list') as acl;
```

### List clusters having elastic disks enabled
Determine the areas in which elastic disks are enabled within your clusters. This can help optimize storage usage and enhance the performance of your data operations.

```sql+postgres
select
  cluster_id,
  cluster_name,
  creator_user_name,
  start_time,
  state,
  account_id
from
  databricks_compute_cluster
where
  enable_elastic_disk;
```

```sql+sqlite
select
  cluster_id,
  cluster_name,
  creator_user_name,
  start_time,
  state,
  account_id
from
  databricks_compute_cluster
where
  enable_elastic_disk = 1;
```

### List clusters that will automatically terminate within 1 hour of inactivity
Explore clusters set to automatically terminate after an hour of inactivity, which can be useful for managing resources and preventing unnecessary costs. This query helps in identifying such clusters and can guide decisions on resource allocation and usage.

```sql+postgres
select
  cluster_id,
  cluster_name,
  creator_user_name,
  start_time,
  state,
  account_id
from
  databricks_compute_cluster
where
  autotermination_minutes < 60;
```

```sql+sqlite
select
  cluster_id,
  cluster_name,
  creator_user_name,
  start_time,
  state,
  account_id
from
  databricks_compute_cluster
where
  autotermination_minutes < 60;
```

### List clusters created by the Databricks Jobs Scheduler
Explore which clusters have been created by the Databricks Jobs Scheduler. This can help in understanding the distribution of resources and determining if there are any idle or underutilized clusters, enabling more efficient resource management.

```sql+postgres
select
  cluster_id,
  cluster_name,
  creator_user_name,
  start_time,
  state,
  account_id
from
  databricks_compute_cluster
where
  cluster_source = 'JOB';
```

```sql+sqlite
select
  cluster_id,
  cluster_name,
  creator_user_name,
  start_time,
  state,
  account_id
from
  databricks_compute_cluster
where
  cluster_source = 'JOB';
```

### List clusters that can be accessed by a single user
Determine the areas in which a single user has access to multiple clusters. This aids in understanding the distribution of resources and permissions across different users in a Databricks environment.

```sql+postgres
select
  cluster_id,
  cluster_name,
  creator_user_name,
  start_time,
  state,
  account_id
from
  databricks_compute_cluster
where
  data_security_mode = 'SINGLE_USER';
```

```sql+sqlite
select
  cluster_id,
  cluster_name,
  creator_user_name,
  start_time,
  state,
  account_id
from
  databricks_compute_cluster
where
  data_security_mode = 'SINGLE_USER';
```