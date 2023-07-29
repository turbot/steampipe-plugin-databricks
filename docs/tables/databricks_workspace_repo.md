# Table: databricks_workspace_repo

Databricks Repos is a visual Git client in Databricks. It supports common Git operations such a cloning a repository, committing and pushing, pulling, branch management, and visual comparison of diffs when committing.

## Examples

### Basic info

```sql
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

### List patterns included for sparse checkout

```sql
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

### List total repos per provider

```sql
select
  provider,
  count(*) as total_repos
from
  databricks_workspace_repo
group by
  provider;
```