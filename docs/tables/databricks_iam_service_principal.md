---
title: "Steampipe Table: databricks_iam_service_principal - Query Databricks IAM Service Principals using SQL"
description: "Allows users to query Databricks IAM Service Principals, providing insights into the identities that can be granted access to Azure resources."
---

# Table: databricks_iam_service_principal - Query Databricks IAM Service Principals using SQL

A Databricks IAM Service Principal is an identity that you can create in your Azure AD and assign roles to, allowing it to manage Azure resources. Service Principals are the Azure AD version of a service account, providing access to Azure resources within your organization. They can be used to automate tasks and manage Azure resources at scale.

## Table Usage Guide

The `databricks_iam_service_principal` table provides insights into IAM Service Principals within Azure Databricks. As a DevOps engineer, explore Service Principal-specific details through this table, including roles, permissions, and associated metadata. Utilize it to uncover information about Service Principals, such as those with specific permissions, the roles assigned to them, and the resources they have access to.

## Examples

### Basic info
Explore which service principals are currently active in your Databricks environment. This can help in managing access and security within your applications.

```sql+postgres
select
  id,
  display_name,
  active,
  application_id,
  account_id
from
  databricks_iam_service_principal;
```

```sql+sqlite
select
  id,
  display_name,
  active,
  application_id,
  account_id
from
  databricks_iam_service_principal;
```

### List all inactive service principals
Discover the segments that consist of inactive service principals, allowing you to identify potential areas for optimization or decommissioning. This is particularly useful in managing resources and maintaining security within your Databricks IAM.

```sql+postgres
select
  id,
  display_name,
  active,
  application_id,
  account_id
from
  databricks_iam_service_principal
where
  not active;
```

```sql+sqlite
select
  id,
  display_name,
  active,
  application_id,
  account_id
from
  databricks_iam_service_principal
where
  active = 0;
```

### List assigned roles for each service principal
Explore the roles assigned to each service principal to understand their permissions and responsibilities within the Databricks IAM system. This is useful for auditing purposes and ensuring correct access levels are maintained.

```sql+postgres
select
  u.id,
  u.display_name,
  r ->> 'value' as role,
  r ->> 'type' as type,
  u.account_id
from
  databricks_iam_service_principal u,
  jsonb_array_elements(roles) as r;
```

```sql+sqlite
select
  u.id,
  u.display_name,
  json_extract(r.value, '$.value') as role,
  json_extract(r.value, '$.type') as type,
  u.account_id
from
  databricks_iam_service_principal u,
  json_each(u.roles) as r;
```

### List groups each service principal belongs to
Explore the affiliations of each service principal within a given system to understand their respective group memberships. This is useful for assessing access controls and managing permissions effectively.

```sql+postgres
select
  u.id,
  u.display_name,
  r ->> 'display' as group_display_name,
  r ->> 'value' as role,
  r ->> 'type' as type,
  u.account_id
from
  databricks_iam_service_principal u,
  jsonb_array_elements(groups) as r;
```

```sql+sqlite
select
  u.id,
  u.display_name,
  json_extract(r.value, '$.display') as group_display_name,
  json_extract(r.value, '$.value') as role,
  json_extract(r.value, '$.type') as type,
  u.account_id
from
  databricks_iam_service_principal u,
  json_each(u.groups) as r;
```

### Get service principal with a specific name
Explore which service principal corresponds to a specific user email. This is useful in identifying and managing user access and permissions in a Databricks environment.

```sql+postgres
select
  id,
  display_name,
  active,
  account_id
from
  databricks_iam_service_principal
where
  display_name = 'user@turbot.com';
```

```sql+sqlite
select
  id,
  display_name,
  active,
  account_id
from
  databricks_iam_service_principal
where
  display_name = 'user@turbot.com';
```

### List service principal entitlements
Explore which service principals have specific entitlements in your Databricks IAM setup. This can help you manage access and permissions effectively across different account holders.

```sql+postgres
select
  id,
  display_name,
  r ->> 'value' as entitlement,
  u.account_id
from
  databricks_iam_service_principal u,
  jsonb_array_elements(entitlements) as r;
```

```sql+sqlite
select
  u.id,
  u.display_name,
  json_extract(r.value, '$.value') as entitlement,
  u.account_id
from
  databricks_iam_service_principal u,
  json_each(u.entitlements) as r;
```

### Find the account with the most service principals
Explore which account utilizes the most service principals, aiding in the identification of potential areas of high resource usage or security concerns. This can be beneficial for optimizing resource allocation and enhancing security measures.

```sql+postgres
select
  account_id,
  count(*) as service_principal_count
from
  databricks_iam_service_principal
group by
  account_id
order by
  service_principal_count desc
limit 1;
```

```sql+sqlite
select
  account_id,
  count(*) as service_principal_count
from
  databricks_iam_service_principal
group by
  account_id
order by
  service_principal_count desc
limit 1;
```