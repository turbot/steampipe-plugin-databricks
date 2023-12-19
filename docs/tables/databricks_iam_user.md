---
title: "Steampipe Table: databricks_iam_user - Query Databricks IAM Users using SQL"
description: "Allows users to query Databricks IAM Users, providing comprehensive data about user identities, access rights, and associated metadata."
---

# Table: databricks_iam_user - Query Databricks IAM Users using SQL

Databricks IAM (Identity and Access Management) Users are entities that represent a person or service that interacts with Databricks resources. IAM Users can be assigned individual security credentials, such as access keys, passwords, and multi-factor authentication devices. They can also be assigned permissions to access Databricks resources.

## Table Usage Guide

The `databricks_iam_user` table provides insights into IAM users within Databricks Identity and Access Management (IAM). As a DevOps engineer, explore user-specific details through this table, including permissions, associated metadata, and security credentials. Utilize it to uncover information about users, such as those with extensive access rights, and to monitor and manage user access to Databricks resources.

## Examples

### Basic info
Explore which Databricks IAM users are currently active and associated with a specific account, providing a quick way to assess user activity and account associations. This can be useful in managing user access and understanding user activity patterns.

```sql+postgres
select
  id,
  user_name,
  display_name,
  active,
  account_id
from
  databricks_iam_user;
```

```sql+sqlite
select
  id,
  user_name,
  display_name,
  active,
  account_id
from
  databricks_iam_user;
```

### List all inactive users
Explore which users in your Databricks IAM are inactive, helping to maintain security by identifying potential unused or unnecessary accounts.

```sql+postgres
select
  id,
  user_name,
  display_name,
  active,
  account_id
from
  databricks_iam_user
where
  not active;
```

```sql+sqlite
select
  id,
  user_name,
  display_name,
  active,
  account_id
from
  databricks_iam_user
where
  active = 0;
```

### List users and their primary emails
Discover the segments that consist of users and their primary email addresses. This can be useful to understand user demographics and for communicating with users directly.

```sql+postgres
select
  id,
  user_name,
  display_name,
  e ->> 'value' as email,
  e ->> 'type' as type,
  account_id
from
  databricks_iam_user,
  jsonb_array_elements(emails) as e
where
  e ->> 'primary' = 'true';
```

```sql+sqlite
select
  id,
  user_name,
  display_name,
  json_extract(e.value, '$.value') as email,
  json_extract(e.value, '$.type') as type,
  account_id
from
  databricks_iam_user,
  json_each(emails) as e
where
  json_extract(e.value, '$.primary') = 'true';
```

### List users and their work emails
Explore which users have registered their work emails. This is useful for understanding the professional contact information associated with each user.

```sql+postgres
select
  id,
  user_name,
  display_name,
  e ->> 'value' as email,
  e ->> 'type' as type,
  e ->> 'primary' as is_primary,
  account_id
from
  databricks_iam_user,
  jsonb_array_elements(emails) as e
where
  e ->> 'type' = 'work';
```

```sql+sqlite
select
  id,
  user_name,
  display_name,
  json_extract(e.value, '$.value') as email,
  json_extract(e.value, '$.type') as type,
  json_extract(e.value, '$.primary') as is_primary,
  account_id
from
  databricks_iam_user,
  json_each(emails) as e
where
  json_extract(e.value, '$.type') = 'work';
```

### List assigned roles for each user
Explore the roles assigned to each user in your organization to understand their access rights and responsibilities. This can help in managing user permissions and ensuring proper security protocols.

```sql+postgres
select
  u.id,
  u.user_name,
  u.display_name,
  r ->> 'value' as role,
  r ->> 'type' as type,
  u.account_id
from
  databricks_iam_user u,
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
  databricks_iam_user u,
  json_each(u.roles) as r;
```

### List groups each user belongs to
Explore which groups each user is associated with in your Databricks IAM setup. This is useful for understanding user permissions and roles within your organization.

```sql+postgres
select
  u.id,
  u.user_name,
  u.display_name,
  r ->> 'display' as group_display_name,
  r ->> 'value' as role,
  r ->> 'type' as type,
  u.account_id
from
  databricks_iam_user u,
  jsonb_array_elements(groups) as r;
```

```sql+sqlite
select
  u.id,
  u.user_name,
  u.display_name,
  json_extract(r.value, '$.display') as group_display_name,
  json_extract(r.value, '$.value') as role,
  json_extract(r.value, '$.type') as type,
  u.account_id
from
  databricks_iam_user u,
  json_each(u.groups) as r;
```

### Get user with a specific user name
Explore which user has a specific username to understand their activity status and associated account details. This can be particularly useful in managing user access and permissions within the Databricks platform.

```sql+postgres
select
  id,
  user_name,
  display_name,
  active,
  account_id
from
  databricks_iam_user
where
  user_name = 'user@turbot.com';
```

```sql+sqlite
select
  id,
  user_name,
  display_name,
  active,
  account_id
from
  databricks_iam_user
where
  user_name = 'user@turbot.com';
```

### List user entitlements
Explore which entitlements are associated with each user in your Databricks IAM user list. This can be useful for auditing user privileges and ensuring appropriate access levels.

```sql+postgres
select
  id,
  user_name,
  display_name,
  r ->> 'value' as entitlement,
  u.account_id
from
  databricks_iam_user u,
  jsonb_array_elements(entitlements) as r;
```

```sql+sqlite
select
  u.id,
  user_name,
  display_name,
  json_extract(r.value, '$.value') as entitlement,
  u.account_id
from
  databricks_iam_user u,
  json_each(entitlements) as r;
```

### Find the account with the most users
Determine the account that has the highest number of users. This is useful for understanding which account is most utilized, potentially indicating areas of high activity or demand.

```sql+postgres
select
  account_id,
  count(*) as user_count
from
  databricks_iam_user
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
  databricks_iam_user
group by
  account_id
order by
  user_count desc
limit 1;
```