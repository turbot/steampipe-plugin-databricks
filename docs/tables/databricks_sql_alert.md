---
title: "Steampipe Table: databricks_sql_alert - Query Databricks SQL Alerts using SQL"
description: "Allows users to query Databricks SQL Alerts, providing insights into the alerts and notifications created within the Databricks SQL platform."
---

# Table: databricks_sql_alert - Query Databricks SQL Alerts using SQL

Databricks SQL Alerts is a feature of the Databricks SQL platform that enables users to create alerts based on SQL queries. These alerts can be set to notify users when certain conditions are met within the data. It provides a way to monitor and respond to changes in data, ensuring the health and performance of Databricks SQL applications.

## Table Usage Guide

The `databricks_sql_alert` table provides insights into the alerts set up within Databricks SQL. As a data engineer or data analyst, you can explore alert-specific details through this table, including alert conditions, schedules, and associated metadata. Utilize it to monitor and manage your Databricks SQL alerts, ensuring the data integrity and performance of your Databricks SQL applications.

## Examples

### Basic info
Uncover the details of alerts in Databricks SQL by identifying their unique identifiers, names, creation dates and associated account IDs. This can be beneficial for tracking alert history and managing account-related alerts.

```sql+postgres
select
  id,
  name,
  created_at,
  query ->> 'query' as query,
  account_id
from
  databricks_sql_alert;
```

```sql+sqlite
select
  id,
  name,
  created_at,
  json_extract(query, '$.query') as query,
  account_id
from
  databricks_sql_alert;
```

### List alerts triggered in the past 24 hours
Gain insights into recent system alerts by identifying those that have been triggered within the past day. This can be useful for monitoring system health and responding promptly to potential issues.

```sql+postgres
select
  id,
  name,
  parent,
  created_at,
  last_triggered_at,
  query ->> 'query' as query,
  account_id
from
  databricks_sql_alert
where
  last_triggered_at > now() - interval '24' hour;
```

```sql+sqlite
select
  id,
  name,
  parent,
  created_at,
  last_triggered_at,
  json_extract(query, '$.query') as query,
  account_id
from
  databricks_sql_alert
where
  last_triggered_at > datetime('now', '-24 hours');
```

### List alerts that trigger every minute
Explore which alerts on Databricks are set to trigger every minute. This is useful in understanding the frequency of these alerts and potentially identifying areas where alert frequency could be optimized.

```sql+postgres
select
  id,
  name,
  parent,
  created_at,
  last_triggered_at,
  query ->> 'query' as query,
  account_id
from
  databricks_sql_alert
where
  rearm = 60;
```

```sql+sqlite
select
  id,
  name,
  parent,
  created_at,
  last_triggered_at,
  json_extract(query, '$.query') as query,
  account_id
from
  databricks_sql_alert
where
  rearm = 60;
```

### List dashboards modified in the past 7 days
Explore which dashboards have been updated recently to keep track of changes and modifications. This is particularly useful for maintaining up-to-date information and ensuring the accuracy of data displayed on your dashboards.

```sql+postgres
select
  id,
  name,
  created_at,
  query ->> 'query' as query,
  account_id
from
  databricks_sql_alert
where
  updated_at > now() - interval '7' day;
```

```sql+sqlite
select
  id,
  name,
  created_at,
  json_extract(query, '$.query') as query,
  account_id
from
  databricks_sql_alert
where
  updated_at > datetime('now', '-7 day');
```

### List alerts that did not fulfill trigger conditions
Explore alerts that were created but did not meet the necessary trigger conditions. This is useful for identifying potential issues or inefficiencies in your alert configuration.  

```sql+postgres
select
  id,
  name,
  parent,
  created_at,
  last_triggered_at,
  query ->> 'query' as query,
  account_id
from
  databricks_sql_alert
where
  state = 'ok';
```

```sql+sqlite
select
  id,
  name,
  parent,
  created_at,
  last_triggered_at,
  json_extract(query, '$.query') as query,
  account_id
from
  databricks_sql_alert
where
  state = 'ok';
```

### Get configuration options of an alert
Analyze the settings to understand the configuration options of an alert in a Databricks SQL environment. This could be beneficial in optimizing alert management and understanding the specifics of alert triggers.

```sql+postgres
select
  id,
  name,
  options ->> 'column' as column,
  options ->> 'op' as operator,
  options ->> 'value' as value,
  options ->> 'custom_body' as custom_body,
  options ->> 'custom_subject' as custom_subject,
  options ->> 'muted' as muted,
  account_id
from
  databricks_sql_alert;
```

```sql+sqlite
select
  id,
  name,
  json_extract(options, '$.column') as column,
  json_extract(options, '$.op') as operator,
  json_extract(options, '$.value') as value,
  json_extract(options, '$.custom_body') as custom_body,
  json_extract(options, '$.custom_subject') as custom_subject,
  json_extract(options, '$.muted') as muted,
  account_id
from
  databricks_sql_alert;
```