---
title: "Steampipe Table: databricks_settings_token - Query Databricks Settings Tokens using SQL"
description: "Allows users to query Databricks Settings Tokens, specifically providing insights into token settings and associated metadata."
---

# Table: databricks_settings_token - Query Databricks Settings Tokens using SQL

Databricks Settings Tokens are a part of the Databricks service that allows users to manage and control access to the Databricks environment. These tokens are used for authentication and can be scoped to provide specific permissions, allowing for granular control over access and actions within the environment. They are crucial for maintaining security and managing user access within the Databricks platform.

## Table Usage Guide

The `databricks_settings_token` table provides insights into Settings Tokens within Databricks. As a DevOps engineer or security analyst, explore token-specific details through this table, including permissions, lifespan, and associated metadata. Utilize it to uncover information about tokens, such as those with extensive permissions or nearing expiry, aiding in maintaining security and managing user access within the Databricks platform.

## Examples

### Basic info
Explore the creation and expiration details of tokens in your Databricks settings to identify who created them and when, providing a comprehensive view of token activity for account management and security purposes.

```sql+postgres
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

```sql+sqlite
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

### List tokens created in the last 30 days
Discover the segments that have been recently created within the last 30 days. This can provide insights into the users' activity and help track any unusual or suspicious behavior.

```sql+postgres
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
  databricks_settings_token
where
  creation_time >= datetime('now', '-30 day');
```

### List all tokens expiring in the next 7 days
Identify tokens that are set to expire within the upcoming week. This is useful for proactively managing access permissions and avoiding unexpected disruptions.

```sql+postgres
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

```sql+sqlite
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
  expiry_time > datetime('now') and expiry_time < datetime('now', '+7 day');
```

### Get number of days each token is valid for
Analyze the validity duration of each token in your Databricks settings to prioritize renewal or removal actions. This helps optimize resource usage and enhance security by preventing the misuse of expired tokens.

```sql+postgres
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

```sql+sqlite
select
  token_id,
  comment,
  julianday(expiry_time) - julianday('now') as days_remaining,
  account_id
from
  databricks_settings_token
order by
  days_remaining desc;
```

### List the owner in order of the number of tokens
Analyze the settings to understand the allocation of tokens among owners. This can help in identifying the users who have been assigned the most tokens, enabling better management and distribution of resources.

```sql+postgres
select
  owner_id,
  created_by_username,
  count(*) as token_count
from
  databricks_settings_token
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
  databricks_settings_token
group by
  owner_id,
  created_by_username
order by
  token_count desc;
```