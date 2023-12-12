---
title: "Steampipe Table: databricks_compute_cluster_policy - Query Databricks Compute Cluster Policies using SQL"
description: "Allows users to query Databricks Compute Cluster Policies, providing insights into the specifications and restrictions for the clusters that users can create."
---

# Table: databricks_compute_cluster_policy - Query Databricks Compute Cluster Policies using SQL

A Databricks Compute Cluster Policy is a feature within Databricks that allows administrators to manage the specifications and restrictions for the clusters that users can create. It provides a centralized way to control the resources usage, including virtual machines, databases, and more. Databricks Compute Cluster Policy helps you maintain cost and resource utilization by enforcing predefined conditions for cluster creation.

## Table Usage Guide

The `databricks_compute_cluster_policy` table provides insights into Compute Cluster Policies within Databricks. As a DevOps engineer, explore policy-specific details through this table, including permissions, restrictions, and associated metadata. Utilize it to uncover information about policies, such as those with specific resource restrictions, the relationships between policies and clusters, and the verification of policy conditions.

## Examples

### Basic info
Explore the creation details and associated information of compute cluster policies in Databricks to understand their origin and usage. This is useful for auditing and managing resource allocation policies in your Databricks environment.

```sql+postgres
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

```sql+sqlite
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

### List clusters created in the last 7 days
Gain insights into the recently created clusters within the past week to understand their configurations and creators. This can be beneficial for tracking the usage and growth of your Databricks environment.

```sql+postgres
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
  created_at_timestamp >= now() - interval '7 days';
```

```sql+sqlite
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
  created_at_timestamp >= datetime('now', '-7 days');
```

### List all default policies
Explore which policies are set as default in your Databricks compute cluster. This is useful to understand the standard configurations applied across your account and to identify any potential security or performance implications.

```sql+postgres
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

```sql+sqlite
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
Discover the policies that have no restrictions on the number of active clusters using them. This can be useful in managing resource allocation and identifying potential areas of system overload.

```sql+postgres
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

```sql+sqlite
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
Explore the access control levels (ACLs) associated with various policies to understand who has what level of permissions. This can be useful for maintaining security and ensuring appropriate access rights within your Databricks compute cluster.

```sql+postgres
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
```

```sql+sqlite
select
  name,
  policy_id,
  created_at_timestamp,
  json_extract(acl.value, '$.user_name') as principal_user_name,
  json_extract(acl.value, '$.group_name') as principal_group_name,
  json_extract(acl.value, '$.all_permissions') as permission_level
from
  databricks_compute_cluster_policy,
  json_each(definition, '$.access_control_list') as acl;
```

### Find the account with the most cluster policies
Discover the account that has the highest number of cluster policies. This query can be used to identify potential areas of policy concentration or overload within an account.

```sql+postgres
select
  account_id,
  count(*) as policy_count
from
  databricks_compute_cluster_policy
group by
  account_id
order by
  policy_count desc
limit 1;
```

```sql+sqlite
select
  account_id,
  count(*) as policy_count
from
  databricks_compute_cluster_policy
group by
  account_id
order by
  policy_count desc
limit 1;
```