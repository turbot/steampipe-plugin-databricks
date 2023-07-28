# Table: databricks_sharing_recipient

A recipient is an object you create using Recipients/Create to represent an organization which you want to allow access shares.

## Examples

### Basic info

```sql
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

```sql
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

### List allowed ip addresses for each recipient

```sql
select
  name,
  comment,
  ip_address
from
  databricks_sharing_recipient,
  jsonb_array_elements_text(ip_access_list -> 'allowed_ip_addresses') as ip_address;
```

### List sharing reciepients that have databricks accounts

```sql
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

```sql
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

### Get external recipients token details

```sql
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