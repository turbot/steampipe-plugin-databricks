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

### List updated caused by an api call

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

### List all pipelines requiring full refresh before each run

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
