# Table: databricks_ml_webhook

MLflow Model Registry is a centralized model repository and a UI and set of APIs that enable you to manage the full lifecycle of MLflow Models. Webhooks are a way to get notified when an event happens in the Model Registry. You can use webhooks to integrate Model Registry with other systems.

## Examples

### Basic info

```sql
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

### List models modified in the last 7 days

```sql
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
  last_updated_timestamp > now() - interval '7' day;
```

### List events that can trigger a webhook

```sql
select
  id,
  model_name,
  e as event,
  account_id
from
  databricks_ml_webhook,
  jsonb_array_elements_text(events) as e;
```

### List all webhooks that are disabled

```sql
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

```sql
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

### Get URL spec for each webhook

```sql
select
  id,
  model_name,
  http_url_spec ->> 'enable_ssl_verification' as enable_ssl_verification,
  http_url_spec ->> 'url' as url,
  account_id
from
  databricks_ml_webhook;
```

### Get job spec for each webhook

```sql
select
  id,
  model_name,
  job_spec ->> 'job_id' as job_id,
  job_spec ->> 'workspace_url' as workspace_url,
  account_id
from
  databricks_ml_webhook;
```