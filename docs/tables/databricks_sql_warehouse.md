# Table: databricks_sql_warehouse

A SQL warehouse is a compute resource that lets you run SQL commands on data objects within Databricks SQL. Compute resources are infrastructure resources that provide processing capabilities in the cloud.

## Examples

### Basic info

```sql
select
  id,
  name,
  cluster_size,
  creator_name,
  jdbc_url,
  account_id
from
  databricks_sql_warehouse;
```

### Get cluster configuration for each warehouse

```sql
select
  id,
  name,
  cluster_size,
  enable_serverless_compute,
  max_num_clusters,
  min_num_clusters,
  num_active_sessions,
  num_clusters,
  spot_instance_policy,
  account_id
from
  databricks_sql_warehouse;
```

### List all stopped warehouse objects

```sql
select
  id,
  name,
  cluster_size,
  creator_name,
  jdbc_url,
  account_id
from
  databricks_sql_warehouse
where
  state = 'STOPPED';
```

### List unhealthy warehouse objects

```sql
select
  id,
  name,
  health ->> 'details' as health_details,
  health ->> 'failure_reason' as failure_reason,
  health ->> 'status' as health_status,
  health ->> 'summary' as health_summary,
  account_id
from
  databricks_sql_warehouse
where
  health ->> 'status' in ('FAILED', 'DEGRADED');
```

### Get count of warehouse objects by type

```sql
select
  warehouse_type,
  count(*) as count
from
  databricks_sql_warehouse
group by
  warehouse_type;
```

### Get warehouse odbc parameters

```sql
select
  id,
  name,
  odbc_params ->> 'hostname' as odbc_hostname,
  odbc_params ->> 'port' as odbc_port,
  odbc_params ->> 'path' as odbc_path,
  odbc_params ->> 'protocol' as odbc_protocol,
  account_id
from
  databricks_sql_warehouse;
```

### Get the permissions associated to each warehouse

```sql
select
  id,
  name,
  acl ->> 'user_name' as principal_user_name,
  acl ->> 'group_name' as principal_group_name,
  acl ->> 'all_permissions' as permission_level
from
  databricks_sql_warehouse,
  jsonb_array_elements(warehouse_permissions -> 'access_control_list') as acl;
```
