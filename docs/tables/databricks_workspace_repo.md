---
title: "Steampipe Table: databricks_workspace_repo - Query Databricks Workspace Repositories using SQL"
description: "Allows users to query Databricks Workspace Repositories, providing detailed information about each repository including its ID, name, git status, and more."
---

# Table: databricks_workspace_repo - Query Databricks Workspace Repositories using SQL

Databricks Workspace Repositories are a feature of Databricks that allow users to manage and version control notebooks and other workspace objects. These repositories can be linked to a remote Git repository, enabling seamless integration with existing version control workflows. This functionality provides a robust and efficient way to manage, track and version control data science and machine learning workflows.

## Table Usage Guide

The `databricks_workspace_repo` table provides insights into Databricks Workspace Repositories. As a data scientist or DevOps engineer, explore repository-specific details through this table, including repository ID, name, and git status. Utilize it to manage and track your data science and machine learning workflows, ensuring efficient version control and workflow management.

## Examples

### Basic info
Explore which Databricks workspace repositories are linked to your account. This helps you assess the elements within your workspace, such as the repository path, branch, and provider, and pinpoint the specific locations where changes have been made.

```sql+postgres
select
  id,
  path,
  branch,
  provider,
  head_commit_id,
  url,
  account_id
from
  databricks_workspace_repo;
```

```sql+sqlite
select
  id,
  path,
  branch,
  provider,
  head_commit_id,
  url,
  account_id
from
  databricks_workspace_repo;
```

### List the master repositories
Determine the areas in which master repositories are used within your Databricks workspace. This can help in understanding your workspace's code base, tracking changes, and managing versions.

```sql+postgres
select
  id,
  path,
  branch,
  provider,
  head_commit_id,
  url,
  account_id
from
  databricks_workspace_repo
where
  branch = 'master';
```

```sql+sqlite
select
  id,
  path,
  branch,
  provider,
  head_commit_id,
  url,
  account_id
from
  databricks_workspace_repo
where
  branch = 'master';
```

### List repositories for github provider
Analyze your Databricks workspace to identify all repositories linked with the GitHub provider. This can help in understanding the codebase distribution and managing repositories effectively.

```sql+postgres
select
  id,
  path,
  branch,
  provider,
  head_commit_id,
  url,
  account_id
from
  databricks_workspace_repo
where
  provider = 'gitHub';
```

```sql+sqlite
select
  id,
  path,
  branch,
  provider,
  head_commit_id,
  url,
  account_id
from
  databricks_workspace_repo
where
  provider = 'gitHub';
```

### List patterns included for sparse checkout
Explore which patterns are included for sparse checkout in a Databricks workspace repository. This can help in understanding the specific files or directories that are included in the workspace without downloading the entire repository, aiding in efficient data management.

```sql+postgres
select
  id,
  path,
  branch,
  patterns,
  account_id
from
  databricks_workspace_repo,
  jsonb_array_elements_text(sparse_checkout_patterns) as patterns;
```

```sql+sqlite
select
  id,
  path,
  branch,
  patterns.value as patterns,
  account_id
from
  databricks_workspace_repo,
  json_each(sparse_checkout_patterns) as patterns;
```

### List total repos per provider
Gain insights into the distribution of repositories across different providers. This is useful for understanding which providers are most commonly used for hosting repositories in your Databricks workspace.

```sql+postgres
select
  provider,
  count(*) as total_repos
from
  databricks_workspace_repo
group by
  provider;
```

```sql+sqlite
select
  provider,
  count(*) as total_repos
from
  databricks_workspace_repo
group by
  provider;
```