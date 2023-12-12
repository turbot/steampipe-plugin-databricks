---
title: "Steampipe Table: databricks_pipeline_event - Query Databricks Pipeline Events using SQL"
description: "Allows users to query Databricks Pipeline Events, providing insights into the execution history of Databricks pipelines."
---

# Table: databricks_pipeline_event - Query Databricks Pipeline Events using SQL

Databricks Pipelines is a feature of Databricks that allows users to build, test, and manage end-to-end, continuous data pipelines. These pipelines can ingest data from multiple sources, transform and aggregate it, and store the result in any data sink. Databricks Pipelines also supports automatic tracking of data lineage and audit logs, making it easier to maintain and troubleshoot pipelines.

## Table Usage Guide

The `databricks_pipeline_event` table provides insights into the execution history of Databricks pipelines. As a data engineer, you can explore pipeline-specific details through this table, including run times, status, errors, and associated metadata. Utilize it to monitor pipeline performance, troubleshoot issues, and understand usage patterns.

## Examples

### Basic info
Analyze the settings to understand the characteristics and details of various pipeline events in your Databricks account. This can help in assessing the maturity level and type of events occurring, providing valuable insights for optimizing pipeline performance.

```sql+postgres
select
  id,
  pipeline_id,
  event_type,
  level,
  maturity_level,
  message,
  account_id
from
  databricks_pipeline_event;
```

```sql+sqlite
select
  id,
  pipeline_id,
  event_type,
  level,
  maturity_level,
  message,
  account_id
from
  databricks_pipeline_event;
```

### List events between a specific time range
Explore which pipeline events occurred within a specific time frame to analyze their impact and assess any potential issues. This can be particularly useful in understanding the performance and reliability of your pipelines during peak usage times.

```sql+postgres
select
  id,
  pipeline_id,
  event_type,
  level,
  maturity_level,
  message,
  account_id
from
  databricks_pipeline_event
where
  timestamp between '2023-07-27T02:00:00' and '2023-07-27T22:00:00';
```

```sql+sqlite
select
  id,
  pipeline_id,
  event_type,
  level,
  maturity_level,
  message,
  account_id
from
  databricks_pipeline_event
where
  timestamp between '2023-07-27T02:00:00' and '2023-07-27T22:00:00';
```

### List all events that are errors
Explore which pipeline events in your Databricks account have been flagged as errors. This analysis can help you identify potential issues and troubleshoot them effectively.

```sql+postgres
select
  id,
  pipeline_id,
  event_type,
  level,
  message,
  error ->> 'exceptions' as exceptions,
  error ->> 'fatal' as fatal,
  account_id
from
  databricks_pipeline_event
where
  level = 'ERROR';
```

```sql+sqlite
select
  id,
  pipeline_id,
  event_type,
  level,
  message,
  json_extract(error, '$.exceptions') as exceptions,
  json_extract(error, '$.fatal') as fatal,
  account_id
from
  databricks_pipeline_event
where
  level = 'ERROR';
```

### Get origin for all events
Explore the origins of all events in order to gain insights into their source, such as the cloud, region, and pipeline they originated from. This can be useful for understanding event trends, tracking specific events, and monitoring the overall health of your pipeline.

```sql+postgres
select
  id,
  pipeline_id,
  event_type,
  message,
  origin ->> 'cloud' as origin_cloud,
  origin ->> 'region' as origin_region,
  origin ->> 'pipeline_id' as origin_pipeline_id,
  origin ->> 'pipeline_name' as origin_pipeline_name,
  origin ->> 'request_id' as origin_request_id,
  origin ->> 'update_id' as origin_update_id,
  account_id
from
  databricks_pipeline_event;
```

```sql+sqlite
select
  id,
  pipeline_id,
  event_type,
  message,
  json_extract(origin, '$.cloud') as origin_cloud,
  json_extract(origin, '$.region') as origin_region,
  json_extract(origin, '$.pipeline_id') as origin_pipeline_id,
  json_extract(origin, '$.pipeline_name') as origin_pipeline_name,
  json_extract(origin, '$.request_id') as origin_request_id,
  json_extract(origin, '$.update_id') as origin_update_id,
  account_id
from
  databricks_pipeline_event;
```

### List all events caused due to user actions
Explore which events have been triggered by user actions in your Databricks pipelines. This is useful in understanding user behavior and identifying potential areas for system improvement or troubleshooting.

```sql+postgres
select
  id,
  pipeline_id,
  event_type,
  level,
  maturity_level,
  message,
  account_id
from
  databricks_pipeline_event
where
  event_type = 'user_action';
```

```sql+sqlite
select
  id,
  pipeline_id,
  event_type,
  level,
  maturity_level,
  message,
  account_id
from
  databricks_pipeline_event
where
  event_type = 'user_action';
```

### List all events having stable maturity level
Explore which events in your Databricks pipeline have reached a stable maturity level. This can be useful in assessing the progress and stability of your data processing workflows.

```sql+postgres
select
  id,
  pipeline_id,
  event_type,
  level,
  maturity_level,
  message,
  account_id
from
  databricks_pipeline_event
where
  maturity_level = 'STABLE';
```

```sql+sqlite
select
  id,
  pipeline_id,
  event_type,
  level,
  maturity_level,
  message,
  account_id
from
  databricks_pipeline_event
where
  maturity_level = 'STABLE';
```

### Find the account with the most pipeline events
Identify the account with the highest number of pipeline events to understand the most active user in the Databricks environment. This can be useful in resource allocation and performance monitoring.

```sql+postgres
select
  account_id,
  count(*) as event_count
from
  databricks_pipeline_event
group by
  account_id
order by
  event_count desc
limit 1;
```

```sql+sqlite
select
  account_id,
  count(*) as event_count
from
  databricks_pipeline_event
group by
  account_id
order by
  event_count desc
limit 1;
```