# Table: databricks_compute_policy_family

A policy family contains a policy definition providing best practices for configuring clusters for a particular use case. Databricks manages and provides policy families for several common cluster use cases.

## Examples

### Basic info

```sql
select
  policy_family_id,
  name,
  description,
  definition,
  account_id
from
  databricks_compute_policy_family;
```

### List cluster policies associated with policy family

```sql
select
  f.policy_family_id,
  f.name as policy_family_name,
  p.policy_id,
  p.name as policy_name,
  f.description as policy_family_description,
  f.account_id
from
  databricks_compute_policy_family f,
  databricks_compute_cluster_policy p
where
  f.policy_family_id = p.policy_family_id
  and f.account_id = p.account_id;
```

### Find the account with the most policy families

```sql
select
  account_id,
  count(*) as policy_family_count
from
  databricks_compute_policy_family
group by
  account_id
order by
  policy_family_count desc
limit 1;
```