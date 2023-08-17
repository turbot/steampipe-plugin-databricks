# Table: databricks_workspace_workspace

Workspace manages the notebooks and folders in databricks. A notebook is a web-based interface to a document that contains runnable code, visualizations, and explanatory text.

## Examples

### Basic info

```sql
select
  object_id,
  created_at,
  language,
  object_type,
  path,
  size,
  account_id
from
  databricks_workspace_workspace
where
  path = '/Users/user@turbot.com/NotebookDev';
```

### List all objects in workspace created in the past 7 days

```sql
select
  object_id,
  created_at,
  language,
  object_type,
  path,
  size,
  account_id
from
  databricks_workspace_workspace
where
  created_at >= now() - interval '7' day;
```

### List all objects in workspace modified in the past 30 days

```sql
select
  object_id,
  modified_at,
  language,
  object_type,
  path,
  size,
  account_id
from
  databricks_workspace_workspace
where
  modified_at >= now() - interval '30' day;
```

### List total objects per type in workspace

```sql
select
  object_type,
  count(*) as total_objects
from
  databricks_workspace_workspace
group by
  object_type;
```

### List total notebook objects per language in workspace

```sql
select
  language,
  count(*) as total_notebooks
from
  databricks_workspace_workspace
where
  object_type = 'NOTEBOOK'
group by
  language;
```