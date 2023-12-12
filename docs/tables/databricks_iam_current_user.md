---
title: "Steampipe Table: databricks_iam_current_user - Query Databricks IAM Users using SQL"
description: "Allows users to query Databricks IAM Users, specifically the details of the current user, providing insights into user-specific data and potential security configurations."
---

# Table: databricks_iam_current_user - Query Databricks IAM Users using SQL

Databricks Identity and Access Management (IAM) is a web service that helps you securely control access to Databricks resources. It provides authentication, authorization, and audit for your Databricks environment. IAM ensures that only authenticated and authorized users are able to access your Databricks resources.

## Table Usage Guide

The `databricks_iam_current_user` table provides insights into the current user within Databricks IAM. As a security analyst, explore user-specific details through this table, including user roles, permissions, and associated metadata. Utilize it to uncover information about the current user, such as their roles and permissions, to verify security configurations and ensure adherence to best practices.

## Examples

### Basic info
Explore the current user's basic information in Databricks to gain insights into their activity status and associated account. This can be useful for monitoring user activity and managing user-related issues.

```sql+postgres
select
  id,
  user_name,
  display_name,
  active,
  account_id
from
  databricks_iam_current_user;
```

```sql+sqlite
select
  id,
  user_name,
  display_name,
  active,
  account_id
from
  databricks_iam_current_user;
```

### List assigned roles for the user
Explore which roles are currently assigned to a user on Databricks. This is useful for auditing user permissions and ensuring appropriate access levels are maintained.

```sql+postgres
select
  u.id,
  u.user_name,
  u.display_name,
  r ->> 'value' as role,
  r ->> 'type' as type,
  u.account_id
from
  databricks_iam_current_user u,
  jsonb_array_elements(roles) as r;
```

```sql+sqlite
select
  u.id,
  u.user_name,
  u.display_name,
  json_extract(r.value, '$.value') as role,
  json_extract(r.value, '$.type') as type,
  u.account_id
from
  databricks_iam_current_user u,
  json_each(roles) as r;
```

### List groups the user belongs to
Explore which groups a particular user is associated with, to understand their permissions and roles within an organization. This is useful for auditing user access and ensuring appropriate security measures are in place.

```sql+postgres
select
  u.id,
  u.user_name,
  u.display_name,
  g.id as group_id,
  g.display_name as group_name,
  u.account_id
from
  databricks_iam_current_user u,
  databricks_iam_account_group g,
  jsonb_array_elements(g.members) m
where
  m ->> 'value' = u.id
  and g.account_id = u.account_id;
```

```sql+sqlite
select
  u.id,
  u.user_name,
  u.display_name,
  g.id as group_id,
  g.display_name as group_name,
  u.account_id
from
  databricks_iam_current_user u,
  databricks_iam_account_group g,
  json_each(g.members) as m
where
  json_extract(m.value, '$.value') = u.id
  and g.account_id = u.account_id;
```

### List user's entitlements
Explore which entitlements are associated with each user in your Databricks environment. This can be beneficial to understand user permissions and roles for better access management.

```sql+postgres
select
  u.id,
  u.user_name,
  u.display_name,
  r ->> 'value' as entitlement,
  u.account_id
from
  databricks_iam_current_user u,
  jsonb_array_elements(entitlements) as r;
```

```sql+sqlite
select
  u.id,
  u.user_name,
  u.display_name,
  json_extract(r.value, '$.value') as entitlement,
  u.account_id
from
  databricks_iam_current_user u,
  json_each(u.entitlements) as r;
```

### Find the account with the most users
Pinpoint the specific account that has the highest number of users. This can be useful in identifying the most active account and understanding user distribution across accounts.

```sql+postgres
select
  account_id,
  count(*) as user_count
from
  databricks_iam_current_user
group by
  account_id
order by
  user_count desc
limit 1;
```

```sql+sqlite
select
  account_id,
  count(*) as user_count
from
  databricks_iam_current_user
group by
  account_id
order by
  user_count desc
limit 1;
```

### List users with multiple email IDs
Explore which users have registered multiple email IDs to their accounts. This can help in identifying instances where a single user might be using multiple accounts, which could be a sign of potential misuse or a security concern.

```sql+postgres
select
  id,
  user_name,
  display_name,
  active,
  account_id,
  jsonb_pretty(emails) as email_ids
from
  databricks_iam_current_user
where
  jsonb_array_length(emails) > 1;
```

```sql+sqlite
select
  id,
  user_name,
  display_name,
  active,
  account_id,
  emails as email_ids
from
  databricks_iam_current_user
where
  json_array_length(emails) > 1;
```