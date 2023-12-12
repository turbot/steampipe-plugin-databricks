---
title: "Steampipe Table: databricks_settings_token_management - Query Databricks Token Management Settings using SQL"
description: "Allows users to query Databricks Token Management Settings, specifically the details about token settings in the Databricks platform."
---

# Table: databricks_settings_token_management - Query Databricks Token Management Settings using SQL

Databricks Token Management is a feature within Databricks that provides control over the generation and usage of tokens. It allows users to manage the lifespan and permissions of tokens, enhancing the security of the platform. Token Management helps in maintaining the integrity of the data and operations performed in Databricks.

## Table Usage Guide

The `databricks_settings_token_management` table provides insights into token settings within Databricks Token Management. As a security engineer, explore token-specific details through this table, including lifespan, permissions, and associated metadata. Utilize it to uncover information about tokens, such as their validity period, the operations they can perform, and the verification of their permissions.

## Examples

### Basic info
Explore which user-created tokens are currently active in your Databricks settings. This is useful to understand token management, including who created each token and when they will expire.

```sql+postgres
select
  token_id,
  comment,
  created_by_username,
  creation_time,
  expiry_time,
  account_id
from
  databricks_settings_token_management;
```

```sql+sqlite
select
  token_id,
  comment,
  created_by_username,
  creation_time,
  expiry_time,
  account_id
from
  databricks_settings_token_management;
```

### List tokens created in the last 30 days
Identify the tokens that have been created in the past month. This can be useful to monitor recent activity and manage security in your Databricks environment.

```sql+postgres
select
  token_id,
  comment,
  created_by_username,
  creation_time,
  expiry_time,
  account_id
from
  databricks_settings_token_management
where
  creation_time >= now() - interval '30' day;
```

```sql+sqlite
select
  token_id,
  comment,
  created_by_username,
  creation_time,
  expiry_time,
  account_id
from
  databricks_settings_token_management
where
  creation_time >= datetime('now','-30 day');
```

### List all tokens expiring in the next 7 days
Determine the areas in which tokens are set to expire within the upcoming week. This is useful for preemptively managing access and maintaining security within your Databricks environment.

```sql+postgres
select
  token_id,
  comment,
  created_by_username,
  creation_time,
  expiry_time,
  account_id
from
  databricks_settings_token_management
where
  expiry_time > now() and expiry_time < now() + interval '7' day;
```

```sql+sqlite
select
  token_id,
  comment,
  created_by_username,
  creation_time,
  expiry_time,
  account_id
from
  databricks_settings_token_management
where
  expiry_time > datetime('now') and expiry_time < datetime('now', '+7 day');
```

### Get number of days each token is valid for
Determine the validity duration of each token by calculating the number of days left before expiration. This can help in managing and planning the token usage effectively.

```sql+postgres
select
  token_id,
  comment,
  expiry_time - now() as days_remaining,
  account_id
from
  databricks_settings_token_management
order by
  days_remaining desc;
```

```sql+sqlite
select
  token_id,
  comment,
  julianday(expiry_time) - julianday('now') as days_remaining,
  account_id
from
  databricks_settings_token_management
order by
  days_remaining desc;
```

### List the owner in order of the number of tokens
Explore which user has created the most tokens in your Databricks configuration to better understand usage patterns and potentially optimize resource allocation.

```sql+postgres
select
  owner_id,
  created_by_username,
  count(*) as token_count
from
  databricks_settings_token_management
group by
  owner_id,
  created_by_username
order by
  token_count desc;
```

```sql+sqlite
select
  owner_id,
  created_by_username,
  count(*) as token_count
from
  databricks_settings_token_management
group by
  owner_id,
  created_by_username
order by
  token_count desc;
```