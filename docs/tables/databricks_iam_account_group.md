---
title: "Steampipe Table: databricks_iam_account_group - Query Databricks IAM Account Groups using SQL"
description: "Allows users to query Databricks IAM Account Groups, providing detailed information about the groups associated with a Databricks account."
---

# Table: databricks_iam_account_group - Query Databricks IAM Account Groups using SQL

Databricks IAM Account Groups represent a collection of Databricks IAM users, roles, and other groups. They are utilized to manage permissions and access to Databricks resources. Account groups streamline the process of granting and revoking access, making it easier to manage security and access control.

## Table Usage Guide

The `databricks_iam_account_group` table provides insights into IAM account groups within Databricks. As a security engineer, you can explore group-specific details through this table, including member lists, access controls, and associated metadata. Utilize it to understand the configuration of access controls, identify groups with excessive permissions, and verify the proper assignment of users and roles.

## Examples

### Basic info
Explore which account groups are associated with specific account IDs to manage and organize your Databricks IAM resources more effectively. This can help in understanding the structure of your account and its security settings.

```sql+postgres
select
  id,
  display_name,
  account_id
from
  databricks_iam_account_group;
```

```sql+sqlite
select
  id,
  display_name,
  account_id
from
  databricks_iam_account_group;
```

### List all members of a specific group
Explore which members belong to a particular group. This can be useful in managing access controls and understanding group composition within the Databricks IAM account.

```sql+postgres
select
  g.id,
  g.display_name,
  m ->> 'display' as member_display_name,
  m ->> 'value' as member_id,
  m ->> 'type' as member_type,
  g.account_id
from
  databricks_iam_account_group g,
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
  databricks_iam_account_group g,
  json_each(g.members) m
where
  g.display_name = 'dev';
```

### List all members that are users in a specific group
Discover the segments that consist of users belonging to a specific group. This is useful in managing user access and permissions in a more organized manner.

```sql+postgres
select
  g.id,
  g.display_name,
  m ->> 'display' as member_display_name,
  m ->> 'value' as member_id,
  m ->> 'type' as member_type,
  g.account_id
from
  databricks_iam_account_group g,
  jsonb_array_elements(g.members) m
where
  g.display_name = 'dev'
  and m ->> '$ref' like 'User%';
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
  databricks_iam_account_group g,
  json_each(g.members) m
where
  g.display_name = 'dev'
  and json_extract(m.value, '$.$ref') like 'User%';
```

### List all members that are groups in a specific group
This example helps you identify all the groups that are part of a specific group within your organization. It can be useful for understanding the structure and hierarchy of your group memberships.

```sql+postgres
select
  g.id,
  g.display_name,
  m ->> 'display' as member_display_name,
  m ->> 'value' as member_id,
  m ->> 'type' as member_type,
  g.account_id
from
  databricks_iam_account_group g,
  jsonb_array_elements(g.members) m
where
  g.display_name = 'dev'
  and m ->> '$ref' like 'Group%';
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
  databricks_iam_account_group g,
  json_each(g.members) m
where
  g.display_name = 'dev'
  and json_extract(m.value, '$.$ref') like 'Group%';
```

### List all the entitlements associated to a particular account group
Determine the areas in which specific account group entitlements apply, enabling the identification of access privileges for development-related tasks. This is useful for managing and monitoring access controls within your Databricks environment.

```sql+postgres
select
  id,
  display_name,
  account_id,
  jsonb_pretty(entitlements) as entitlements
from
  databricks_iam_account_group
where
  display_name = 'dev';
```

```sql+sqlite
select
  id,
  display_name,
  account_id,
  entitlements
from
  databricks_iam_account_group
where
  display_name = 'dev';
```