# Table: databricks_sharing_share

A share is a container instantiated with Shares/Create. Once created you can iteratively register a collection of existing data assets defined within the metastore using Shares/Update.

## Examples

### Basic info

```sql
select
  name,
  comment,
  created_at,
  created_by,
  account_id
from
  databricks_sharing_share;
```

### List objects shared in past 7 days

```sql
select
  name,
  comment,
  created_at,
  created_by,
  account_id
from
  databricks_sharing_share
where
  created_at > now() - interval '7' day;
```

### List all shared objects

```sql
select
  name as share_name,
  o ->> 'name' as object_name,
  o ->> 'added_at' as added_at,
  o ->> 'added_by' as added_by,
  o ->> 'data_object_type' as data_object_type,
  o ->> 'shared_as' as shared_as,
  o ->> 'status' as status,
  account_id
from
  databricks_sharing_share,
  jsonb_array_elements(objects) as o
```

### Get permissions for each share

```sql
select
  name,
  p ->> 'principal' as principal_name,
  p ->> 'privileges' as permissions
from
  databricks_sharing_share,
  jsonb_array_elements(permissions) p;
```
