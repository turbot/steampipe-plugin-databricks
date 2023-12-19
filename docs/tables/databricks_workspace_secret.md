---
title: "Steampipe Table: databricks_workspace_secret - Query Databricks Workspace Secrets using SQL"
description: "Allows users to query Databricks Workspace Secrets, specifically providing insights into the secret metadata and values, aiding in the management and security of sensitive data."
---

# Table: databricks_workspace_secret - Query Databricks Workspace Secrets using SQL

A Databricks Workspace Secret is a resource in Databricks that allows you to store and manage sensitive data such as passwords, OAuth tokens, and SSH keys. Secrets are stored in a workspace and are available to all notebooks in that workspace. The use of secrets eliminates the need to hard code sensitive data, enhancing security and simplifying the management of sensitive data.

## Table Usage Guide

The `databricks_workspace_secret` table provides insights into the secrets stored within a Databricks Workspace. As a security analyst, explore secret-specific details through this table, including secret metadata and values. Utilize it to uncover information about sensitive data, such as secret values and metadata, aiding in the management and security of sensitive data.

## Examples

### Basic info
Explore the latest updates to secret keys within your Databricks workspace for account management purposes. This can help maintain the security and integrity of your workspace by identifying any recent changes.

```sql+postgres
select
  scope_name,
  key,
  last_updated_timestamp,
  account_id
from
  databricks_workspace_secret;
```

```sql+sqlite
select
  scope_name,
  key,
  last_updated_timestamp,
  account_id
from
  databricks_workspace_secret;
```

### List all secrets updated in the past 7 days
Explore which confidential data elements were modified in the last week. This query can be used to maintain security and ensure that changes to sensitive information are monitored regularly.

```sql+postgres
select
  scope_name,
  key,
  last_updated_timestamp,
  account_id
from
  databricks_workspace_secret
where
  last_updated_timestamp > now() - interval '7' day;
```

```sql+sqlite
select
  scope_name,
  key,
  last_updated_timestamp,
  account_id
from
  databricks_workspace_secret
where
  last_updated_timestamp > datetime('now', '-7 day');
```

### List total secrets per scope
Assess the elements within your Databricks workspace by identifying the total number of secrets each scope holds. This allows for better management and understanding of your workspace's security configuration.

```sql+postgres
select
  scope_name,
  count(*) as total_secrets
from
  databricks_workspace_secret
group by
  scope_name;
```

```sql+sqlite
select
  scope_name,
  count(*) as total_secrets
from
  databricks_workspace_secret
group by
  scope_name;
```

### Get all secrets for a specific scope
Explore the secrets within a specific scope in your Databricks workspace. This is useful for auditing or reviewing the last updated timestamp and associated account details within a given scope.

```sql+postgres
select
  scope_name,
  key,
  last_updated_timestamp,
  account_id
from
  databricks_workspace_secret
where
  scope_name = 'my_scope';
```

```sql+sqlite
select
  scope_name,
  key,
  last_updated_timestamp,
  account_id
from
  databricks_workspace_secret
where
  scope_name = 'my_scope';
```