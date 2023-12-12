---
title: "Steampipe Table: databricks_compute_cluster_node_type - Query Databricks Compute Cluster Node Types using SQL"
description: "Allows users to query Databricks Compute Cluster Node Types, specifically the attributes, limitations, and capabilities of each node type."
---

# Table: databricks_compute_cluster_node_type - Query Databricks Compute Cluster Node Types using SQL

Databricks Compute Cluster Node Types are the units of processing power and memory that Databricks uses to run computations. Each node type has specific attributes, limitations, and capabilities. The node types are designed to optimize the performance of Databricks workloads.

## Table Usage Guide

The `databricks_compute_cluster_node_type` table provides insights into the node types available in Databricks Compute Clusters. As a Data Engineer, explore node type-specific details through this table, including memory, CPU, and storage attributes, as well as any limitations or special capabilities. Utilize it to select the most suitable node type for your specific Databricks workloads, ensuring optimal performance and cost-effectiveness.

## Examples

### Basic info
Explore the different categories of compute cluster node types in Databricks, understanding their memory and core capacities. This information can help optimize resource allocation and performance across different accounts.

```sql+postgres
select
  node_type_id,
  category,
  description,
  memory_mb,
  num_cores,
  account_id
from
  databricks_compute_cluster_node_type;
```

```sql+sqlite
select
  node_type_id,
  category,
  description,
  memory_mb,
  num_cores,
  account_id
from
  databricks_compute_cluster_node_type;
```

### List total node types per category
Explore the distribution of node types across different categories within your account. This allows you to understand your usage patterns and potentially optimize resource allocation.

```sql+postgres
select
  category,
  count(*) as num_node_types,
  account_id
from
  databricks_compute_cluster_node_type
group by
  category,
  account_id;
```

```sql+sqlite
select
  category,
  count(*) as num_node_types,
  account_id
from
  databricks_compute_cluster_node_type
group by
  category,
  account_id;
```

### List node types encrypted in transit
Explore which node types within your Databricks compute cluster are encrypted in transit. This can be useful for ensuring security compliance across your data processing infrastructure.

```sql+postgres
select
  node_type_id,
  category,
  description,
  memory_mb
  num_cores,
  account_id
from
  databricks_compute_cluster_node_type
where
  is_encrypted_in_transit;
```

```sql+sqlite
select
  node_type_id,
  category,
  description,
  memory_mb,
  num_cores,
  account_id
from
  databricks_compute_cluster_node_type
where
  is_encrypted_in_transit = 1;
```

### List node types with I/O caching enabled
Explore which node types have I/O caching enabled in your Databricks compute cluster. This can help you determine which nodes could potentially offer enhanced performance due to caching, aiding in efficient resource allocation.

```sql+postgres
select
  node_type_id,
  category,
  description,
  memory_mb
  num_cores,
  account_id
from
  databricks_compute_cluster_node_type
where
  is_io_cache_enabled;
```

```sql+sqlite
select
  node_type_id,
  category,
  description,
  memory_mb,
  num_cores,
  account_id
from
  databricks_compute_cluster_node_type
where
  is_io_cache_enabled;
```

### List node types that support port forwarding
Discover the types of nodes that support port forwarding to understand their characteristics and capabilities. This can help optimize your network configuration by choosing nodes that best meet your port forwarding needs.

```sql+postgres
select
  node_type_id,
  category,
  description,
  memory_mb
  num_cores,
  account_id
from
  databricks_compute_cluster_node_type
where
  support_port_forwarding;
```

```sql+sqlite
select
  node_type_id,
  category,
  description,
  memory_mb,
  num_cores,
  account_id
from
  databricks_compute_cluster_node_type
where
  support_port_forwarding = 1;
```

### Get node instance type details for each node type
Discover the details of each node type in your system by understanding the specifics of the instance type, such as disk size. This could be useful to manage resources and plan for infrastructure upgrades.

