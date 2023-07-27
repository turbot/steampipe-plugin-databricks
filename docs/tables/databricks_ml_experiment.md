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

### List experiments modified in the last 7 days

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
  last_update_time > now() - interval '7' day;
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