---
title: "Steampipe Table: databricks_sharing_recipient - Query Databricks Sharing Recipients using SQL"
description: "Allows users to query Databricks Sharing Recipients, specifically providing insights into the sharing status of Databricks notebooks."
---

# Table: databricks_sharing_recipient - Query Databricks Sharing Recipients using SQL

Databricks Sharing Recipients is a feature within Databricks that provides information about the sharing status of Databricks notebooks. It allows users to monitor and manage the sharing of notebooks across different users and groups within the organization. This feature is crucial for maintaining proper access control and collaborative workflows in Databricks.

## Table Usage Guide

The `databricks_sharing_recipient` table provides insights into the sharing status of Databricks notebooks. As a data engineer or data scientist, explore notebook-specific details through this table, including recipient type, recipient ID, and the notebook path. Utilize it to uncover information about sharing activities, such as who has access to which notebooks, and the types of access they have.

## Examples

### Basic info
Explore which Databricks sharing recipients have been created and by whom. This can be useful to understand who has access to your shared Databricks resources and when they were given access, providing insights into your data security and management.

```sql+postgres
select
  name,
  comment,
  data_recipient_global_metastore_id,
  metastore_id,
  activation_url,
  created_at,
  created_by,
  account_id
from
  databricks_sharing_recipient;
```

```sql+sqlite
select
  name,
  comment,
  data_recipient_global_metastore_id,
  metastore_id,
  activation_url,
  created_at,
  created_by,
  account_id
from
  databricks_sharing_recipient;
```

### List all inactive recipients
Explore which data recipients are inactive on your Databricks account. This can help in identifying unused resources, and potentially optimizing resource allocation.

```sql+postgres
select
  name,
  comment,
  data_recipient_global_metastore_id,
  metastore_id,
  activation_url,
  created_at,
  created_by,
  account_id
from
  databricks_sharing_recipient
where
  not activated;
```

```sql+sqlite
select
  name,
  comment,
  data_recipient_global_metastore_id,
  metastore_id,
  activation_url,
  created_at,
  created_by,
  account_id
from
  databricks_sharing_recipient
where
  activated is not 1;
```

### List allowed ip addresses for each recipient
Explore which recipient names and comments are associated with specific IP addresses. This is useful for understanding and managing access permissions in your Databricks sharing environment.

```sql+postgres
select
  name,
  comment,
  ip_address
from
  databricks_sharing_recipient,
  jsonb_array_elements_text(ip_access_list -> 'allowed_ip_addresses') as ip_address;
```

```sql+sqlite
select
  name,
  comment,
  ip_address.value as ip_address
from
  databricks_sharing_recipient,
  json_each(databricks_sharing_recipient.ip_access_list, '$.allowed_ip_addresses') as ip_address;
```

### List sharing reciepients that have databricks accounts
Explore which sharing recipients have Databricks accounts to manage and control access to shared data effectively. This is useful for maintaining security standards and ensuring appropriate data access.

```sql+postgres
select
  name,
  comment,
  data_recipient_global_metastore_id,
  metastore_id,
  activation_url,
  cloud,
  region,
  sharing_code,
  account_id
from
  databricks_sharing_recipient
where
  not authentication_type = 'DATABRICKS';
```

```sql+sqlite
select
  name,
  comment,
  data_recipient_global_metastore_id,
  metastore_id,
  activation_url,
  cloud,
  region,
  sharing_code,
  account_id
from
  databricks_sharing_recipient
where
  not authentication_type = 'DATABRICKS';
```

### Get permissions for each share
Explore the permissions assigned to each shared resource within your Databricks environment. This can help in identifying any potential security risks, such as overly permissive access rights.

```sql+postgres
select
  name,
  p ->> 'share_name' as share_name,
  pa ->> 'principal' as principal_name,
  pa ->> 'privileges' as privileges
from
  databricks_sharing_recipient,
  jsonb_array_elements(permissions) p,
  jsonb_array_elements(p -> 'privilege_assignments') as pa;
```

```sql+sqlite
select
  name,
  json_extract(p.value, '$.share_name') as share_name,
  json_extract(pa.value, '$.principal') as principal_name,
  json_extract(pa.value, '$.privileges') as privileges
from
  databricks_sharing_recipient,
  json_each(permissions) as p,
  json_each(json_extract(p.value, '$.privilege_assignments')) as pa;
```

### Get external recipients token details
Explore which external recipients have token details in the Databricks sharing recipient list. This can help identify and manage token-based access to Databricks resources, including understanding who created the tokens and when they expire.

```sql+postgres
select
  name,
  comment,
  t ->> 'id' as token_id,
  t ->> 'activation_url' as token_activation_url,
  t ->> 'created_at' as token_created_at,
  t ->> 'created_by' as token_created_by,
  t ->> 'expiration_time' as token_expiration_time,
  account_id
from
  databricks_sharing_recipient,
  jsonb_array_elements(tokens) as t
where
  not authentication_type = 'TOKEN';
```

```sql+sqlite
select
  name,
  comment,
  json_extract(t.value, '$.id') as token_id,
  json_extract(t.value, '$.activation_url') as token_activation_url,
  json_extract(t.value, '$.created_at') as token_created_at,
  json_extract(t.value, '$.created_by') as token_created_by,
  json_extract(t.value, '$.expiration_time') as token_expiration_time,
  account_id
from
  databricks_sharing_recipient,
  json_each(tokens) as t
where
  not authentication_type = 'TOKEN';
```