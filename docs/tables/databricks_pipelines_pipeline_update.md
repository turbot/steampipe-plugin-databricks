# Table: databricks_pipelines_pipeline_update

Delta Live Tables is a framework for building reliable, maintainable, and testable data processing pipelines. You define the transformations to perform on your data, and Delta Live Tables manages task orchestration, cluster management, monitoring, data quality, and error handling. Pipeline updates are the pipeline update logs.

## Examples

### Basic info

```sql
select
  update_id,
  pipeline_id,
  cause,
  cluster_id,
  creation_time,
  account_id
from
  databricks_pipelines_pipeline_update;
```

### List updates created in the last 7 days

```sql
select
  update_id,
  pipeline_id,
  cause,
  cluster_id,
  creation_time,
  account_id
from
  databricks_pipelines_pipeline_update
where
  creation_time >= now() - interval '7' day;
```

### List updates caused by an API call

```sql
select
  update_id,
  pipeline_id,
  cause,
  cluster_id,
  creation_time,
  account_id
from
  databricks_pipelines_pipeline_update
where
  cause = 'API_CALL';
```

### List all failed updates

```sql
select
  update_id,
  pipeline_id,
  cause,
  cluster_id,
  creation_time,
  account_id
from
  databricks_pipelines_pipeline_update
where
  state = 'FAILED';
```

### List all pipelines that require full refresh before each run

```sql
select
  update_id,
  pipeline_id,
  cause,
  cluster_id,
  creation_time,
  full_refresh_selection,
  account_id
from
  databricks_pipelines_pipeline_update
where
  full_refresh;
```

### Find the account with the most pipeline updates

```sql
select
  account_id,
  count(*) as update_count
from
  databricks_pipelines_pipeline_update
group by
  account_id
order by
  update_count desc
limit 1;
```