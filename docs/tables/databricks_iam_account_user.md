# Table: databricks_iam_account_user

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
  databricks_iam_account_user;
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
  databricks_iam_account_user
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
  databricks_iam_account_user,
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
  databricks_iam_account_user,
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
  databricks_iam_account_user u,
  jsonb_array_elements(roles) as r;
```

### List groups each user belongs to

```sql
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

### Get user with a specific user name

```sql
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