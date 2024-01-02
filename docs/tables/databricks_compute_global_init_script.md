---
title: "Steampipe Table: databricks_compute_global_init_script - Query Databricks Global Init Scripts using SQL"
description: "Allows users to query Databricks Global Init Scripts, specifically the details of scripts that run during the initialization of all clusters in a Databricks workspace."
---

# Table: databricks_compute_global_init_script - Query Databricks Global Init Scripts using SQL

Databricks Global Init Scripts are scripts that run during the initialization of all clusters in a Databricks workspace. These scripts can be used to install software, download data, or change configurations. They provide a flexible way to customize the environment of all clusters in a workspace.

## Table Usage Guide

The `databricks_compute_global_init_script` table provides insights into the global initialization scripts within Databricks. As a DevOps engineer, explore script-specific details through this table, including script ID, content, and associated metadata. Utilize it to uncover information about scripts, such as their content, the clusters they are associated with, and the effects they have on the initialization of clusters.

## Examples

### Basic info
Explore which scripts were created, when, and by whom in your Databricks compute environment. This is beneficial to understand the timeline and authorship of scripts, aiding in accountability and management.

```sql+postgres
select
  script_id,
  name,
  created_at,
  created_by,
  account_id
from
  databricks_compute_global_init_script;
```

```sql+sqlite
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
Identify the scripts that were created in the past week in Databricks. This can be useful for keeping track of recent additions and changes to your scripts.

```sql+postgres
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

```sql+sqlite
select
  script_id,
  name,
  created_at,
  created_by,
  account_id
from
  databricks_compute_global_init_script
where
  created_at >= datetime('now', '-7 days');
```

### List scripts that are disabled
Determine the scripts that have been deactivated, to understand which scripts are not currently in use. This can be helpful in managing resources and maintaining an organized script library.

```sql+postgres
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

```sql+sqlite
select
  script_id,
  name,
  created_at,
  created_by,
  account_id
from
  databricks_compute_global_init_script
where
  enabled = 0;
```

### List scripts that have not been modified in last 90 days
Discover the scripts that have remained unaltered for the past 90 days. This query can be useful in identifying outdated or unused scripts in the Databricks compute environment, helping to maintain clean and efficient code repositories.

```sql+postgres
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

```sql+sqlite
select
  script_id,
  name,
  created_at,
  created_by,
  account_id
from
  databricks_compute_global_init_script
where
  updated_at <= datetime('now', '-90 day');
```

### Get script details for a given script id
Explore the specifics of a particular script, such as its name, creation date, author, and associated account. This could be beneficial for auditing purposes or to understand the context and history of a script.

```sql+postgres
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

```sql+sqlite
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
Discover which account utilizes the most global initialization scripts. This could be particularly useful in identifying high resource usage or potential optimization opportunities within your Databricks environment.

```sql+postgres
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

```sql+sqlite
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