# Table: databricks_compute_global_init_script

The Global Init Scripts API enables Workspace administrators to configure global initialization scripts for their workspace. These scripts run on every node in every cluster in the workspace.

## Examples

### Basic info

```sql
select
  script_id,
  name,
  created_at,
  created_by,
  account_id
from
  databricks_compute_global_init_script;
```

### List scripts created in the last 7 days

```sql
select
  script_id,
  name,
  created_at,
  created_by,
  account_id
from
  databricks_compute_global_init_script
where
  created_at >= now() - interval '7' day;;
```

### List scripts that are disabled

```sql
select
  script_id,
  name,
  created_at,
  created_by,
  account_id
from
  databricks_compute_global_init_script
where
  not enabled;
```

### List scripts that have not been modified in last 90 days

```sql
select
  script_id,
  name,
  created_at,
  created_by,
  account_id
from
  databricks_compute_global_init_script
where
  updated_at <= now() - interval '90' day;
```

### Get script details for a given script id

```sql
select
  script_id,
  name,
  created_at,
  created_by,
  script,
  account_id
from
  databricks_compute_global_init_script
where
  script_id = 'script_id';
```

### Find the account with the most global init scripts

```sql
select
  account_id,
  count(*) as script_count
from
  databricks_compute_global_init_script
group by
  account_id
order by
  script_count desc
limit 1;
```