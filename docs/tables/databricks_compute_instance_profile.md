# Table: databricks_compute_instance_profile

The Instance Profiles API allows admins to add, list, and remove instance profiles that users can launch clusters with. Regular users can list the instance profiles available to them.

## Examples

### Basic info

```sql
select
  instance_profile_arn,
  iam_role_arn,
  is_meta_instance_profile,
  account_id
from
  databricks_compute_instance_profile;
```

### List all valid instance profiles

```sql
select
  instance_profile_arn,
  iam_role_arn,
  is_meta_instance_profile,
  account_id
from
  databricks_compute_instance_profile
where
  is_meta_instance_profile;
```

### List instance profiles associated with clusters

```sql
select
  p.instance_profile_arn,
  p.iam_role_arn,
  c.cluster_id,
  c.cluster_name,
  p.account_id
from
  databricks_compute_instance_profile p,
  databricks_compute_cluster c
where
  p.instance_profile_arn = c.aws_attributes ->> 'instance_profile_arn'
  and p.account_id = c.account_id;
```

### Get instance profile used by all SQL warehouses in a workspace

```sql
select
  p.instance_profile_arn,
  p.iam_role_arn,
  p.is_meta_instance_profile,
  p.account_id
from
  databricks_compute_instance_profile p,
  databricks_sql_warehouse_config c
where
  p.instance_profile_arn = c.instance_profile_arn
  and p.account_id = c.account_id;
```