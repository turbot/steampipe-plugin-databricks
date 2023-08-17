# Table: databricks_catalog_metastore

A metastore is the top-level container of objects in Unity Catalog. It stores data assets (tables and views) and the permissions that govern access to them. Databricks account admins can create metastores and assign them to Databricks workspaces to control which workloads use each metastore.

## Examples

### Basic info

```sql
select
  metastore_id,
  name,
  cloud,
  created_at,
  owner,
  account_id
from
  databricks_catalog_metastore;
```

### List cloud provider configuration for the metastores

```sql
select
  metastore_id,
  name,
  cloud,
  global_metastore_id,
  region,
  storage_root,
  storage_root_credential_id,
  storage_root_credential_name
from
  databricks_catalog_metastore;
```

### List metastores that could be shared externally

```sql
select
  metastore_id,
  name,
  cloud,
  global_metastore_id,
  owner,
  account_id
from
  databricks_catalog_metastore
where
  delta_sharing_scope = 'INTERNAL_AND_EXTERNAL';
```

### List metastores that were updated in the last 7 days

```sql
select
  metastore_id,
  name,
  cloud,
  owner,
  updated_at,
  account_id
from
  databricks_catalog_metastore
where
  updated_at >= now() - interval '7 days';
```

### Get effective permissions for each function

```sql
select
  metastore_id,
  name,
  p ->> 'principal' as principal_name,
  p ->> 'privileges' as permissions
from
  databricks_catalog_metastore,
  jsonb_array_elements(metastore_effective_permissions) p;
```

### List metastores with the highest number of effective permissions

```sql
select
  name,
  cloud,
  jsonb_array_length(metastore_effective_permissions) as permission_count
from
  databricks_catalog_metastore
order by
  permission_count desc
limit 10;
```

### Find the most recently updated metastores

```sql
select
  name,
  cloud,
  updated_at
from
  databricks_catalog_metastore
order by
  updated_at desc
limit 10;
```

### Count the number of metastores per cloud

```sql
select
  cloud,
  count(*) as metastore_count
from
  databricks_catalog_metastore
group by
  cloud;
```