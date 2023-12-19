---
title: "Steampipe Table: databricks_pipeline - Query Databricks Pipelines using SQL"
description: "Allows users to query Databricks Pipelines, providing insights into the configuration, status, and metadata of each pipeline."
---

# Table: databricks_pipeline - Query Databricks Pipelines using SQL

Databricks Pipelines are a set of tools within the Databricks platform that allows you to build, test, and deploy machine learning workflows. It provides a unified way to automate the end-to-end machine learning lifecycle, including feature extraction, model training and testing, and model deployment. Databricks Pipelines help you manage and streamline your machine learning workflows, ensuring reproducibility and facilitating collaboration.

## Table Usage Guide

The `databricks_pipeline` table provides insights into Pipelines within Databricks. As a Data Scientist or Machine Learning Engineer, explore pipeline-specific details through this table, including configuration, status, and associated metadata. Utilize it to manage and streamline your machine learning workflows, ensuring reproducibility and facilitating collaboration.

## Examples

### Basic info
Explore which pipelines are active in your Databricks environment, identifying their associated cluster and creator. This can help in managing resources and understanding user activity.

```sql+postgres
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

```sql+sqlite
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
Identify instances where certain pipelines have failed to initiate. This provides insights into potential issues within your system, allowing you to troubleshoot and resolve these problems effectively.

```sql+postgres
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

```sql+sqlite
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
Analyze the settings to understand the relationship between different pipelines and their associated clusters in Databricks. This can help in resource allocation and optimization by providing insights into the computational resources each pipeline utilizes.

```sql+postgres
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

```sql+sqlite
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
This query allows you to gain insights into the most recent updates completed for each pipeline in your Databricks environment. It can be useful for monitoring pipeline performance, identifying potential issues, and understanding the timeline of your data processing tasks.

```sql+postgres
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

```sql+sqlite
select
  p.pipeline_id,
  p.name,
  p.state,
  json_extract(u.value, '$.creation_time') as update_creation_time,
  json_extract(u.value, '$.state') as update_state,
  json_extract(u.value, '$.update_id') as update_id,
  p.account_id
from
  databricks_pipeline p,
  json_each(p.latest_updates) as u
where
  json_extract(u.value, '$.state') = 'COMPLETED'
order by
  update_creation_time desc limit 1;
```

### Get the last failed pipeline update for each pipeline
This example demonstrates how to pinpoint the most recent pipeline update that failed. This can help in troubleshooting the issues that led to the failure, thus improving the reliability and efficiency of your pipelines.

```sql+postgres
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

```sql+sqlite
select
  p.pipeline_id,
  p.name,
  p.state,
  json_extract(u.value, '$.creation_time') as update_creation_time,
  json_extract(u.value, '$.state') as update_state,
  json_extract(u.value, '$.update_id') as update_id,
  account_id
from
  databricks_pipeline p,
  json_each(p.latest_updates) as u
where
  json_extract(u.value, '$.state') = 'FAILED'
order by
  update_creation_time desc limit 1;
```

### Get pipelines publishing data in a catalog table
Discover the segments that have active data pipelines publishing to a catalog in Databricks. This is useful for identifying and managing data flow across different pipelines in your account.

```sql+postgres
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

```sql+sqlite
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
Explore which pipelines in your Databricks environment are manually triggered, allowing you to understand which processes require user initiation and potentially optimize for automation.

```sql+postgres
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

```sql+sqlite
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
  continuous = 0;
```

### List unhealthy pipelines
Discover the segments that consist of unhealthy pipelines in your Databricks environment. This allows for quick identification and troubleshooting of problematic pipelines, enhancing system performance and reliability.

```sql+postgres
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

```sql+sqlite
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
Explore which pipelines are currently in development mode. This can be useful to identify the pipelines that are under development and not yet in production, allowing you to manage and track your development resources effectively.

```sql+postgres
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

```sql+sqlite
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
  development = 1;
```

### Get the permissions associated to each pipeline
Explore the different permissions assigned to each data pipeline to understand who has access and at what level. This can help manage security and access control within your Databricks environment.

```sql+postgres
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

```sql+sqlite
select
  pipeline_id,
  name,
  json_extract(acl.value, '$.user_name') as principal_user_name,
  json_extract(acl.value, '$.group_name') as principal_group_name,
  json_extract(acl.value, '$.all_permissions') as permission_level
from
  databricks_pipeline,
  json_each(databricks_pipeline.pipeline_permissions, '$.access_control_list') as acl;
```

### List libraries installed on each pipeline
Explore the various libraries installed on each pipeline in your Databricks environment. This can help you manage dependencies and understand the resources each pipeline is utilizing.

```sql+postgres
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

```sql+sqlite
select
  pipeline_id,
  name,
  json_extract(l.value, '$.notebook.path') as notebook_path,
  json_extract(l.value, '$.maven') as maven,
  json_extract(l.value, '$.whl') as whl,
  json_extract(l.value, '$.jar') as jar,
  json_extract(l.value, '$.file.path') as file_path,
  account_id
from
  databricks_pipeline,
  json_each(libraries) as l;
```

### Get trigger settings for each pipeline
Analyze the settings to understand the trigger configurations for each data processing pipeline. This can help in assessing the frequency and type of triggers, which is crucial for managing and optimizing data workflows.

```sql+postgres
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

```sql+sqlite
select
  pipeline_id,
  name,
  json_extract(trigger, '$.cron') as cron,
  json_extract(trigger, '$.manual') as is_manual,
  account_id
from
  databricks_pipeline
where
  trigger is not null;
```

### Get cluster settings for each pipeline
Explore the configuration of each pipeline to understand its associated settings, such as the type of node it uses, whether it has autoscale enabled, and the number of workers it employs. This information can be useful in optimizing pipeline performance and resource usage.

```sql+postgres
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

```sql+sqlite
select
  pipeline_id,
  name,
  json_extract(c.value, '$.instance_pool_id') as instance_pool_id,
  json_extract(c.value, '$.node_type_id') as node_type_id,
  json_extract(c.value, '$.autoscale') as autoscale,
  json_extract(c.value, '$.num_workers') as num_workers,
  json_extract(c.value, '$.policy_id') as policy_id,
  account_id
from
  databricks_pipeline,
  json_each(clusters) as c;
```