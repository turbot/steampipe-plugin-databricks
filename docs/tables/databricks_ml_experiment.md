# Table: databricks_ml_experiment

Experiments are the primary unit of organization in MLflow; all MLflow runs belong to an experiment. Each experiment lets you visualize, search, and compare runs, as well as download run artifacts or metadata for analysis in other tools.

## Examples

### Basic info

```sql
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

```sql
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

### List experiments that have not been modified in the last 90 days

```sql
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

### List all active experiments

```sql
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

```sql
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