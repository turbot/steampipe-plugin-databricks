---
title: "Steampipe Table: databricks_workspace_git_credential - Query Databricks Workspace Git Credentials using SQL"
description: "Allows users to query Databricks Workspace Git Credentials, providing insights into the Git credentials configured in Databricks workspaces."
---

# Table: databricks_workspace_git_credential - Query Databricks Workspace Git Credentials using SQL

Databricks Workspace is a collaborative environment for data engineers, data scientists, machine learning engineers, and business users. It allows these users to create, share, and collaborate on Databricks notebooks, jobs, and data. Git Credentials in Databricks Workspace are used to connect to Git repositories and allow users to version control notebooks and other work in their workspace.

## Table Usage Guide

The `databricks_workspace_git_credential` table provides insights into the Git credentials configured in Databricks workspaces. As a data engineer or data scientist, explore the details of these credentials through this table, including the URL of the Git repository, the username used for authentication, and the last time the credentials were updated. Utilize it to manage and monitor the use of Git repositories in your Databricks workspaces.

## Examples

### Basic info
Explore which credentials are associated with specific Git providers and usernames in your Databricks workspace. This can be useful in managing and auditing access to your code repositories.

```sql+postgres
select
  credential_id,
  git_provider,
  git_username,
  account_id
from
  databricks_workspace_git_credential;
```

```sql+sqlite
select
  credential_id,
  git_provider,
  git_username,
  account_id
from
  databricks_workspace_git_credential;
```

### Get git credential info for gitHub
Gain insights into the GitHub credentials used within your Databricks workspace. This can be useful for understanding the distribution of different user accounts linked to your workspace, aiding in user management and security.

```sql+postgres
select
  credential_id,
  git_provider,
  git_username,
  account_id
from
  databricks_workspace_git_credential
where
  git_provider = 'gitHub';
```

```sql+sqlite
select
  credential_id,
  git_provider,
  git_username,
  account_id
from
  databricks_workspace_git_credential
where
  git_provider = 'gitHub';
```

### List the account in order of git credentials
Analyze the settings to understand the distribution of Git credentials across various accounts. This is useful for identifying accounts that might be overusing or underutilizing Git credentials, aiding in resource allocation and potential security auditing.

```sql+postgres
select
  account_id,
  count(*) as git_cred_count
from
  databricks_workspace_git_credential
group by
  account_id
order by
  git_cred_count desc;
```

```sql+sqlite
select
  account_id,
  count(*) as git_cred_count
from
  databricks_workspace_git_credential
group by
  account_id
order by
  git_cred_count desc;
```