---
title: "Steampipe Table: databricks_sharing_share - Query Databricks Sharing Shares using SQL"
description: "Allows users to query Sharing Shares in Databricks, specifically the details of each share, offering insights into shared resources and their permissions."
---

# Table: databricks_sharing_share - Query Databricks Sharing Shares using SQL

Databricks Sharing is a feature in the Databricks platform that enables users to share notebooks, dashboards, and other resources with others. It provides a centralized way to manage the sharing of resources in a secure and controlled manner. Databricks Sharing helps users collaborate effectively by providing access to shared resources based on defined permissions.

## Table Usage Guide

The `databricks_sharing_share` table provides insights into the shared resources within Databricks. As a data engineer or data scientist, you can explore details about each share through this table, including the owner, permissions, and associated metadata. Utilize it to uncover information about shared resources, such as who has access to them and what level of permissions they have.

## Examples

### Basic info
Explore the details of shared resources in your Databricks environment, including who created them and when, to better manage your resources and understand user activity.

```sql+postgres
select
  name,
  comment,
  created_at,
  created_by,
  account_id
from
  databricks_sharing_share;
```

```sql+sqlite
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
Gain insights into the recent sharing activity on your Databricks platform. This query allows you to identify objects that have been shared in the past week, providing transparency and control over data distribution.

```sql+postgres
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

```sql+sqlite
select
  name,
  comment,
  created_at,
  created_by,
  account_id
from
  databricks_sharing_share
where
  created_at > datetime('now', '-7 day');
```

### List all shared objects
Explore which objects have been shared across your account, identifying who added them and when, to gain insights into your shared resources management. This can be particularly useful for auditing or tracking changes over time.

```sql+postgres
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
  jsonb_array_elements(objects) as o;
```

```sql+sqlite
select
  name as share_name,
  json_extract(o.value, '$.name') as object_name,
  json_extract(o.value, '$.added_at') as added_at,
  json_extract(o.value, '$.added_by') as added_by,
  json_extract(o.value, '$.data_object_type') as data_object_type,
  json_extract(o.value, '$.shared_as') as shared_as,
  json_extract(o.value, '$.status') as status,
  account_id
from
  databricks_sharing_share,
  json_each(objects) as o;
```

### Get permissions for each share
Explore which permissions are assigned to each shared resource in Databricks to understand access control and security aspects. This can help in identifying instances where permissions may be overly broad or restrictive, aiding in efficient security management.

```sql+postgres
select
  name,
  p ->> 'principal' as principal_name,
  p ->> 'privileges' as permissions
from
  databricks_sharing_share,
  jsonb_array_elements(permissions) p;
```

```sql+sqlite
select
  name,
  json_extract(p.value, '$.principal') as principal_name,
  json_extract(p.value, '$.privileges') as permissions
from
  databricks_sharing_share,
  json_each(permissions) as p;
```

### List objects shared by a particular owner
Discover the items that a specific user has shared, providing insights into their collaborative behavior and resource distribution. This can be useful for understanding user engagement and tracking resource usage within a team.

```sql+postgres
select
  name,
  comment,
  created_at,
  created_by,
  account_id
from
  databricks_sharing_share
where
  owner = 'owner-username';
```

```sql+sqlite
select
  name,
  comment,
  created_at,
  created_by,
  account_id
from
  databricks_sharing_share
where
  owner = 'owner-username';
```

### Find the account that has the most objects shared
Discover the account with the highest number of shared objects to understand the most active users in terms of resource distribution. This can be useful for identifying key contributors or potential bottlenecks in your data sharing process.

```sql+postgres
select
  account_id,
  count(*) as object_sharing_count
from
  databricks_sharing_share
group by
  account_id
order by
  object_sharing_count desc
limit 1;
```

```sql+sqlite
select
  account_id,
  count(*) as object_sharing_count
from
  databricks_sharing_share
group by
  account_id
order by
  object_sharing_count desc
limit 1;
```