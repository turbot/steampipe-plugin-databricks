# Table: databricks_workspace_git_credential

Git credentials enable us to connect a remote repo to Databricks Repos. Registers personal access token for Databricks to do operations on behalf of the user.

## Examples

### Basic info

```sql
select
  credential_id,
  git_provider,
  git_username,
  account_id
from
  databricks_workspace_git_credential;
```

### Get git credential info for gitHub

```sql
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

```sql
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