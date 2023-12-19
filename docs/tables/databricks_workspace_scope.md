---
title: "Steampipe Table: databricks_workspace_scope - Query Databricks Workspace Scopes using SQL"
description: "Allows users to query Databricks Workspace Scopes, providing insights into the scopes in a Databricks workspace that are available for secrets."
---

# Table: databricks_workspace_scope - Query Databricks Workspace Scopes using SQL

A Databricks Workspace Scope is a feature within Databricks that allows you to manage and organize secrets. Secrets are sensitive data such as database connection strings, tokens, and passwords that are necessary for running notebooks and jobs. Scopes allow for the grouping of secrets to facilitate their management and organization.

## Table Usage Guide

The `databricks_workspace_scope` table provides insights into the Workspace Scopes in Databricks. As a data engineer or security analyst, you can explore the details of these scopes, including their names and other metadata. This table is particularly useful for managing and organizing secrets in a Databricks workspace, ensuring secure and efficient operations.

## Examples

### Basic info
Explore which backend types are associated with specific account IDs to understand the structure of your Databricks workspace. This can aid in managing resources and planning for future workspace configurations.

```sql+postgres
select
  name,
  backend_type,
  account_id
from
  databricks_workspace_scope;
```

```sql+sqlite
select
  name,
  backend_type,
  account_id
from
  databricks_workspace_scope;
```

### List scopes of a desired backend type
Explore which scopes within your Databricks workspace are of a certain backend type. This can help to manage and organize your workspace by identifying areas that utilize specific backend types.

```sql+postgres
select
  name,
  backend_type,
  account_id
from
  databricks_workspace_scope
where
  backend_type = 'DATABRICKS';
```

```sql+sqlite
select
  name,
  backend_type,
  account_id
from
  databricks_workspace_scope
where
  backend_type = 'DATABRICKS';
```

### List scopes using the azure keyvault
Discover the segments that utilize the Azure Keyvault in Databricks workspaces. This can be beneficial in understanding the distribution and usage of Azure Keyvault across different accounts.

```sql+postgres
select
  name,
  keyvault_metadata ->> 'dns_name' as keyvault_dns_name,
  keyvault_metadata ->> 'resource_id' as keyvault_resource_id,
  account_id
from
  databricks_workspace_scope
where
  backend_type = 'AZURE_KEYVAULT';
```

```sql+sqlite
select
  name,
  json_extract(keyvault_metadata, '$.dns_name') as keyvault_dns_name,
  json_extract(keyvault_metadata, '$.resource_id') as keyvault_resource_id,
  account_id
from
  databricks_workspace_scope
where
  backend_type = 'AZURE_KEYVAULT';
```

### List acls for each scope
Explore the access control lists (ACLs) for each scope in your Databricks workspace. This query is useful for identifying who has what permissions, helping to maintain security and access management.

```sql+postgres
select
  name as scope_name,
  backend_type,
  acl ->> 'principal' as principal,
  acl ->> 'permission' as permission
from
  databricks_workspace_scope,
  jsonb_array_elements(acls) as acl;
```

```sql+sqlite
select
  name as scope_name,
  backend_type,
  json_extract(acl.value, '$.principal') as principal,
  json_extract(acl.value, '$.permission') as permission
from
  databricks_workspace_scope,
  json_each(acls) as acl;
```

### Find the account with the maximum workspace scopes
Discover which account has the highest number of workspace scopes, allowing you to identify the most extensively used account in your Databricks setup. This can be beneficial for resource management and understanding usage patterns.

```sql+postgres
select
  account_id,
  count(*) as scope_count
from
  databricks_workspace_scope
group by
  account_id
order by
  scope_count desc;
```

```sql+sqlite
select
  account_id,
  count(*) as scope_count
from
  databricks_workspace_scope
group by
  account_id
order by
  scope_count desc;
```