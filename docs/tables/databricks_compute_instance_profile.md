---
title: "Steampipe Table: databricks_compute_instance_profile - Query Databricks Compute Instance Profiles using SQL"
description: "Allows users to query Compute Instance Profiles in Databricks, specifically information about the instance profile ARN, role ARN, and the instance profile ID, providing insights into the configuration and status of compute instance profiles."
---

# Table: databricks_compute_instance_profile - Query Databricks Compute Instance Profiles using SQL

A Compute Instance Profile in Databricks is a predefined set of permissions that you can assign to users or groups. These profiles define what actions users or groups can perform and on what resources. Compute Instance Profiles are used to manage access to Databricks resources, including clusters, jobs, and notebooks.

## Table Usage Guide

The `databricks_compute_instance_profile` table provides insights into Compute Instance Profiles within Databricks. As a DevOps engineer, explore profile-specific details through this table, including the instance profile ARN, role ARN, and instance profile ID. Utilize it to uncover information about compute instance profiles, such as those with specific permissions, the resources they can access, and their current status.

## Examples

### Basic info
Explore which Databricks compute instances are associated with specific IAM roles and accounts. This can be useful in understanding the security configuration and access rights within your Databricks environment.

```sql+postgres
select
  instance_profile_arn,
  iam_role_arn,
  is_meta_instance_profile,
  account_id
from
  databricks_compute_instance_profile;
```

```sql+sqlite
select
  instance_profile_arn,
  iam_role_arn,
  is_meta_instance_profile,
  account_id
from
  databricks_compute_instance_profile;
```

### List all valid instance profiles
Determine the areas in which valid instance profiles are being used within your Databricks compute environment. This can help assess the elements within your account that are associated with these profiles, providing insights into your resource allocation and usage.

```sql+postgres
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

```sql+sqlite
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
Explore the association between instance profiles and clusters to understand the connection between specific IAM roles and Databricks clusters within your account. This can help in managing access control and resource allocation.

```sql+postgres
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

```sql+sqlite
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
  p.instance_profile_arn = json_extract(c.aws_attributes, '$.instance_profile_arn')
  and p.account_id = c.account_id;
```

### Get instance profile used by all SQL warehouses in a workspace
Explore the relationship between your SQL warehouses and instance profiles across your workspace. This can help you understand the different roles and permissions assigned to your warehouses, providing valuable insights for managing access and security.

```sql+postgres
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

```sql+sqlite
select
  p.instance_profile_arn,
  p.iam_role_arn,
  p.is_meta_instance_profile,
  p.account_id
from
  databricks_compute_instance_profile p
join
  databricks_sql_warehouse_config c
on
  p.instance_profile_arn = c.instance_profile_arn
  and p.account_id = c.account_id;
```