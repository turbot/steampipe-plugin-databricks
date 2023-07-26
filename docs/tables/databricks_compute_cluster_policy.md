# Table: databricks_compute_cluster_policy

Cluster policy limits the ability to configure clusters based on a set of rules. The policy rules limit the attributes or attribute values available for cluster creation. Cluster policies have ACLs that limit their use to specific users and groups

## Examples

### Basic info

```sql
select
  name,
  policy_id,
  created_at_timestamp,
  creator_user_name,
  description,
  account_id
from
  databricks_compute_cluster_policy;
```

### List all default policies

```sql
select
  name,
  policy_id,
  created_at_timestamp,
  creator_user_name,
  description,
  account_id
from
  databricks_compute_cluster_policy
where
  is_default;
```

### List policies having no limit on the number of active clusters using it

```sql
select
  name,
  policy_id,
  created_at_timestamp,
  creator_user_name,
  description,
  account_id
from
  databricks_compute_cluster_policy
where
  max_clusters_per_user is null;
```

### Get the ACLs for the policies

```sql
select
  name,
  policy_id,
  created_at_timestamp,
  acl ->> 'user_name' as principal_user_name,
  acl ->> 'group_name' as principal_group_name,
  acl ->> 'all_permissions' as permission_level
from
  databricks_compute_cluster_policy,
  jsonb_array_elements(definition -> 'access_control_list') as acl;
