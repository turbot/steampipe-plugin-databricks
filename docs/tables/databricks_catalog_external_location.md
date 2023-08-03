# Table: databricks_catalog_external_location

An external location is an object that combines a cloud storage path with a storage credential that authorizes access to the cloud storage path. Each external location is subject to Unity Catalog access-control policies that control which users and groups can access the credential.

## Examples

### Basic info

```sql
select
  name,
  comment,
  created_at,
  created_by,
  metastore_id,
  url,
  account_id
from
  databricks_catalog_external_location;
```

### List external locations modified in the last 7 days

```sql
select
  name,
  comment,
  created_at,
  created_by,
  metastore_id,
  url,
  account_id
from
  databricks_catalog_external_location
where
  updated_at >= now() - interval '7 days';
```

### List read only external locations

```sql
select
  name,
  comment,
  created_at,
  created_by,
  metastore_id,
  url,
  account_id
from
  databricks_catalog_external_location
where
  read_only;
```

### Get assocated credential for each external location

```sql
select
  l.name,
  l.comment,
  l.url,
  c.name as credential_name,
  c.id,
  c.aws_iam_role as credential_aws_iam_role,
  l.account_id
from
  databricks_catalog_external_location l,
  databricks_catalog_storage_credential c
where
  l.credential_id = c.id
  and l.account_id = c.account_id;
```

### Get effective permissions for each external location

```sql
select
  name,
  p ->> 'principal' as principal_name,
  p ->> 'privileges' as permissions
from
  databricks_catalog_external_location,
  jsonb_array_elements(external_location_effective_permissions) p;
```

### Count the number of external locations per account ID

```sql
select
  account_id,
  count(*) AS location_count
from
  databricks_catalog_external_location
group by
  account_id;
```

### List users who created the most external locations

```sql
select
  created_by,
  count(*) as location_count
from
  databricks_catalog_external_location
group by
  created_by
order by
  location_count desc;
```