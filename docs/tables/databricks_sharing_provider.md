---
title: "Steampipe Table: databricks_sharing_provider - Query Databricks Sharing Providers using SQL"
description: "Allows users to query Databricks Sharing Providers, specifically the sharing provider configurations, providing insights into the sharing settings and potential configurations."
---

# Table: databricks_sharing_provider - Query Databricks Sharing Providers using SQL

Databricks Sharing Providers is a resource within Databricks that allows you to manage and configure the sharing settings for your workspace. It provides a centralized way to set up and manage sharing providers, including email, Slack, and more. Databricks Sharing Providers helps you stay informed about the sharing configurations and take appropriate actions when predefined conditions are met.

## Table Usage Guide

The `databricks_sharing_provider` table provides insights into the sharing providers within Databricks. As a DevOps engineer, explore sharing provider-specific details through this table, including configurations and associated metadata. Utilize it to uncover information about sharing providers, such as those with specific configurations, the relationships between sharing providers, and the verification of sharing settings.

## Examples

### Basic info
Explore the basic information from your Databricks sharing provider to understand who created specific data and when. This can help you track data provenance and maintain accountability within your team.

```sql+postgres
select
  name,
  comment,
  data_provider_global_metastore_id,
  metastore_id
  created_at,
  created_by,
  account_id
from
  databricks_sharing_provider;
```

```sql+sqlite
select
  name,
  comment,
  data_provider_global_metastore_id,
  metastore_id,
  created_at,
  created_by,
  account_id
from
  databricks_sharing_provider;
```

### List providers created in the last 7 days
Explore which data sharing providers were created in the past week. This can be beneficial to maintain an up-to-date overview of recent changes in your Databricks environment.

```sql+postgres
select
  name,
  comment,
  data_provider_global_metastore_id,
  metastore_id
  created_at,
  created_by,
  account_id
from
  databricks_sharing_provider
where
  created_at >= now() - interval '7' day;
```

```sql+sqlite
select
  name,
  comment,
  data_provider_global_metastore_id,
  metastore_id,
  created_at,
  created_by,
  account_id
from
  databricks_sharing_provider
where
  created_at >= datetime('now', '-7 day');
```

### List providers authenticated by Databricks
Explore which providers are authenticated by Databricks to understand the security and access configuration of your data sharing setup. This is useful in auditing your data sharing environment and ensuring only authorized providers have access.

```sql+postgres
select
  name,
  comment,
  data_provider_global_metastore_id,
  metastore_id
  created_at,
  created_by,
  account_id
from
  databricks_sharing_provider
where
  authentication_type = 'DATABRICKS';
```

```sql+sqlite
select
  name,
  comment,
  data_provider_global_metastore_id,
  metastore_id,
  created_at,
  created_by,
  account_id
from
  databricks_sharing_provider
where
  authentication_type = 'DATABRICKS';
```

### List all shares by each provider
Explore the distribution of shares among different providers to better manage and optimize resource allocation. This can assist in identifying which providers have the most shares, aiding in strategic decision-making.

```sql+postgres
select
  name as provider_name,
  s ->> 'name' as provider_share_name,
  account_id
from
  databricks_sharing_provider,
  jsonb_array_elements(shares) as s;
```

```sql+sqlite
select
  name as provider_name,
  json_extract(s.value, '$.name') as provider_share_name,
  account_id
from
  databricks_sharing_provider,
  json_each(shares) as s;
```

### Get recipient profile for each shared provider
Explore which shared providers have been configured with token-based authentication. This can help in understanding the distribution and usage of different authentication types across your shared providers.

```sql+postgres
select
  name as share_name,
  recipient_profile ->> 'bearer_token' as bearer_token,
  recipient_profile ->> 'endpoint' as endpoint,
  recipient_profile ->> 'share_credentials_version' as share_credentials_version,
  account_id
from
  databricks_sharing_provider
where
  authentication_type = 'TOKEN';
```

```sql+sqlite
select
  name as share_name,
  json_extract(recipient_profile, '$.bearer_token') as bearer_token,
  json_extract(recipient_profile, '$.endpoint') as endpoint,
  json_extract(recipient_profile, '$.share_credentials_version') as share_credentials_version,
  account_id
from
  databricks_sharing_provider
where
  authentication_type = 'TOKEN';
```

### List the owner in order of the number of providers
Discover the segments that have the highest number of providers, organized by the owner. This can help assess the distribution and management of providers across different owners.

```sql+postgres
select
  owner,
  count(*) as provider_count
from
  databricks_sharing_provider
group by
  owner
order by
  provider_count desc;
```

```sql+sqlite
select
  owner,
  count(*) as provider_count
from
  databricks_sharing_provider
group by
  owner
order by
  provider_count desc;
```