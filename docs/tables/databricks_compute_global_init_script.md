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

### List scripts modified in last 1 week

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
  updated_at >= now() - interval '7' day;
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