---
title: "Steampipe Table: databricks_pipeline_update - Query Databricks Pipelines using SQL"
description: "Allows users to query Databricks Pipelines, specifically the update details, providing insights into the pipeline status and changes."
---

# Table: databricks_pipeline_update - Query Databricks Pipelines using SQL

Databricks Pipelines is a service within Databricks that allows you to build, test, deploy, and manage machine learning workflows. It provides a centralized way to set up and manage pipelines for various Databricks resources, including models, data, and more. Databricks Pipelines helps you stay informed about the health and performance of your machine learning workflows and take appropriate actions when predefined conditions are met.

## Table Usage Guide

The `databricks_pipeline_update` table provides insights into Pipelines within Databricks. As a Data Scientist or Machine Learning Engineer, explore pipeline-specific details through this table, including status, update details, and associated metadata. Utilize it to uncover information about pipelines, such as those with recent updates, the status of pipelines, and the verification of changes.

## Examples

### Basic info
Assess the elements within your Databricks pipeline updates to gain insights into the causes and timing of updates. This can be particularly useful in identifying patterns or issues related to pipeline updates in your Databricks account.

```sql+postgres
select
  update_id,
  pipeline_id,
  cause,
  cluster_id,
  creation_time,
  account_id
from
  databricks_pipeline_update;
```

```sql+sqlite
select
  update_id,
  pipeline_id,
  cause,
  cluster_id,
  creation_time,
  account_id
from
  databricks_pipeline_update;
```

### List updates created in the last 7 days
Explore recent changes by identifying updates made within the last week. This can be beneficial in tracking modifications, understanding their causes, and assessing their impact on different accounts and clusters.

```sql+postgres
select
  update_id,
  pipeline_id,
  cause,
  cluster_id,
  creation_time,
  account_id
from
  databricks_pipeline_update
where
  creation_time >= now() - interval '7' day;
```

```sql+sqlite
select
  update_id,
  pipeline_id,
  cause,
  cluster_id,
  creation_time,
  account_id
from
  databricks_pipeline_update
where
  creation_time >= datetime('now', '-7 day');
```

### List updates caused by an API call
Discover the updates triggered by an API call. This is particularly useful for tracking changes and auditing purposes, as it allows you to see which updates were not user-initiated but were instead triggered by an API call.

```sql+postgres
select
  update_id,
  pipeline_id,
  cause,
  cluster_id,
  creation_time,
  account_id
from
  databricks_pipeline_update
where
  cause = 'API_CALL';
```

```sql+sqlite
select
  update_id,
  pipeline_id,
  cause,
  cluster_id,
  creation_time,
  account_id
from
  databricks_pipeline_update
where
  cause = 'API_CALL';
```

### List all failed updates
Explore which updates failed in your Databricks pipeline. This is useful for identifying problematic updates and understanding the cause of their failure.

```sql+postgres
select
  update_id,
  pipeline_id,
  cause,
  cluster_id,
  creation_time,
  account_id
from
  databricks_pipeline_update
where
  state = 'FAILED';
```

```sql+sqlite
select
  update_id,
  pipeline_id,
  cause,
  cluster_id,
  creation_time,
  account_id
from
  databricks_pipeline_update
where
  state = 'FAILED';
```

### List all pipelines that require full refresh before each run
Explore which pipelines necessitate a full refresh prior to each run, aiding in resource allocation and ensuring efficient pipeline management. This can be particularly useful in scenarios where pipeline performance is crucial and resources are limited.

```sql+postgres
select
  update_id,
  pipeline_id,
  cause,
  cluster_id,
  creation_time,
  full_refresh_selection,
  account_id
from
  databricks_pipeline_update
where
  full_refresh;
```

```sql+sqlite
select
  update_id,
  pipeline_id,
  cause,
  cluster_id,
  creation_time,
  full_refresh_selection,
  account_id
from
  databricks_pipeline_update
where
  full_refresh = 1;
```

### Find the account with the most pipeline updates
Identify the account with the highest frequency of pipeline updates. This can be useful in understanding which account is most actively managing and modifying their pipelines.

```sql+postgres
select
  account_id,
  count(*) as update_count
from
  databricks_pipeline_update
group by
  account_id
order by
  update_count desc
limit 1;
```

```sql+sqlite
select
  account_id,
  count(*) as update_count
from
  databricks_pipeline_update
group by
  account_id
order by
  update_count desc
limit 1;
```