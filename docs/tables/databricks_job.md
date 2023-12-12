---
title: "Steampipe Table: databricks_job - Query Databricks Jobs using SQL"
description: "Allows users to query Databricks Jobs, specifically job configurations, providing insights into job settings and potential issues."
---

# Table: databricks_job - Query Databricks Jobs using SQL

Databricks Jobs is a feature within Databricks that allows you to schedule and run computations in a reliable and scalable manner. It provides a centralized way to manage and monitor jobs for various Databricks resources, including clusters, notebooks, and more. Databricks Jobs helps you stay informed about the status and configuration of your Databricks jobs and take appropriate actions when predefined conditions are met.

## Table Usage Guide

The `databricks_job` table provides insights into job configurations within Databricks. As a DevOps engineer, explore job-specific details through this table, including settings, schedules, and associated metadata. Utilize it to uncover information about jobs, such as those with specific settings, the schedules of jobs, and the verification of job configurations.

## Examples

### Basic info
Explore which jobs have been created in Databricks, who created them, and under whose user name they're running. This can help assess accountability and monitor usage patterns within your Databricks environment.

```sql+postgres
select
  job_id,
  name,
  created_time,
  creator_user_name,
  run_as_user_name,
  format,
  account_id
from
  databricks_job;
```

```sql+sqlite
select
  job_id,
  name,
  created_time,
  creator_user_name,
  run_as_user_name,
  format,
  account_id
from
  databricks_job;
```

### Get compute requirements for each job
Assess the elements within each job to understand the specific compute requirements. This can be useful in optimizing resource allocation and managing job performance efficiently.

```sql+postgres
select
  job_id,
  name,
  compute ->> 'compute_key' as compute_key,
  compute ->> 'spec' as compute_spec,
  account_id
from
  databricks_job;
```

```sql+sqlite
select
  job_id,
  name,
  json_extract(compute, '$.compute_key') as compute_key,
  json_extract(compute, '$.spec') as compute_spec,
  account_id
from
  databricks_job;
```

### List all continuous jobs
Explore which continuous jobs are currently in operation. This can help you manage and monitor ongoing tasks within your Databricks environment, allowing you to assess their status and identify any potential issues.

```sql+postgres
select
  job_id,
  name,
  format,
  continuous ->> 'pause_status' as pause_status,
  account_id
from
  databricks_job
where
  continuous is not null;
```

```sql+sqlite
select
  job_id,
  name,
  format,
  json_extract(continuous, '$.pause_status') as pause_status,
  account_id
from
  databricks_job
where
  continuous is not null;
```

### Get email notification configuration for each job
Explore the email notification settings for each job in order to manage and customize alerts for different job outcomes such as start, success, or failure. This is particularly useful to proactively track job progress and troubleshoot any issues promptly.

```sql+postgres
select
  job_id,
  name,
  email_notifications ->> 'on_start' as email_on_start,
  email_notifications ->> 'on_success' as email_on_success,
  email_notifications ->> 'on_failure' as email_on_failure,
  email_notifications ->> 'no_alert_for_skipped_runs' as no_alert_for_skipped_runs,
  account_id
from
  databricks_job;
```

```sql+sqlite
select
  job_id,
  name,
  json_extract(email_notifications, '$.on_start') as email_on_start,
  json_extract(email_notifications, '$.on_success') as email_on_success,
  json_extract(email_notifications, '$.on_failure') as email_on_failure,
  json_extract(email_notifications, '$.no_alert_for_skipped_runs') as no_alert_for_skipped_runs,
  account_id
from
  databricks_job;
```

### Get git settings for each job
Explore the configuration of each job in terms of its associated Git settings to understand how different jobs are linked to different versions and branches of your codebase. This can help you track the evolution of your project and identify potential inconsistencies or issues.

```sql+postgres
select
  job_id,
  name,
  git_source ->> 'git_branch' as git_branch,
  git_source ->> 'git_commit' as git_commit,
  git_source ->> 'git_provider' as git_provider,
  git_source ->> 'git_snapshot' as git_snapshot,
  git_source ->> 'git_tag' as git_tag,
  git_source ->> 'git_url' as git_url,
  account_id
from
  databricks_job
where
  git_source is not null;
```

```sql+sqlite
select
  job_id,
  name,
  json_extract(git_source, '$.git_branch') as git_branch,
  json_extract(git_source, '$.git_commit') as git_commit,
  json_extract(git_source, '$.git_provider') as git_provider,
  json_extract(git_source, '$.git_snapshot') as git_snapshot,
  json_extract(git_source, '$.git_tag') as git_tag,
  json_extract(git_source, '$.git_url') as git_url,
  account_id
from
  databricks_job
where
  git_source is not null;
```

### Get clusters on which each job runs
Explore which clusters are utilized by each job, allowing for a clear understanding of resource allocation and potential bottlenecks in your system. This aids in optimizing job distribution and resource management for improved system performance.

```sql+postgres
select
  job_id,
  name,
  jc ->> 'job_cluster_key' as job_cluster_key,
  jc -> 'new_cluster' ->> 'cluster_name' as new_cluster_name,
  jc -> 'new_cluster' ->> 'cluster_source' as new_cluster_source,
  account_id
from
  databricks_job,
  jsonb_array_elements(job_clusters) as jc
where
  job_clusters is not null;
```

```sql+sqlite
select
  job_id,
  name,
  json_extract(jc.value, '$.job_cluster_key') as job_cluster_key,
  json_extract(jc.value, '$.new_cluster.cluster_name') as new_cluster_name,
  json_extract(jc.value, '$.new_cluster.cluster_source') as new_cluster_source,
  account_id
from
  databricks_job,
  json_each(job_clusters) as jc
where
  job_clusters is not null;
```

