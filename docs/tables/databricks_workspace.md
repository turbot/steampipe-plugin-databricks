---
title: "Steampipe Table: databricks_workspace - Query Databricks Workspaces using SQL"
description: "Allows users to query Databricks Workspaces, providing insights into workspace properties such as name, location, and SKU."
---

# Table: databricks_workspace - Query Databricks Workspaces using SQL

A Databricks Workspace is an environment for accessing all of your Databricks assets. The workspace organizes objects (notebooks, libraries, and experiments) into folders and provides access to data objects and computational resources including clusters, jobs, and models. Each workspace is associated with a Databricks account and includes a number of features for collaborative work.

## Table Usage Guide

The `databricks_workspace` table provides insights into Databricks Workspaces. As a data engineer or data scientist, explore workspace-specific details through this table, including the workspace name, location, and SKU. Utilize it to uncover information about workspaces, such as the workspace's managed resource group ID, managed private network, and the provisioning state of the workspace.

## Examples

### Basic info
Explore which objects were created within the Databricks workspace for a specific user. This can help in understanding the user's activities and the resources they've used.

```sql+postgres
select
  object_id,
  created_at,
  language,
  object_type,
  path,
  size,
  account_id
from
  databricks_workspace
where
  path = '/Users/user@turbot.com/NotebookDev';
```

```sql+sqlite
select
  object_id,
  created_at,
  language,
  object_type,
  path,
  size,
  account_id
from
  databricks_workspace
where
  path = '/Users/user@turbot.com/NotebookDev';
```

### List all objects in workspace created in the past 7 days
Explore the most recent additions to your workspace by identifying objects that have been created within the past week. This is particularly useful for keeping track of recent changes and additions, ensuring you stay updated on the most current workspace content.

```sql+postgres
select
  object_id,
  created_at,
  language,
  object_type,
  path,
  size,
  account_id
from
  databricks_workspace
where
  created_at >= now() - interval '7' day;
```

```sql+sqlite
select
  object_id,
  created_at,
  language,
  object_type,
  path,
  size,
  account_id
from
  databricks_workspace
where
  created_at >= datetime('now', '-7 day');
```

### List all objects in workspace modified in the past 30 days
Explore which items in your workspace have been updated in the past month. This can be useful for tracking recent changes and understanding the current state of your workspace.

```sql+postgres
select
  object_id,
  modified_at,
  language,
  object_type,
  path,
  size,
  account_id
from
  databricks_workspace
where
  modified_at >= now() - interval '30' day;
```

```sql+sqlite
select
  object_id,
  modified_at,
  language,
  object_type,
  path,
  size,
  account_id
from
  databricks_workspace
where
  modified_at >= datetime('now', '-30 day');
```

### List total objects per type in workspace
Explore the distribution of different object types within your workspace to understand the composition and organization of your data. This can assist in managing resources and identifying potential areas for optimization or reorganization.

```sql+postgres
select
  object_type,
  count(*) as total_objects
from
  databricks_workspace
group by
  object_type;
```

```sql+sqlite
select
  object_type,
  count(*) as total_objects
from
  databricks_workspace
group by
  object_type;
```

### List total notebook objects per language in workspace
Analyze the distribution of notebook objects across different programming languages in your workspace. This could be useful to understand the most commonly used languages and guide future training or tool development.

```sql+postgres
select
  language,
  count(*) as total_notebooks
from
  databricks_workspace
where
  object_type = 'NOTEBOOK'
group by
  language;
```

```sql+sqlite
select
  language,
  count(*) as total_notebooks
from
  databricks_workspace
where
  object_type = 'NOTEBOOK'
group by
  language;
```