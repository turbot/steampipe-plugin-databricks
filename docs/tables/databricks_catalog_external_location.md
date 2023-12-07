---
title: "Steampipe Table: databricks_catalog_external_location - Query Databricks Catalog External Locations using SQL"
description: "Allows users to query Catalog External Locations in Databricks, providing information about the external location's name, database name, and owner name."
---

# Table: databricks_catalog_external_location - Query Databricks Catalog External Locations using SQL

Databricks Catalog External Locations is a feature within Databricks that allows you to manage and organize your data. It provides a centralized way to catalog and manage data across various external sources, making data discovery and governance more efficient. Databricks Catalog External Locations helps you stay informed about the data you have and where it is located.

## Table Usage Guide

The `databricks_catalog_external_location` table provides insights into the catalog external locations within Databricks. As a data engineer or data scientist, explore specific details about these locations through this table, including the name of the location, the associated database, and the owner. Utilize it to gain a deeper understanding of your data's organization, distribution, and ownership in Databricks.

## Examples

### Basic info
Explore the metadata of external locations in your Databricks catalog to understand when and by whom they were created, as well as their associated account details. This can help you manage and track the usage of your resources.

```sql+postgres
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

```sql+sqlite
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
Discover the segments that have undergone modification in the external locations during the past week. This can be useful in tracking recent changes and understanding who made them and when.

```sql+postgres
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

```sql+sqlite
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
  updated_at >= datetime('now', '-7 days');
```

### List read only external locations
Discover the segments that are read-only in your external locations. This can be useful to ensure data integrity by preventing accidental changes to these areas.

```sql+postgres
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

```sql+sqlite
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
  read_only = 1;
```

### Get assocated credential for each external location
Determine the associated credentials for each external location in Databricks to understand the correlation between data storage and access permissions. This can aid in managing security and access control across different data repositories.

```sql+postgres
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

```sql+sqlite
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
Explore which permissions are associated with each external location in your Databricks catalog. This can be particularly useful for understanding access rights and maintaining security compliance.

```sql+postgres
select
  name,
  p ->> 'principal' as principal_name,
  p ->> 'privileges' as permissions
from
  databricks_catalog_external_location,
  jsonb_array_elements(external_location_effective_permissions) p;
```

```sql+sqlite
select
  name,
  json_extract(p.value, '$.principal') as principal_name,
  json_extract(p.value, '$.privileges') as permissions
from
  databricks_catalog_external_location,
  json_each(external_location_effective_permissions) as p;
```

### Count the number of external locations per account ID
Discover the segments that have varying numbers of external locations linked to them. This is useful to understand the distribution of resources across different account segments.

```sql+postgres
select
  account_id,
  count(*) AS location_count
from
  databricks_catalog_external_location
group by
  account_id;
```

```sql+sqlite
select
  account_id,
  count(*) AS location_count
from
  databricks_catalog_external_location
group by
  account_id;
```

### List users who created the most external locations
Discover who has been the most active in creating external locations within your Databricks catalog. This is beneficial for understanding usage patterns and identifying key contributors to your data infrastructure.

```sql+postgres
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

```sql+sqlite
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