### Get all scheduled jobs
Explore which tasks are programmed to run automatically, including their status and timing details. This is useful to understand the operational flow and scheduling within your Databricks environment.

```sql+postgres
select
  job_id,
  name,
  schedule ->> 'pause_status' as pause_status,
  schedule ->> 'quartz_cron_expression' as quartz_cron_expression,
  schedule ->> 'timezone_id' as timezone_id,
  account_id
from
  databricks_job
where
  schedule is not null;
```

```sql+sqlite
select
  job_id,
  name,
  json_extract(schedule, '$.pause_status') as pause_status,
  json_extract(schedule, '$.quartz_cron_expression') as quartz_cron_expression,
  json_extract(schedule, '$.timezone_id') as timezone_id,
  account_id
from
  databricks_job
where
  schedule is not null;
```

### Get task details for each job
Determine the specifics of each task for every job to understand the configuration and settings in place, aiding in task management and optimization.

```sql+postgres
select
  job_id,
  name,
  t ->> 'task_key' as task_key,
  t ->> 'notebook_task' as notebook_task,
  t ->> 'timeout_seconds' as timeout_seconds,
  t ->> 'email_notifications' as email_notifications,
  t ->> 'existing_cluster_id' as existing_cluster_id,
  t ->> 'new_cluster' as new_cluster,
  t ->> 'notification_settings' as notification_settings,
  t ->> 'min_retry_interval_millis' as min_retry_interval_millis,
  t ->> 'depends_on' as depends_on,
  account_id
from
  databricks_job,
  jsonb_array_elements(tasks) as t
where
  tasks is not null;
```

```sql+sqlite
select
  job_id,
  name,
  json_extract(t.value, '$.task_key') as task_key,
  json_extract(t.value, '$.notebook_task') as notebook_task,
  json_extract(t.value, '$.timeout_seconds') as timeout_seconds,
  json_extract(t.value, '$.email_notifications') as email_notifications,
  json_extract(t.value, '$.existing_cluster_id') as existing_cluster_id,
  json_extract(t.value, '$.new_cluster') as new_cluster,
  json_extract(t.value, '$.notification_settings') as notification_settings,
  json_extract(t.value, '$.min_retry_interval_millis') as min_retry_interval_millis,
  json_extract(t.value, '$.depends_on') as depends_on,
  account_id
from
  databricks_job,
  json_each(tasks) as t
where
  tasks is not null;
```

### Get task trigger settings
Analyze the settings to understand the trigger conditions for specific tasks in Databricks. This can be useful in assessing task performance and identifying potential issues related to task triggers.

```sql+postgres
select
  job_id,
  name,
  trigger ->> 'file_arrival' as file_arrival,
  trigger ->> 'pause_status' as pause_status,
  account_id
from
  databricks_job
where
  trigger is not null;
```

```sql+sqlite
select
  job_id,
  name,
  json_extract(trigger, '$.file_arrival') as file_arrival,
  json_extract(trigger, '$.pause_status') as pause_status,
  account_id
from
  databricks_job
where
  trigger is not null;
```

### Get task trigger history
Explore the history of task triggers to determine instances where tasks have failed, not been triggered, or have been successfully triggered in the Databricks environment. This aids in understanding and troubleshooting job execution patterns and anomalies.

```sql+postgres
select
  job_id,
  name,
  trigger_history ->> 'last_failed' as last_failed,
  trigger_history ->> 'last_not_triggered' as last_not_triggered,
  trigger_history ->> 'last_triggered' as last_triggered,
  account_id
from
  databricks_job
where
  trigger_history is not null;
```

```sql+sqlite
select
  job_id,
  name,
  json_extract(trigger_history, '$.last_failed') as last_failed,
  json_extract(trigger_history, '$.last_not_triggered') as last_not_triggered,
  json_extract(trigger_history, '$.last_triggered') as last_triggered,
  account_id
from
  databricks_job
where
  trigger_history is not null;
```

### Get the permissions associated to each job
Discover the level of access granted to various users and groups for different tasks. This is particularly useful for auditing and managing security within your organization.

```sql+postgres
select
  job_id,
  name,
  acl ->> 'user_name' as principal_user_name,
  acl ->> 'group_name' as principal_group_name,
  acl ->> 'all_permissions' as permission_level
from
  databricks_job,
  jsonb_array_elements(job_permissions -> 'access_control_list') as acl;
```

```sql+sqlite
select
  job_id,
  name,
  json_extract(acl.value, '$.user_name') as principal_user_name,
  json_extract(acl.value, '$.group_name') as principal_group_name,
  json_extract(acl.value, '$.all_permissions') as permission_level
from
  databricks_job,
  json_each(job_permissions, '$.access_control_list') as acl;
```

### Find the account with the most jobs
Determine the account that has the highest job count. This is useful for identifying the most active account in terms of job execution, which can aid in resource allocation and load balancing.

```sql+postgres
select
  account_id,
  count(*) as job_count
from
  databricks_job
group by
  account_id
order by
  job_count desc
limit 1;
```

```sql+sqlite
select
  account_id,
  count(*) as job_count
from
  databricks_job
group by
  account_id
order by
  job_count desc
limit 1;
```

### List the collection of system notification IDs associated to a particular job
Explore which system notifications are linked to a specific job to better manage and track job progress and alerts. This can be useful in understanding the communication flow and debugging issues related to job notifications.

```sql+postgres
select
  job_id,
  name,
  jsonb_pretty(webhook_notifications) as notification_ids
from
  databricks_job;
```

```sql+sqlite
select
  job_id,
  name,
  webhook_notifications as notification_ids
from
  databricks_job;
```