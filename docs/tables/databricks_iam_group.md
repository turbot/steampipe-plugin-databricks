---
title: "Steampipe Table: databricks_iam_group - Query Databricks IAM Groups using SQL"
description: "Allows users to query IAM Groups in Databricks, providing insights into group details, membership, and associated roles."
---

# Table: databricks_iam_group - Query Databricks IAM Groups using SQL

Databricks IAM Groups is a feature within Databricks that allows you to manage access to Databricks resources. It provides a way to create and manage groups, add and remove users from groups, and assign IAM roles to groups. Databricks IAM Groups help you manage user access and permissions in a centralized and efficient manner.

## Table Usage Guide

The `databricks_iam_group` table provides insights into IAM groups within Databricks. As a Security Analyst or DevOps engineer, explore group-specific details through this table, including group membership and associated roles. Utilize it to manage user access, ensure proper permissions are assigned, and maintain security compliance.

## Examples

### Basic info
Explore which user groups exist within your Databricks account, including their unique identifiers and display names, to manage and control access effectively.

```sql+postgres
select
  id,
  display_name,
  account_id
from
  databricks_iam_group;
```

```sql+sqlite
select
  id,
  display_name,
  account_id
from
  databricks_iam_group;
```

### List all members of a specific group
Identify instances where you can explore the composition of a specific group in your organization. This is useful for gaining insights into group structure and membership, aiding in management and oversight.

```sql+postgres
select
  g.id,
  g.display_name,
  m ->> 'display' as member_display_name,
  m ->> 'value' as member_id,
  m ->> 'type' as member_type,
  g.account_id
from
  databricks_iam_group g,
  jsonb_array_elements(g.members) m
where
  g.display_name = 'dev';
```

```sql+sqlite
select
  g.id,
  g.display_name,
  json_extract(m.value, '$.display') as member_display_name,
  json_extract(m.value, '$.value') as member_id,
  json_extract(m.value, '$.type') as member_type,
  g.account_id
from
  databricks_iam_group g,
  json_each(g.members) m
where
  g.display_name = 'dev';
```

### List all groups in a specific group
Analyze the settings to understand the hierarchical relationships within a specific user group, such as the 'admin' group. This query is beneficial in managing and understanding user roles and permissions within a specific group.

```sql+postgres
select
  g.id,
  g.display_name,
  m ->> 'display' as group_display_name,
  m ->> 'value' as group_id,
  m ->> 'type' as group_type,
  g.account_id
from
  databricks_iam_group g,
  jsonb_array_elements(g.groups) m
where
  g.display_name = 'admin';
```

```sql+sqlite
select
  g.id,
  g.display_name,
  json_extract(m.value, '$.display') as group_display_name,
  json_extract(m.value, '$.value') as group_id,
  json_extract(m.value, '$.type') as group_type,
  g.account_id
from
  databricks_iam_group g,
  json_each(g.groups) m
where
  g.display_name = 'admin';
```

### List group entitlements
Determine the areas in which specific group entitlements are being used within your Databricks IAM setup. This allows you to better manage access and permissions across your organization.

```sql+postgres
select
  u.id,
  u.display_name,
  r ->> 'value' as entitlement,
  u.account_id
from
  databricks_iam_group u,
  jsonb_array_elements(entitlements) as r;
```

```sql+sqlite
select
  u.id,
  u.display_name,
  json_extract(r.value, '$.value') as entitlement,
  u.account_id
from
  databricks_iam_group u,
  json_each(entitlements) as r;
```

### List all workspace local groups
Discover the segments that consist of local groups within a workspace. This is useful for understanding the organization and access control within your Databricks environment.

```sql+postgres
select
  id,
  display_name,
  account_id
from
  databricks_iam_group
where
  meta ->> 'resourceType' = 'WorkspaceGroup';
```

```sql+sqlite
select
  id,
  display_name,
  account_id
from
  databricks_iam_group
where
  json_extract(meta, '$.resourceType') = 'WorkspaceGroup';
```

### Find the account with the most groups
Uncover the details of the account that is associated with the highest number of groups. This can be useful in understanding the distribution and organization of groups within your network.

```sql+postgres
select
  account_id,
  count(*) as group_count
from
  databricks_iam_group
group by
  account_id
order by
  group_count desc
limit 1;
```

```sql+sqlite
select
  account_id,
  count(*) as group_count
from
  databricks_iam_group
group by
  account_id
order by
  group_count desc
limit 1;
```

### List groups assigned with multiple roles
Explore which groups have been assigned more than one role in Databricks IAM, allowing a better understanding of permission distribution and potential security risks.

```sql+postgres
select
  id,
  display_name,
  account_id,
  jsonb_pretty(roles) as iam_group_roles
from
  databricks_iam_group
where
  jsonb_array_length(roles) > 1;
```

```sql+sqlite
select
  id,
  display_name,
  account_id,
  roles as iam_group_roles
from
  databricks_iam_group
where
  json_array_length(roles) > 1;
```