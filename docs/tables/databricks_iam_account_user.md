---
title: "Steampipe Table: databricks_iam_account_user - Query Databricks IAM Account Users using SQL"
description: "Allows users to query IAM Account Users in Databricks, specifically the user details such as login name, home directory, and last login time."
---

# Table: databricks_iam_account_user - Query Databricks IAM Account Users using SQL

Databricks Identity and Access Management (IAM) is a service that helps manage access to Databricks resources. It allows you to control who is authenticated (signed in) and authorized (has permissions) to use resources. IAM Account Users are individual user identities in Databricks that can be given permission to access Databricks resources.

## Table Usage Guide

The `databricks_iam_account_user` table provides insights into individual user identities within Databricks IAM. As a system administrator, explore user-specific details through this table, including login name, home directory, and last login time. Utilize it to uncover information about users, such as their access permissions, the resources they can access, and the frequency of their logins.

## Examples

### Basic info
Explore which Databricks IAM account users are active, by displaying their user names and IDs. This can be used to manage account access and maintain security within your Databricks environment.

```sql+postgres
select
  id,
  user_name,
  display_name,
  active,
  account_id
from
  databricks_iam_account_user;
```

```sql+sqlite
select
  id,
  user_name,
  display_name,
  active,
  account_id
from
  databricks_iam_account_user;
```

### List all inactive users
Explore which users in your Databricks account are inactive to help manage resources and maintain security. This is particularly useful for administrators who need to keep track of user activity and status.

```sql+postgres
select
  id,
  user_name,
  display_name,
  active,
  account_id
from
  databricks_iam_account_user
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
  databricks_iam_account_user
where
  not active;
```

### List all the entitlements associated to a particular user
Explore which entitlements are linked to a specific user within a Databricks account. This insight can be useful for auditing user permissions and ensuring appropriate access levels.

```sql+postgres
select
  id,
  display_name,
  account_id,
  jsonb_pretty(entitlements) as entitlements
from
  databricks_iam_account_user
where
  display_name = 'abc-user';
```

```sql+sqlite
select
  id,
  display_name,
  account_id,
  entitlements
from
  databricks_iam_account_user
where
  display_name = 'abc-user';
```

### List users and their primary emails
Gain insights into the primary email addresses associated with each user. This is useful for understanding the main point of contact for each individual in your system.

```sql+postgres
select
  id,
  user_name,
  display_name,
  e ->> 'value' as email,
  e ->> 'type' as type,
  account_id
from
  databricks_iam_account_user,
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
  databricks_iam_account_user,
  json_each(emails) as e
where
  json_extract(e.value, '$.primary') = 'true';
```

### List users and their work emails
Determine the work emails associated with each user in a Databricks account. This can be useful for administrators who need to manage or communicate with users based on their professional contact information.

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
  databricks_iam_account_user,
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
  databricks_iam_account_user,
  json_each(emails) as e
where
  json_extract(e.value, '$.type') = 'work';
```

### List assigned roles for each user
Explore which roles are assigned to each user in your Databricks IAM account. This is particularly useful for auditing purposes, ensuring users have the correct permissions and identifying any potential security risks.

```sql+postgres
select
  u.id,
  u.user_name,
  u.display_name,
  r ->> 'value' as role,
  r ->> 'type' as type,
  u.account_id
from
  databricks_iam_account_user u,
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
  databricks_iam_account_user u,
  json_each(u.roles) as r;
```

### List groups each user belongs to
Determine the areas in which each user is associated by identifying the groups they belong to. This is useful for managing user permissions and understanding user roles within an organization.

```sql+postgres
select
  u.id,
  u.user_name,
  u.display_name,
  g.id as group_id,
  g.display_name as group_name,
  u.account_id
from
  databricks_iam_account_user u,
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
  databricks_iam_account_user u,
  databricks_iam_account_group g,
  json_each(g.members) as m
where
  json_extract(m.value, '$.value') = u.id
  and g.account_id = u.account_id;
```

### Get user with a specific user name
Explore which user accounts are associated with a specific username. This can be particularly useful in managing user permissions and roles, or investigating potential security issues.

```sql+postgres
select
  id,
  user_name,
  display_name,
  active,
  account_id
from
  databricks_iam_account_user
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
  databricks_iam_account_user
where
  user_name = 'user@turbot.com';
```

### Find the account with the most users
Explore which account has the highest number of users. This is useful for identifying potential areas of heavy traffic or resource usage.

```sql+postgres
select
  account_id,
  count(*) as user_count
from
  databricks_iam_account_user
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
  databricks_iam_account_user
group by
  account_id
order by
  user_count desc
limit 1;
```

### List users with multiple email IDs
Discover the segments that include users with more than one email ID. This is particularly useful for managing user accounts and ensuring data integrity within your Databricks IAM account.

```sql+postgres
select
  id,
  user_name,
  display_name,
  active,
  account_id,
  jsonb_pretty(emails) as email_ids
from
  databricks_iam_account_user
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
  databricks_iam_account_user
where
  json_array_length(emails) > 1;
```