# Table: databricks_jobs_job

You can use a Databricks job to run a data processing or data analysis task in a Databricks cluster with scalable resources. Your job can consist of a single task or can be a large, multi-task workflow with complex dependencies.

## Examples

### Basic info

```sql
select
  job_id,
  name,
  created_time,
  creator_user_name,
  run_as_user_name,
  format,
  account_id
from
  databricks_jobs_job;
```

### Get compute requirements for each job

```sql
select
  job_id,
  name,
  compute ->> 'compute_key' as compute_key,
  compute ->> 'spec' as compute_spec,
  account_id
from
  databricks_jobs_job;
```

### List all continuous jobs

```sql
select
  job_id,
  name,
  format,
  continuous ->> 'pause_status' as pause_status,
  account_id
from
  databricks_jobs_job
where
  continuous is not null;
```

### Get email notification configuration for each job

```sql
select
  job_id,
  name,
  email_notifications ->> 'on_start' as email_on_start,
  email_notifications ->> 'on_success' as email_on_success,
  email_notifications ->> 'on_failure' as email_on_failure,
  email_notifications ->> 'no_alert_for_skipped_runs' as no_alert_for_skipped_runs,
  account_id
from
  databricks_jobs_job;
```

### Get git settings for each job

```sql
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
  databricks_jobs_job
where
  git_source is not null;
```

### Get clusters on which each job runs

```sql
select
  job_id,
  name,
  jc ->> 'job_cluster_key' as job_cluster_key,
  jc -> 'new_cluster' ->> 'cluster_name' as new_cluster_name,
  jc -> 'new_cluster' ->> 'cluster_source' as new_cluster_source,
  account_id
from
  databricks_jobs_job,
  jsonb_array_elements(job_clusters) as jc
where
  job_clusters is not null;
```

### Get all scheduled jobs

```sql
select
  job_id,
  name,
  schedule ->> 'pause_status' as pause_status,
  schedule ->> 'quartz_cron_expression' as quartz_cron_expression,
  schedule ->> 'timezone_id' as timezone_id,
  account_id
from
  databricks_jobs_job
where
  schedule is not null;
```

### Get task details for each job

```sql
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
  databricks_jobs_job,
  jsonb_array_elements(tasks) as t
where
  tasks is not null;
```

### Get task trigger settings

```sql
select
  job_id,
  name,
  trigger ->> 'file_arrival' as file_arrival,
  trigger ->> 'pause_status' as pause_status,
  account_id
from
  databricks_jobs_job
where
  trigger is not null;
```

### Get task trigger history

```sql
select
  job_id,
  name,
  trigger_history ->> 'last_failed' as last_failed,
  trigger_history ->> 'last_not_triggered' as last_not_triggered,
  trigger_history ->> 'last_triggered' as last_triggered,
  account_id
from
  databricks_jobs_job
where
  trigger_history is not null;
```

### Get the permissions associated to each job

```sql
select
  job_id,
  name,
  acl ->> 'user_name' as principal_user_name,
  acl ->> 'group_name' as principal_group_name,
  acl ->> 'all_permissions' as permission_level
from
  databricks_jobs_job,
  jsonb_array_elements(job_permissions -> 'access_control_list') as acl;
```

### Find the account with the most jobs

```sql
select
  account_id,
  count(*) as job_count
from
  databricks_jobs_job
group by
  account_id
order by
  job_count desc
limit 1;
```

### List the collection of system notification IDs associated to a particular job

```sql
select
  job_id,
  name,
  jsonb_pretty(webhook_notifications) as notification_ids
from
  databricks_jobs_job;
```