```sql+postgres
select
  node_type_id,
  node_instance_type ->> 'instance_type_id' as instance_type_id,
  node_instance_type ->> 'local_disk_size_gb' as local_disk_size_gb,
  node_instance_type ->> 'local_disks' as local_disks,
  account_id
from
  databricks_compute_cluster_node_type;
```

```sql+sqlite
select
  node_type_id,
  json_extract(node_instance_type, '$.instance_type_id') as instance_type_id,
  json_extract(node_instance_type, '$.local_disk_size_gb') as local_disk_size_gb,
  json_extract(node_instance_type, '$.local_disks') as local_disks,
  account_id
from
  databricks_compute_cluster_node_type;
```

### List hidden node types
Discover the hidden node types within your Databricks compute cluster to better manage resources and understand the configuration of your system. This helps in optimizing the usage of memory and cores, thereby improving overall system efficiency.

```sql+postgres
select
  node_type_id,
  category,
  description,
  memory_mb
  num_cores,
  account_id
from
  databricks_compute_cluster_node_type
where
  is_hidden;
```

```sql+sqlite
select
  node_type_id,
  category,
  description,
  memory_mb,
  num_cores,
  account_id
from
  databricks_compute_cluster_node_type
where
  is_hidden = 1;
```

### List gravition node types
Analyze the configuration of your Databricks compute cluster to identify instances where Graviton node types are being used. This could be useful for assessing the efficiency and performance of your data processing tasks.

```sql+postgres
select
  node_type_id,
  category,
  description,
  memory_mb
  num_cores,
  account_id
from
  databricks_compute_cluster_node_type
where
  is_graviton;
```

```sql+sqlite
select
  node_type_id,
  category,
  description,
  memory_mb,
  num_cores,
  account_id
from
  databricks_compute_cluster_node_type
where
  is_graviton = 1;
```

### List all non-deprecated node types
Explore the different types of non-deprecated nodes in your Databricks compute cluster to understand their categories, descriptions, and hardware specifications such as memory and core count. This can be useful for optimizing resource allocation and identifying suitable node types for your specific workload requirements.

```sql+postgres
select
  node_type_id,
  category,
  description,
  memory_mb
  num_cores,
  account_id
from
  databricks_compute_cluster_node_type
where
  not is_deprecated;
```

```sql+sqlite
select
  node_type_id,
  category,
  description,
  memory_mb,
  num_cores,
  account_id
from
  databricks_compute_cluster_node_type
where
  is_deprecated = 0;
```

### List node types having more than one GPUs
Discover the segments that have more than one GPU within your Databricks compute cluster node types. This can be useful in identifying high-performance node types, thus aiding in resource allocation and optimization.

```sql+postgres
select
  node_type_id,
  category,
  description,
  memory_mb
  num_cores,
  account_id
from
  databricks_compute_cluster_node_type
where
  num_gpus > 1;
```

```sql+sqlite
select
  node_type_id,
  category,
  description,
  memory_mb,
  num_cores,
  account_id
from
  databricks_compute_cluster_node_type
where
  num_gpus > 1;
```

### List node types that support EBS volumes
Discover the types of nodes that are compatible with EBS volumes in a Databricks computing environment. This can be useful when planning resource allocation or designing data processing tasks.

```sql+postgres
select
  node_type_id,
  category,
  description,
  memory_mb
  num_cores,
  account_id
from
  databricks_compute_cluster_node_type
where
  support_ebs_volumes;
```

```sql+sqlite
select
  node_type_id,
  category,
  description,
  memory_mb,
  num_cores,
  account_id
from
  databricks_compute_cluster_node_type
where
  support_ebs_volumes = 1;
```

### List node types in order of available memory
Explore the types of nodes in your Databricks compute cluster, ordered by the amount of available memory. This can help prioritize resource allocation and optimize cluster performance.

```sql+postgres
select
  node_type_id,
  category,
  memory_mb
from
  databricks_compute_cluster_node_type
order by
  memory_mb desc;
```

```sql+sqlite
select
  node_type_id,
  category,
  memory_mb
from
  databricks_compute_cluster_node_type
order by
  memory_mb desc;
```