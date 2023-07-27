# Table: databricks_iam_account_group

Groups simplify identity management, making it easier to assign access to Databricks account, data, and other securable objects.

## Examples

### Basic info

```sql
select
  id,
  display_name,
  account_id
from
  databricks_iam_account_group;
```

### List all members of a specific group

```sql
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
```

### List all members that are users in a specific group

```sql
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
  and m ->> '$ref' like 'User%'
```

### List all members that are groups in a specific group

```sql
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
  and m ->> '$ref' like 'Group%'
```