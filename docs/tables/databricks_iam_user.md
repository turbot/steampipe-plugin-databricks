# Table: databricks_iam_user

User identities recognized by Databricks and represented by email addresses.

## Examples

### Basic info

```sql
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

```sql
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

### List users and their primary emails

```sql
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

### List users and their work emails

```sql
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

### List assigned roles for each user

```sql
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

### List groups each user belongs to

```sql
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

### Get user with a specific user name

```sql
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

```sql
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

### Find the account with the most users

```sql
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