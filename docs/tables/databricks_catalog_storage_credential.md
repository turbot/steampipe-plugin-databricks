# Table: databricks_catalog_storage_credential

A storage credential represents an authentication and authorization mechanism for accessing data stored on your cloud tenant. Each storage credential is subject to Unity Catalog access-control policies that control which users and groups can access the credential.

## Examples

### Basic info

```sql
select
  id,
  name,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from
  databricks_catalog_storage_credential;
```

### List credentials modified in the last 7 days

```sql
select
  id,
  name,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from  
  databricks_catalog_storage_credential
where
  updated_at >= now() - interval '7 days';
```

### List credentials that are read only

```sql
select
  id,
  name,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from  
  databricks_catalog_storage_credential
where
  read_only;
```

### List credentials that are the current metastore's root storage credential

```sql
select
  id,
  name,
  comment,
  created_at,
  created_by,
  metastore_id,
  account_id
from  
  databricks_catalog_storage_credential
where
  used_for_managed_storage;
```

### List storage credentials for AWS

```sql
select
  id,
  name,
  metastore_id,
  aws_iam_role ->> 'external_id' as external_id,
  aws_iam_role ->> 'role_arn' aws_iam_role_arn,
  aws_iam_role ->> 'unity_catalog_iam_arn' as aws_unity_catalog_iam_arn,
  account_id
from  
  databricks_catalog_storage_credential;
```

### List storage credentials for Azure

```sql
select
  id,
  name,
  metastore_id,
  azure_managed_identity ->> 'access_connector_id' as azure_access_connector_id,
  azure_managed_identity ->> 'credential_id' as credential_id,
  azure_managed_identity ->> 'managed_identity_id' as azure_managed_identity_id,
  azure_service_principal ->> 'application_id' as azure_service_principal_application_id,
  azure_service_principal ->> 'client_secret' as azure_service_principal_client_secret,
  azure_service_principal ->> 'directory_id' as azure_service_principal_directory_id,
  account_id
from  
  databricks_catalog_storage_credential;
```

### List storage credentials for GCP

```sql
select
  id,
  name,
  metastore_id,
  databricks_gcp_service_account ->> 'credential_id' as credential_id,
  databricks_gcp_service_account ->> 'email' as gcp_service_account_email,
  account_id
from  
  databricks_catalog_storage_credential;
```

### Get effective permissions for each credential

```sql
select
  name,
  p ->> 'principal' as principal_name,
  p ->> 'privileges' as permissions
from
  databricks_catalog_storage_credential,
  jsonb_array_elements(storage_credential_effective_permissions) p;
```
