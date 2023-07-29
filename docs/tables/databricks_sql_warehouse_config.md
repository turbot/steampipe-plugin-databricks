# Table: databricks_sql_warehouse_config

A SQL warehouse config is the workspace level configuration that is shared by all SQL warehouses in a workspace.

## Examples

### Get warehouse sql configuration

```sql
select
  security_policy,
  cp ->> 'key' as config_parameter_key,
  cp ->> 'value' as config_parameter_value,
  account_id
from
  databricks_sql_warehouse_config,
  jsonb_array_elements(sql_configuration_parameters -> 'configuration_pairs') as cp;
```

### Check if warehouse config uses a security policy

```sql
select
  security_policy,
  google_service_account,
  instance_profile_arn,
  account_id
from
  databricks_sql_warehouse_config
where
  security_policy <> 'NONE';
```

### Get data acces configuration

```sql
select
  security_policy,
  ac ->> 'key' as config_parameter_key,
  ac ->> 'value' as config_parameter_value,
  account_id
from
  databricks_sql_warehouse_config,
  jsonb_array_elements(data_access_config) as ac;
```

### Get all enabled warehouse types for the workspace

```sql
select
  security_policy,
  wt ->> 'warehouse_type' as warehouse_type,
  wt ->> 'enabled' as enabled,
  account_id
from
  databricks_sql_warehouse_config,
  jsonb_array_elements(enabled_warehouse_types) as wt
where
  wt ->> 'enabled' = 'true';
```