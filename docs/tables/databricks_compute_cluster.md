# Table: databricks_compute_cluster

A Databricks cluster is a set of computation resources and configurations on which you run data engineering, data science, and data analytics workloads, such as production ETL pipelines, streaming analytics, ad-hoc analytics, and machine learning.

**Note**: Databricks retains cluster configuration information for up to 200 all-purpose clusters terminated in the last 30 days and up to 30 job clusters recently terminated by the job scheduler

## Examples

### Basic info

```sql
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

```sql
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

### List clusters that are running

```sql
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

```sql
select
  cluster_id,
  cluster_name,
  creator_user_name,
  autoscale ->> 'min_workers' as min_workers,
  autoscale ->> 'max_workers' as max_workers
from
  databricks_compute_cluster;
```

### Get aws attributes associated with each cluster

```sql
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

### Get azure attributes associated with each cluster

```sql
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

### Get azure attributes associated with each cluster

```sql
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

### List clusters that support port forwarding
  
```sql
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
  port_forwarding ->> 'enabled' = 'true';
```

### List clusters terminated due to inactivity

```sql
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
  state = 'TERMINATED' and
  termination_reason ->> 'code' = 'INACTIVITY';
```

### Get the permissions associated to each cluster

```sql
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