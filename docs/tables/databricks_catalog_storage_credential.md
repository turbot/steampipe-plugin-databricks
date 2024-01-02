---
title: "Steampipe Table: databricks_catalog_storage_credential - Query Databricks Catalog Storage Credentials using SQL"
description: "Allows users to query Catalog Storage Credentials in Databricks, providing detailed information about the storage credentials used in the Databricks data catalog."
---

# Table: databricks_catalog_storage_credential - Query Databricks Catalog Storage Credentials using SQL

A Databricks Catalog Storage Credential is a resource within Databricks that is used to authenticate and authorize access to data stored in external storage. It is a key component in securing data access and ensuring only authorized users can access specific data. These credentials are used in conjunction with Databricks data catalog, which provides a unified view of all data assets.

## Table Usage Guide

The `databricks_catalog_storage_credential` table provides insights into Catalog Storage Credentials within Databricks. As a data engineer or a security analyst, you can explore credential-specific details through this table, including the type of credential, the scope of the credential, and other associated metadata. Utilize it to uncover information about the credentials, such as their access level, the data they can access, and their usage in securing data access.

## Examples

### Basic info
Explore the creation details of your storage credentials in Databricks. This can help you understand who created what and when, providing insights into your resource management and usage patterns.

```sql+postgres
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

```sql+sqlite
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
Discover the segments that have undergone modifications in their credentials within the past week. This is particularly useful for maintaining security and keeping track of recent changes in your system.

```sql+postgres
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

```sql+sqlite
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
  updated_at >= datetime('now', '-7 days');
```

### List credentials that are read only
Discover the segments that have read-only access to your data, enabling you to identify potential security risks and ensure proper data management. This is particularly useful in maintaining data integrity and preventing unauthorized modifications.

```sql+postgres
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

```sql+sqlite
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
  read_only = 1;
```

### List credentials that are the current metastore's root storage credential
Determine the areas in which credentials are being used for managed storage within the Databricks catalog. This allows for an understanding of the relationships between different elements, such as accounts and metastores, and can aid in optimizing data management strategies.

```sql+postgres
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

```sql+sqlite
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
  used_for_managed_storage = 1;
```

### List storage credentials for AWS
Explore which AWS storage credentials are linked to your Databricks catalog. This can help you manage access and permissions, ensuring the right roles are assigned to the right resources.

```sql+postgres
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

```sql+sqlite
select
  id,
  name,
  metastore_id,
  json_extract(aws_iam_role, '$.external_id') as external_id,
  json_extract(aws_iam_role, '$.role_arn') as aws_iam_role_arn,
  json_extract(aws_iam_role, '$.unity_catalog_iam_arn') as aws_unity_catalog_iam_arn,
  account_id
from
  databricks_catalog_storage_credential;
```

### List storage credentials for Azure
This example helps you analyze the stored access credentials in Azure. It's useful for auditing purposes, allowing you to keep track of which credentials are being used in your Azure environment.

```sql+postgres
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

```sql+sqlite
select
  id,
  name,
  metastore_id,
  json_extract(azure_managed_identity, '$.access_connector_id') as azure_access_connector_id,
  json_extract(azure_managed_identity, '$.credential_id') as credential_id,
  json_extract(azure_managed_identity, '$.managed_identity_id') as azure_managed_identity_id,
  json_extract(azure_service_principal, '$.application_id') as azure_service_principal_application_id,
  json_extract(azure_service_principal, '$.client_secret') as azure_service_principal_client_secret,
  json_extract(azure_service_principal, '$.directory_id') as azure_service_principal_directory_id,
  account_id
from
  databricks_catalog_storage_credential;
```

### List storage credentials for GCP
Discover the segments that contain storage credentials for your Google Cloud Platform (GCP) account. This can be useful to manage and monitor your storage credentials, ensuring they are correctly configured and secure.

```sql+postgres
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

```sql+sqlite
select
  id,
  name,
  metastore_id,
  json_extract(databricks_gcp_service_account, '$.credential_id') as credential_id,
  json_extract(databricks_gcp_service_account, '$.email') as gcp_service_account_email,
  account_id
from
  databricks_catalog_storage_credential;
```

### Get effective permissions for each credential
Explore the specific permissions associated with each credential to gain insights into your data access and security settings. This is useful for maintaining robust security practices and understanding who has access to what data.

```sql+postgres
select
  name,
  p ->> 'principal' as principal_name,
  p ->> 'privileges' as permissions
from
  databricks_catalog_storage_credential,
  jsonb_array_elements(storage_credential_effective_permissions) p;
```

```sql+sqlite
select
  name,
  json_extract(p.value, '$.principal') as principal_name,
  json_extract(p.value, '$.privileges') as permissions
from
  databricks_catalog_storage_credential,
  json_each(storage_credential_effective_permissions) as p;
```

### Find the account with the most entries
Determine the account with the highest number of entries to understand usage patterns and identify potential areas of optimization or investigation.

```sql+postgres
select
  account_id,
  count(*) as entry_count
from
  databricks_catalog_storage_credential
group by
  account_id
order by
  entry_count desc
limit 1;
```

```sql+sqlite
select
  account_id,
  count(*) as entry_count
from
  databricks_catalog_storage_credential
group by
  account_id
order by
  entry_count desc
limit 1;
```