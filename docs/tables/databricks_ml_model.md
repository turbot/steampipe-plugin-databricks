---
title: "Steampipe Table: databricks_ml_model - Query Databricks ML Models using SQL"
description: "Allows users to query Databricks ML Models, specifically providing detailed information about the machine learning models used in Databricks."
---

# Table: databricks_ml_model - Query Databricks ML Models using SQL

Databricks ML Models is a feature within Databricks that allows users to manage the full lifecycle of machine learning models. It offers capabilities to track, compare, and visualize machine learning experiments. It also enables the deployment of machine learning models in a consistent and reproducible manner.

## Table Usage Guide

The `databricks_ml_model` table provides insights into ML models within Databricks. As a data scientist or machine learning engineer, you can explore model-specific details through this table, including the model name, version, creation timestamp, and user ID of the creator. Utilize it to track and manage the lifecycle of your machine learning models, compare different models, and ensure reproducibility of your experiments.

## Examples

### Basic info
Explore the creation and update details of machine learning models on Databricks. This can help you assess the activity and usage of models over time, providing insights into your data science operations.

```sql+postgres
select
  name,
  creation_timestamp,
  description,
  last_updated_timestamp,
  user_id,
  account_id
from
  databricks_ml_model;
```

```sql+sqlite
select
  name,
  creation_timestamp,
  description,
  last_updated_timestamp,
  user_id,
  account_id
from
  databricks_ml_model;
```

### List models modified in the last 7 days
Discover the modifications made to machine learning models in the past week. This is useful for keeping track of recent changes and understanding how your models are evolving over time.

```sql+postgres
select
  name,
  creation_timestamp,
  description,
  last_updated_timestamp,
  user_id,
  account_id
from
  databricks_ml_model
where
  last_updated_timestamp > now() - interval '7' day;
```

```sql+sqlite
select
  name,
  creation_timestamp,
  description,
  last_updated_timestamp,
  user_id,
  account_id
from
  databricks_ml_model
where
  last_updated_timestamp > datetime('now', '-7 day');
```

### Get users permission level for each model
Explore which users have certain permission levels for each machine learning model in your Databricks environment. This can help ensure appropriate access rights are maintained across your team.

```sql+postgres
select
  name,
  user_id,
  permission_level,
  account_id
from
  databricks_ml_model;
```

```sql+sqlite
select
  name,
  user_id,
  permission_level,
  account_id
from
  databricks_ml_model;
```

### List all models with a specific permission level
Explore which machine learning models within your Databricks account have a specific permission level. This could be beneficial in managing access control and understanding who has the ability to manage certain models.

```sql+postgres
select
  name,
  user_id,
  permission_level,
  account_id
from
  databricks_ml_model
where
  permission_level = 'CAN_MANAGE';
```

```sql+sqlite
select
  name,
  user_id,
  permission_level,
  account_id
from
  databricks_ml_model
where
  permission_level = 'CAN_MANAGE';
```

### List details of version for each model
Determine the version details for each machine learning model in your Databricks environment, including its creation time, status, and source. This can be useful for tracking model evolution and understanding the current stage of each model version.

```sql+postgres
select
  name,
  user_id,
  permission_level,
  account_id,
  v ->> 'Version' as version,
  v ->> 'CreationTimestamp' as creation_time,
  v ->> 'Name' as version_name,
  v ->> 'RunId' as run_id,
  v ->> 'Status' as version_status,
  v ->> 'UserId' as user_id,
  v ->> 'Source' as version_source,
  v ->> 'CurrentStage' as current_version_stage
from
  databricks_ml_model,
  jsonb_array_elements(latest_versions) as v;
```

```sql+sqlite
select
  name,
  user_id,
  permission_level,
  account_id,
  json_extract(v.value, '$.Version') as version,
  json_extract(v.value, '$.CreationTimestamp') as creation_time,
  json_extract(v.value, '$.Name') as version_name,
  json_extract(v.value, '$.RunId') as run_id,
  json_extract(v.value, '$.Status') as version_status,
  json_extract(v.value, '$.UserId') as user_id,
  json_extract(v.value, '$.Source') as version_source,
  json_extract(v.value, '$.CurrentStage') as current_version_stage
from
  databricks_ml_model,
  json_each(latest_versions) as v;
```