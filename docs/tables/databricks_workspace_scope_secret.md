# Table: databricks_workspace_scope_secret

Sometimes accessing data requires that you authenticate to external data sources through JDBC. Instead of directly entering your credentials into a notebook, use Databricks secrets to store your credentials and reference them in notebooks and jobs.

## Examples

### Basic info

```sql
select
  scope_name,
  key,
  last_updated_timestamp,
  account_id
from
  databricks_workspace_scope_secret;
```

### List all secrets updated in the past 7 days

```sql
select
  scope_name,
  key,
  last_updated_timestamp,
  account_id
from
  databricks_workspace_scope_secret
where
  last_updated_timestamp > now() - interval '7' day;
```

### List total secrets per scope

```sql
select
  scope_name,
  count(*) as total_secrets
from
  databricks_workspace_scope_secret
group by
  scope_name;
```

### Get all secrets for a specific scope

```sql
select
  scope_name,
  key,
  last_updated_timestamp,
  account_id
from
  databricks_workspace_scope_secret
where
  scope_name = 'my_scope';
```