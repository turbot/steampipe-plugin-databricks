---
title: "Steampipe Table: databricks_job_run - Query Databricks Job Runs using SQL"
description: "Allows users to query Databricks Job Runs, specifically providing detailed information about each job run, including the job ID, run ID, start time, and state."
---

# Table: databricks_job_run - Query Databricks Job Runs using SQL

Databricks Job Runs are a resource within the Databricks service that allows users to execute a specific job on a schedule or on-demand. They provide detailed information about each execution, including the job ID, run ID, start time, and state. This helps users track the status and progress of their Databricks jobs.

## Table Usage Guide

The `databricks_job_run` table provides insights into Job Runs within Databricks. As a Data Engineer, you can explore specific details about each job run through this table, including the job ID, run ID, start time, and state. Utilize it to monitor the status and progress of your Databricks jobs, helping you to manage and optimize your data processing tasks.

## Examples

### Basic info
Explore which Databricks job runs have taken place, identifying instances such as who initiated them and how long they lasted. This can be particularly beneficial for auditing purposes or for understanding patterns in job run durations and frequencies.

```sql+postgres
select
  run_id,
  run_name,
  job_id,
  original_attempt_run_id,
  attempt_number,
  creator_user_name,
  run_duration as run_duration_ms,
  account_id
from
  databricks_job_run;
```

```sql+sqlite
select
  run_id,
  run_name,
  job_id,
  original_attempt_run_id,
  attempt_number,
  creator_user_name,
  run_duration as run_duration_ms,
  account_id
from
  databricks_job_run;
```

### Get total runs per job
Discover the frequency of each job execution to better understand the workload distribution. This could be useful in identifying heavily used jobs that may need optimization or resource allocation adjustments.

```sql+postgres
select
  job_id,
  count(*) as total_runs
from
  databricks_job_run
group by
  job_id;
```

```sql+sqlite
select
  job_id,
  count(*) as total_runs
from
  databricks_job_run
group by
  job_id;
```

### Get total runs per job per day
Explore the frequency of job runs on a daily basis to understand the workload distribution and identify any potential bottlenecks or high-activity periods. This can help in optimizing job scheduling and resource allocation.

```sql+postgres
select
  job_id,
  date_trunc('day', start_time) as day,
  count(*) as total_runs
from
  databricks_job_run
group by
  job_id,
  day
order by
  day;
```

```sql+sqlite
select
  job_id,
  date(start_time) as day,
  count(*) as total_runs
from
  databricks_job_run
group by
  job_id,
  day
order by
  day;
```

### Get the state of the last run for each job
Determine the status of the most recent job execution for each task. This can help you understand the success or failure of your latest tasks, enabling you to troubleshoot issues or optimize performance.

```sql+postgres
select
  job_id,
  run_id,
  run_name,
  attempt_number,
  state ->> 'state_message',
  state ->> 'life_cycle_state',
  state ->> 'result_state',
  account_id
from
  databricks_job_run
order by
  attempt_number desc
limit 1;
```

```sql+sqlite
select
  job_id,
  run_id,
  run_name,
  attempt_number,
  json_extract(state, '$.state_message'),
  json_extract(state, '$.life_cycle_state'),
  json_extract(state, '$.result_state'),
  account_id
from
  databricks_job_run
order by
  attempt_number desc
limit 1;
```

### Get task details for each job run
This query is used to gain insights into the details of each task within job runs, including the duration, start and end times, and state. It's useful for understanding the performance and efficiency of tasks within job runs in a Databricks environment.

```sql+postgres
select
  job_id,
  run_id,
  run_name,
  t ->> 'task_key' as task_key,
  t ->> 'cleanup_duration' as cleanup_duration,
  t ->> 'cluster_instance' as cluster_instance,
  t ->> 'start_time' as start_time,
  t ->> 'end_time' as end_time,
  t ->> 'existing_cluster_id' as existing_cluster_id,
  t ->> 'notebook_task' as notebook_task,
  t ->> 'cleanup_duration' as cleanup_duration,
  t ->> 'state' as state,
  account_id
from
  databricks_job_run,
  jsonb_array_elements(tasks) as t
where
  tasks is not null;
```

```sql+sqlite
select
  job_id,
  run_id,
  run_name,
  json_extract(t.value, '$.task_key') as task_key,
  json_extract(t.value, '$.cleanup_duration') as cleanup_duration,
  json_extract(t.value, '$.cluster_instance') as cluster_instance,
  json_extract(t.value, '$.start_time') as start_time,
  json_extract(t.value, '$.end_time') as end_time,
  json_extract(t.value, '$.existing_cluster_id') as existing_cluster_id,
  json_extract(t.value, '$.notebook_task') as notebook_task,
  json_extract(t.value, '$.cleanup_duration') as cleanup_duration,
  json_extract(t.value, '$.state') as state,
  account_id
from
  databricks_job_run,
  json_each(tasks) as t
where
  tasks is not null;
```

### List jobs that are waiting for retry
Explore which jobs are currently in a waiting state for a retry attempt. This is useful for identifying potential issues with certain jobs and understanding where intervention may be necessary.

```sql+postgres
select
  run_id,
  run_name,
  job_id,
  original_attempt_run_id,
  attempt_number,
  creator_user_name,
  run_duration as run_duration_ms,
  account_id
from
  databricks_job_run
where
  state ->> 'life_cycle_state' = 'WAITING_FOR_RETRY';
```

```sql+sqlite
select
  run_id,
  run_name,
  job_id,
  original_attempt_run_id,
  attempt_number,
  creator_user_name,
  run_duration as run_duration_ms,
  account_id
from
  databricks_job_run
where
  json_extract(state, '$.life_cycle_state') = 'WAITING_FOR_RETRY';
```

### List retry job runs for a particular job
Analyze the settings to understand the frequency and reasons for job reruns in a specific Databricks job. This can help to pinpoint areas of instability and inform decisions on system optimization.

```sql+postgres
select
  run_id,
  run_name,
  job_id,
  original_attempt_run_id,
  attempt_number,
  creator_user_name,
  run_duration as run_duration_ms,
  account_id
from
  databricks_job_run
where
  job_id = '572473586420586'
  and original_attempt_run_id <> run_id;
```

```sql+sqlite
select
  run_id,
  run_name,
  job_id,
  original_attempt_run_id,
  attempt_number,
  creator_user_name,
  run_duration as run_duration_ms,
  account_id
from
  databricks_job_run
where
  job_id = '572473586420586'
  and original_attempt_run_id <> run_id;
```