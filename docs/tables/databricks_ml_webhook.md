---
title: "Steampipe Table: databricks_ml_webhook - Query Databricks Machine Learning Webhooks using SQL"
description: "Allows users to query Databricks Machine Learning Webhooks, which provide real-time notifications about events in a Databricks workspace."
---

# Table: databricks_ml_webhook - Query Databricks Machine Learning Webhooks using SQL

Databricks Machine Learning Webhooks are a feature of Databricks Machine Learning, a unified platform for data science and machine learning workflows. Webhooks provide real-time notifications about events in a Databricks workspace, such as experiment runs, model registrations, and model deployments. They are useful for integrating Databricks with other tools and services, automating workflows, and monitoring Databricks activity.

## Table Usage Guide

The `databricks_ml_webhook` table provides insights into the Webhooks within Databricks Machine Learning. As a Data Scientist or Machine Learning Engineer, explore webhook-specific details through this table, including the events they track, their payloads, and the endpoints they send notifications to. Utilize it to monitor Databricks activity, automate workflows, and integrate Databricks with other tools and services.

## Examples

### Basic info
Explore the details of your Databricks machine learning webhooks, such as their creation and last updated timestamps, to help manage and monitor their usage and status. This can be particularly useful in maintaining an up-to-date understanding of your machine learning operations.

```sql+postgres
select
  id,
  model_name,
  creation_timestamp,
  description,
  last_updated_timestamp,
  status,
  account_id
from
  databricks_ml_webhook;
```

```sql+sqlite
select
  id,
  model_name,
  creation_timestamp,
  description,
  last_updated_timestamp,
  status,
  account_id
from
  databricks_ml_webhook;
```

### List models created in the last 7 days
Determine the areas in which new models have been created in the past week. This is useful for keeping track of recent developments and updates within your databricks machine learning environment.

```sql+postgres
select
  id,
  model_name,
  creation_timestamp,
  description,
  last_updated_timestamp,
  status,
  account_id
from
  databricks_ml_webhook
where
  creation_timestamp >= now() - interval '7' day;
```

```sql+sqlite
select
  id,
  model_name,
  creation_timestamp,
  description,
  last_updated_timestamp,
  status,
  account_id
from
  databricks_ml_webhook
where
  creation_timestamp >= datetime('now', '-7 day');
```

### List models that have not been modified in the last 90 days
Explore models that have remained unaltered for the past 90 days. This can be useful in identifying dormant or potentially outdated models that may require review or updates.

```sql+postgres
select
  id,
  model_name,
  creation_timestamp,
  description,
  last_updated_timestamp,
  status,
  account_id
from
  databricks_ml_webhook
where
  last_updated_timestamp <= now() - interval '90' day;
```

```sql+sqlite
select
  id,
  model_name,
  creation_timestamp,
  description,
  last_updated_timestamp,
  status,
  account_id
from
  databricks_ml_webhook
where
  last_updated_timestamp <= datetime('now', '-90 day');
```


### List events that can trigger a webhook
Explore which events have the potential to trigger a webhook in your Databricks Machine Learning account. This can help you better understand and manage your automated processes.

```sql+postgres
select
  id,
  model_name,
  e as event,
  account_id
from
  databricks_ml_webhook,
  jsonb_array_elements_text(events) as e;
```

```sql+sqlite
select
  id,
  model_name,
  e.value as event,
  account_id
from
  databricks_ml_webhook,
  json_each(events) as e;
```

### List all webhooks that are disabled
Explore which webhooks are currently disabled in your system. This could be useful to identify potential issues or gaps in your automated processes.

```sql+postgres
select
  id,
  model_name,
  creation_timestamp,
  description,
  last_updated_timestamp,
  status,
  account_id
from
  databricks_ml_webhook
where
  status = 'DISABLED';
```

```sql+sqlite
select
  id,
  model_name,
  creation_timestamp,
  description,
  last_updated_timestamp,
  status,
  account_id
from
  databricks_ml_webhook
where
  status = 'DISABLED';
```

### List all webhooks that require SSL verification
Explore which webhooks necessitate SSL verification to enhance security measures. This can be useful in identifying potential vulnerabilities and ensuring that all webhooks are properly configured for secure data transfer.

```sql+postgres
select
  id,
  model_name,
  creation_timestamp,
  description,
  last_updated_timestamp,
  status,
  account_id
from
  databricks_ml_webhook
where
  http_url_spec ->> 'enable_ssl_verification' = 'true';
```

```sql+sqlite
select
  id,
  model_name,
  creation_timestamp,
  description,
  last_updated_timestamp,
  status,
  account_id
from
  databricks_ml_webhook
where
  json_extract(http_url_spec, '$.enable_ssl_verification') = 'true';
```

### Get URL spec for each webhook
Analyze the settings to understand the configuration of each webhook, specifically focusing on SSL verification status and URL, which is useful for ensuring secure and accurate data transmission between systems.

```sql+postgres
select
  id,
  model_name,
  http_url_spec ->> 'enable_ssl_verification' as enable_ssl_verification,
  http_url_spec ->> 'url' as url,
  account_id
from
  databricks_ml_webhook;
```

```sql+sqlite
select
  id,
  model_name,
  json_extract(http_url_spec, '$.enable_ssl_verification') as enable_ssl_verification,
  json_extract(http_url_spec, '$.url') as url,
  account_id
from
  databricks_ml_webhook;
```

### Get job spec for each webhook
Explore the specific details of each webhook job, such as the job ID and the associated workspace URL. This can be useful to understand the configuration and scope of each job within the Databricks Machine Learning environment.

```sql+postgres
select
  id,
  model_name,
  job_spec ->> 'job_id' as job_id,
  job_spec ->> 'workspace_url' as workspace_url,
  account_id
from
  databricks_ml_webhook;
```

```sql+sqlite
select
  id,
  model_name,
  json_extract(job_spec, '$.job_id') as job_id,
  json_extract(job_spec, '$.workspace_url') as workspace_url,
  account_id
from
  databricks_ml_webhook;
```

### Get details of the model associated to a particular webhook
Explore the specifics of a machine learning model linked to a particular webhook. This query is useful to understand the characteristics and timelines of the model, aiding in tracking its performance and updates.

```sql+postgres
select
  w.id as webhook_id,
  m.name as model_name,
  m.creation_timestamp model_create_time,
  m.description as model_description,
  m.last_updated_timestamp as model_update_time,
  m.account_id as model_account_id
from
  databricks_ml_webhook as w
  left join databricks_ml_model as m on w.model_name = m.name;
```

```sql+sqlite
select
  w.id as webhook_id,
  m.name as model_name,
  m.creation_timestamp model_create_time,
  m.description as model_description,
  m.last_updated_timestamp as model_update_time,
  m.account_id as model_account_id
from
  databricks_ml_webhook as w
  left join databricks_ml_model as m on w.model_name = m.name;
```