# Table: databricks_iam_group

Groups simplify identity management, making it easier to assign access to Databricks workspace, data, and other securable objects.

## Examples

### Basic info

```sql
select
  id,
  display_name,
  account_id
from
  databricks_iam_group;
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
  databricks_iam_group g,
  jsonb_array_elements(g.members) m
where
  g.display_name = 'dev'
```

### List all groups in a specific group

```sql
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

### List group entitlements

```sql
select
  u.id,
  u.display_name,
  r ->> 'value' as entitlement,
  u.account_id
from
  databricks_iam_group u,
  jsonb_array_elements(entitlements) as r;
```

### List all workspace local groups

```sql
select
  id,
  display_name,
  account_id
from
  databricks_iam_group
where
  meta ->> 'resourceType' = 'WorkspaceGroup';
```