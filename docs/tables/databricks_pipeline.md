# Table: databricks_pipeline

Delta Live Tables is a framework for building reliable, maintainable, and testable data processing pipelines. You define the transformations to perform on your data, and Delta Live Tables manages task orchestration, cluster management, monitoring, data quality, and error handling.

## Examples

### Basic info

```sql
select
  pipeline_id,
  name,
  cluster_id,
  creator_user_name,
  state,
  edition,
  account_id
from
  databricks_pipeline;
```

### List pipelines that failed to start

```sql
select
  pipeline_id,
  name,
  cluster_id,
  creator_user_name,
  state,
  account_id
from
  databricks_pipeline
where
  state = 'FAILED';
```

### Get cluster details associated with each pipeline

```sql
select
  p.pipeline_id,
  p.name,
  p.cluster_id,
  c.cluster_name,
  c.cluster_source,
  c.cluster_cores,
  c.cluster_memory_mb,
  c.runtime_engine,
  c.account_id
from
  databricks_pipeline p,
  databricks_compute_cluster c
where
  p.cluster_id = c.cluster_id
  and p.account_id = c.account_id;
```

### Get the last completed pipeline update for each pipeline

```sql
select
  p.pipeline_id,
  p.name,
  p.state,
  u ->> 'creation_time' as update_creation_time,
  u ->> 'state' as update_state,
  u ->> 'update_id' as update_id,
  account_id
from
  databricks_pipeline p,
  jsonb_array_elements(p.latest_updates) as u
where
  u ->> 'state' = 'COMPLETED'
order by
  update_creation_time desc limit 1;
```

### Get the last failed pipeline update for each pipeline

```sql
select
  p.pipeline_id,
  p.name,
  p.state,
  u ->> 'creation_time' as update_creation_time,
  u ->> 'state' as update_state,
  u ->> 'update_id' as update_id,
  account_id
from
  databricks_pipeline p,
  jsonb_array_elements(p.latest_updates) as u
where
  u ->> 'state' = 'FAILED'
order by
  update_creation_time desc limit 1;
```

### Get pipelines publishing data in a catalog table

```sql
select
  pipeline_id,
  name,
  cluster_id,
  catalog,
  target,
  state,
  account_id
from
  databricks_pipeline
where
  catalog is not null;
```

### List pipelines that are manually triggered

```sql
select
  pipeline_id,
  name,
  cluster_id,
  creator_user_name,
  state,
  account_id
from
  databricks_pipeline
where
  not continuous;
```

### List unhealthy pipelines

```sql
select
  pipeline_id,
  name,
  cluster_id,
  creator_user_name,
  state,
  account_id
from
  databricks_pipeline
where
  health = 'UNHEALTHY';
```
### List pipelines in development mode

```sql
select
  pipeline_id,
  name,
  cluster_id,
  creator_user_name,
  state,
  account_id
from
  databricks_pipeline
where
  development;
```

### Get the permissions associated to each pipeline

```sql
select
  pipeline_id,
  name,
  acl ->> 'user_name' as principal_user_name,
  acl ->> 'group_name' as principal_group_name,
  acl ->> 'all_permissions' as permission_level
from
  databricks_pipeline,
  jsonb_array_elements(pipeline_permissions -> 'access_control_list') as acl;
```

### List libraries installed on each pipeline

```sql
select
  pipeline_id,
  name,
  l -> 'notebook' ->> 'path' as notebook_path,
  l ->> 'maven' as maven,
  l ->> 'whl' as whl,
  l ->> 'jar' as jar,
  l -> 'file' ->> 'path' as file_path,
  account_id
from
  databricks_pipeline,
  jsonb_array_elements(libraries) as l;
```

### Get trigger settings for each pipeline

```sql
select
  pipeline_id,
  name,
  trigger ->> 'cron' as cron,
  trigger ->> 'manual' as is_manual,
  account_id
from
  databricks_pipeline
where
  trigger is not null;
```

### Get cluster settings for each pipeline

```sql
select
  pipeline_id,
  name,
  c ->> 'instance_pool_id' as instance_pool_id,
  c ->> 'node_type_id' as node_type_id,
  c ->> 'autoscale' as autoscale,
  c ->> 'num_workers' as num_workers,
  c ->> 'policy_id' as policy_id,
  account_id
from
  databricks_pipeline,
  jsonb_array_elements(clusters) as c;
```