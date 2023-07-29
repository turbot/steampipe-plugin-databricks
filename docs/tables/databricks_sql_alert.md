# Table: databricks_sql_alert

An alert is a Databricks SQL object that periodically runs a query, evaluates a condition of its result, and notifies one or more users and/or notification destinations if the condition was met.

## Examples

### Basic info

```sql
select
  id,
  name,
  created_at,
  query ->> 'query' as query,
  account_id
from
  databricks_sql_alert;
```

### List alerts triggered in the past 24 hours

```sql
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

### List alerts that trigger every minute

```sql
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

### List dashboards modified in the past 7 days

```sql
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

### List alerts that did not fulfill trigger conditions
  
```sql
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

### Get configuration options of an alert

```sql
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