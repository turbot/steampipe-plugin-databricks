---
title: "Steampipe Table: databricks_compute_policy_family - Query Databricks Compute Policy Families using SQL"
description: "Allows users to query Databricks Compute Policy Families, specifically providing details about the ID, name, and member compute policies of each family."
---

# Table: databricks_compute_policy_family - Query Databricks Compute Policy Families using SQL

Databricks Compute Policy Family is a feature within Databricks that groups related compute policies together. It provides a way to manage and organize compute policies, which are rules that control the usage of Databricks compute resources. Databricks Compute Policy Family helps in maintaining a structured and organized approach to resource allocation and usage within Databricks.

## Table Usage Guide

The `databricks_compute_policy_family` table provides insights into Compute Policy Families within Databricks. As a DevOps engineer or a data administrator, explore compute policy family-specific details through this table, including the name, ID, and member compute policies. Utilize it to understand the organization and management of compute resources in your Databricks environment, and to identify any potential resource allocation issues.

## Examples

### Basic info
Explore the various policies within your Databricks compute environment to understand their purpose and details. This can help streamline your cloud operations by identifying any unnecessary or outdated policies.

```sql+postgres
select
  policy_family_id,
  name,
  description,
  definition,
  account_id
from
  databricks_compute_policy_family;
```

```sql+sqlite
select
  policy_family_id,
  name,
  description,
  definition,
  account_id
from
  databricks_compute_policy_family;
```

### List cluster policies associated with policy family
Explore the association between cluster policies and their respective policy families within your account. This could be useful in identifying and understanding the relationships and dependencies between different policies and families, aiding in efficient policy management.

```sql+postgres
select
  f.policy_family_id,
  f.name as policy_family_name,
  p.policy_id,
  p.name as policy_name,
  f.description as policy_family_description,
  f.account_id
from
  databricks_compute_policy_family f,
  databricks_compute_cluster_policy p
where
  f.policy_family_id = p.policy_family_id
  and f.account_id = p.account_id;
```

```sql+sqlite
select
  f.policy_family_id,
  f.name as policy_family_name,
  p.policy_id,
  p.name as policy_name,
  f.description as policy_family_description,
  f.account_id
from
  databricks_compute_policy_family f,
  databricks_compute_cluster_policy p
where
  f.policy_family_id = p.policy_family_id
  and f.account_id = p.account_id;
```

### Find the account with the most policy families
Identify the account with the highest number of policy families. This can help in understanding which account is utilizing the most resources, aiding in resource management and optimization.

```sql+postgres
select
  account_id,
  count(*) as policy_family_count
from
  databricks_compute_policy_family
group by
  account_id
order by
  policy_family_count desc
limit 1;
```

```sql+sqlite
select
  account_id,
  count(*) as policy_family_count
from
  databricks_compute_policy_family
group by
  account_id
order by
  policy_family_count desc
limit 1;
```