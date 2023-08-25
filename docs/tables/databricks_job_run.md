# Table: databricks_job_run

You can use a Databricks job to run a data processing or data analysis task in a Databricks cluster with scalable resources. Your job can consist of a single task or can be a large, multi-task workflow with complex dependencies. Job run is an instance of a job that is triggered by a scheduler or manually.

## Examples

### Basic info

```sql
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

```sql
select
  job_id,
  count(*) as total_runs
from
  databricks_job_run
group by
  job_id;
```

### Get total runs per job per day

```sql
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

### Get the state of the last run for each job

```sql
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

### Get task details for each job run

```sql
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

### List jobs that are waiting for retry

```sql
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

### List retry job runs for a particular job

```sql
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