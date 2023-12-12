---
title: "Steampipe Table: databricks_ml_experiment - Query Databricks ML Experiments using SQL"
description: "Allows users to query Databricks ML Experiments, specifically retrieving information about each experiment such as its ID, name, artifact location, and lifecycle stage."
---

# Table: databricks_ml_experiment - Query Databricks ML Experiments using SQL

Databricks ML Experiments is a feature within Databricks that allows you to track, manage, and visualize machine learning experiments. It provides a centralized way to log parameters, metrics, and artifacts for each run in an experiment. Databricks ML Experiments helps you stay informed about the performance of your machine learning models and take appropriate actions based on the insights derived from the data.

## Table Usage Guide

The `databricks_ml_experiment` table provides insights into ML experiments within Databricks. As a data scientist, explore experiment-specific details through this table, including experiment ID, name, artifact location, and lifecycle stage. Utilize it to uncover information about experiments, such as their current lifecycle stage, the location of their artifacts, and other associated metadata.

## Examples

### Basic info
Gain insights into the creation and last update times of machine learning experiments within a specific account on Databricks. This is useful for tracking the progress and activity of various experiments over time.

```sql+postgres
select
  experiment_id,
  name,
  creation_time,
  last_update_time,
  account_id
from
  databricks_ml_experiment;
```

```sql+sqlite
select
  experiment_id,
  name,
  creation_time,
  last_update_time,
  account_id
from
  databricks_ml_experiment;
```

### List experiments created in the last 7 days
Discover the newly created experiments within the past week. This allows for a timely review and monitoring of the latest experiments, ensuring up-to-date insights and understanding.

```sql+postgres
select
  experiment_id,
  name,
  creation_time,
  last_update_time,
  account_id
from
  databricks_ml_experiment
where
  creation_time >= now() - interval '7' day;
```

```sql+sqlite
select
  experiment_id,
  name,
  creation_time,
  last_update_time,
  account_id
from
  databricks_ml_experiment
where
  creation_time >= datetime('now', '-7 days');
```

### List experiments that have not been modified in the last 90 days
Discover the segments that have not been updated or modified in the last 90 days in your Databricks machine learning experiments. This can help in identifying dormant or inactive experiments for potential clean up or review.

```sql+postgres
select
  experiment_id,
  name,
  creation_time,
  last_update_time,
  account_id
from
  databricks_ml_experiment
where
  last_update_time <= now() - interval '90' day;
```

```sql+sqlite
select
  experiment_id,
  name,
  creation_time,
  last_update_time,
  account_id
from
  databricks_ml_experiment
where
  last_update_time <= datetime('now', '-90 day');
```

### List all active experiments
Explore which experiments are currently active in your Databricks machine learning workflow. This allows you to keep track of ongoing studies and manage resources effectively.

```sql+postgres
select
  experiment_id,
  name,
  creation_time,
  last_update_time,
  account_id
from
  databricks_ml_experiment
where
  lifecycle_stage = 'active';
```

```sql+sqlite
select
  experiment_id,
  name,
  creation_time,
  last_update_time,
  account_id
from
  databricks_ml_experiment
where
  lifecycle_stage = 'active';
```

### Find the account with the most experiments
Discover which account has conducted the most experiments, helping you identify the most active accounts and understand usage patterns. This can be beneficial in assessing resource allocation and planning future capacity needs.

```sql+postgres
select
  account_id,
  count(*) as experiment_count
from
  databricks_ml_experiment
group by
  account_id
order by
  experiment_count desc
limit 1;
```

```sql+sqlite
select
  account_id,
  count(*) as experiment_count
from
  databricks_ml_experiment
group by
  account_id
order by
  experiment_count desc
limit 1;
```