# Table: databricks_catalog_catalog

A catalog is the first layer of Unity Catalog’s three-level namespace. It’s used to organize your data assets.

**Note** Users can see all catalogs on which they have been assigned the USE_CATALOG data permission.

## Examples

### Basic info

```sql
select
  name,
  catalog_type,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog_catalog;
```

### List catalogs modified in the last 7 days

```sql
select
  name,
  catalog_type,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog_catalog
where
  updated_at >= now() - interval '7 days';
```

### List catalogs with auto maintenance enabled

```sql
select
  name,
  catalog_type,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog_catalog
where
  enable_auto_maintenance = 'ENABLE';
```

### List catalogs that are not isolated to the workspace

```sql
select
  name,
  catalog_type,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog_catalog
where
  isolation_mode = 'OPEN';
```

### Get effective permissions for each catalog

```sql
select
  name,
  p ->> 'principal' as principal_name,
  p ->> 'privileges' as permissions
from
  databricks_catalog_catalog,
  jsonb_array_elements(catalog_effective_permissions) p;
```

### Get total catalogs of each type

```sql
select
  catalog_type,
  count(*) as total_catalogs
from
  databricks_catalog_catalog
group by
  catalog_type;
```