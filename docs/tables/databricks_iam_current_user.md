# Table: databricks_iam_current_user

Information about currently authenticated user or service principal

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
  databricks_iam_current_user;
```

### List assigned roles for the user

```sql
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

### List groups the user belongs to

```sql
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

### List user's entitlements

```sql
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