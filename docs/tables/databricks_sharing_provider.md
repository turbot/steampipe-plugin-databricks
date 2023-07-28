# Table: databricks_sharing_provider

A share is a container instantiated with Shares/Create. Once created you can iteratively register a collection of existing data assets defined within the metastore using Shares/Update. A data provider is an object representing the organization in the real world who shares the data

## Examples

### Basic info

```sql
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

### List providers authenticated by Databricks

```sql
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

### List all shares by each provider

```sql
select
  name as provider_name,
  s ->> 'name' as provider_share_name,
  account_id
from
  databricks_sharing_provider,
  jsonb_array_elements(shares) as s;
```

### Get recipient profile for each shared provider

```sql
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