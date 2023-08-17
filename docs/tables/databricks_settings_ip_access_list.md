# Table: databricks_settings_ip_access_list

IP Access List enables admins to configure IP access lists. IP access lists affect web application access and REST API access to this workspace only.

## Examples

### Basic info

```sql
select
  list_id,
  label,
  address_count,
  created_at,
  created_by,
  account_id
from
  databricks_settings_ip_access_list;
```

### List access lists modified in the last 7 days

```sql
select
  list_id,
  label,
  address_count,
  updated_at,
  updated_by,
  enabled,
  account_id
from
  databricks_settings_ip_access_list
where
  updated_at > now() - interval '7' day;
```

### List all access lists which are disabled

```sql
select
  list_id,
  label,
  address_count,
  created_at,
  created_by,
  account_id
from
  databricks_settings_ip_access_list
where
  not enabled;
```

### List all the addresses in each access list

```sql
select
  list_id,
  label,
  address,
  account_id
from
  databricks_settings_ip_access_list,
  jsonb_array_elements(ip_addresses) as address
where
  enabled;
```

### Get access lists that allow all the requests

```sql
select
  list_id,
  label,
  address,
  created_by,
  updated_by,
  account_id
from
  databricks_settings_ip_access_list,
  jsonb_array_elements_text(ip_addresses) as address
where
  (enabled
  and address = '0.0.0.0/0'
  and list_type = 'ALLOW')
  or (not enabled);
```