# Table: databricks_settings_token

Tokens are used to authenticate and access Databricks REST APIs. This table contains information about the tokens created in your Databricks account.

## Examples

### Basic info

```sql
select
  token_id,
  comment,
  created_by_username,
  creation_time,
  expiry_time,
  account_id
from
  databricks_settings_token;
```

### List all tokens expiring in the next 7 days

```sql
select
  token_id,
  comment,
  created_by_username,
  creation_time,
  expiry_time,
  account_id
from
  databricks_settings_token
where
  expiry_time > now() and expiry_time < now() + interval '7' day;
```

### Get number of days each token is valid for

```sql
select
  token_id,
  comment,
  expiry_time - now() as days_remaining,
  account_id
from
  databricks_settings_token
order by
  days_remaining desc;
```