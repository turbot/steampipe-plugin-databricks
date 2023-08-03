# Table: databricks_ml_webhook

Webhooks enable you to listen for Model Registry events so your integrations can automatically trigger actions. You can use webhooks to automate and integrate your machine learning pipeline with existing CI/CD tools and workflows. For example, you can trigger CI builds when a new model version is created or notify your team members through Slack each time a model transition to production is requested.

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