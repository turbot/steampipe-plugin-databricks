# Table: databricks_ml_model

MLflow Model Registry is a centralized model repository and a UI and set of APIs that enable you to manage the full lifecycle of MLflow Models.

## Examples

### Basic info

```sql
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

```sql
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

### Get users permission level for each model

```sql
select
  name,
  user_id,
  permission_level,
  account_id
from
  databricks_ml_model;
```

### List all models with a specific permission level

```sql
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