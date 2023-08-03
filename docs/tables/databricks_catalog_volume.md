# Table: databricks_catalog_volume

Volumes are a Unity Catalog (UC) capability for accessing, storing, governing, organizing and processing files

The `databricks_catalog_volume` table can be used to query information about any volume, and **you must specify the catalog name and schema name** in the where or join clause using the `catalog_name` and `schema_name` columns.

**Note** To query a volume, the user must have the **USE_CATALOG** privilege on the catalog and the **USE_SCHEMA** privilege on the schema, and the output list contains only volumes for which either the user has the **EXECUTE** privilege or the user is the owner

## Examples

### Basic info

```sql
select
  volume_id,
  name,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog_volume
where
  catalog_name = 'catalog'
  and schema_name = 'schema';
```

### List volumes modified in the last 7 days

```sql
select
  volume_id,
  name,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog_volume
where
  updated_at >= now() - interval '7 days'
  and catalog_name = 'catalog'
  and schema_name = 'schema';
```

### List all external volumes

```sql
select
  volume_id,
  name,
  volume_type,
  storage_location,
  account_id
from
  databricks_catalog_volume
where
  volume_type = 'EXTERNAL'
  and catalog_name = 'catalog'
  and schema_name = 'schema';
```

### Get details for a specific volume

```sql
select
  volume_id,
  name,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog_volume
where
  full_name = '__catalog_name__.__schema_name__.__volume_name__';
```