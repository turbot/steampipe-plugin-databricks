---
title: "Steampipe Table: databricks_settings_ip_access_list - Query Databricks IP Access Lists using SQL"
description: "Allows users to query IP Access Lists in Databricks, providing detailed information on the IP addresses and their access permissions."
---

# Table: databricks_settings_ip_access_list - Query Databricks IP Access Lists using SQL

Databricks IP Access Lists is a feature within Databricks that allows you to control which IP addresses have access to your Databricks workspace. It provides a mechanism to ensure that only trusted IP addresses can access your Databricks resources, enhancing the security of your data and applications. This feature is crucial for managing access permissions and maintaining the integrity of your Databricks workspace.

## Table Usage Guide

The `databricks_settings_ip_access_list` table provides insights into the IP Access Lists within Databricks. As a security analyst or a DevOps engineer, you can explore detailed information about the IP addresses and their access permissions through this table. You can use it to audit access permissions, identify trusted IP addresses, and ensure that only authorized IP addresses have access to your Databricks resources.

## Examples

### Basic info
Explore the creation and composition of IP access lists within your Databricks settings. This allows you to understand who created each list, when it was created, and how many addresses it contains, which can be useful for auditing and security purposes.

```sql+postgres
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

```sql+sqlite
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
Discover the segments that have seen modifications in their access lists within the past week. This can be beneficial in monitoring recent changes to enhance security and control access within your digital environment.

```sql+postgres
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

```sql+sqlite
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
  updated_at > datetime('now', '-7 day');
```

### List all access lists which are disabled
Discover the segments that have disabled access lists in your Databricks settings. This is useful for identifying potential security risks or areas where access has been restricted.

```sql+postgres
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

```sql+sqlite
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
  enabled = 0;
```

### List all the addresses in each access list
Explore which IP addresses are included in each access list in your Databricks settings, useful for maintaining network security and controlling access to your data.

```sql+postgres
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

```sql+sqlite
select
  list_id,
  label,
  address.value as address,
  account_id
from
  databricks_settings_ip_access_list,
  json_each(ip_addresses) as address
where
  enabled;
```

### Get access lists that allow all the requests
Explore which access lists are allowing all requests, including those that are currently disabled. This is useful for identifying potential security vulnerabilities in your Databricks settings.

```sql+postgres
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

```sql+sqlite
select
  list_id,
  label,
  address.value as address,
  created_by,
  updated_by,
  account_id
from
  databricks_settings_ip_access_list,
  json_each(ip_addresses) as address
where
  (enabled
  and address.value = '0.0.0.0/0'
  and list_type = 'ALLOW')
  or (not enabled);
```