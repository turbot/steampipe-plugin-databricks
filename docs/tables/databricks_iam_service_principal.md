# Table: databricks_iam_service_principal

Identities for use with jobs, automated tools, and systems such as scripts, apps, and CI/CD platforms.

## Examples

### Basic info

```sql
select
  id,
  display_name,
  active,
  application_id,
  account_id
from
  databricks_iam_service_principal;
```

### List all inactive service principals

```sql
select
  id,
  display_name,
  active,
  application_id,
  account_id
from
  databricks_iam_service_principal
where
  not active;
```

### List assigned roles for each service principal

```sql
select
  u.id,
  u.display_name,
  r ->> 'value' as role,
  r ->> 'type' as type,
  u.account_id
from
  databricks_iam_service_principal u,
  jsonb_array_elements(roles) as r;
```

### List groups each service principal belongs to

```sql
select
  u.id,
  u.display_name,
  r ->> 'display' as group_display_name,
  r ->> 'value' as role,
  r ->> 'type' as type,
  u.account_id
from
  databricks_iam_service_principal u,
  jsonb_array_elements(groups) as r;
```

### Get service principal with a specific name

```sql
select
  id,
  display_name,
  active,
  account_id
from
  databricks_iam_service_principal
where
  display_name = 'user@turbot.com';
```

### List service principal entitlements

```sql
select
  id,
  display_name,
  r ->> 'value' as entitlement,
  u.account_id
from
  databricks_iam_service_principal u,
  jsonb_array_elements(entitlements) as r;
```
