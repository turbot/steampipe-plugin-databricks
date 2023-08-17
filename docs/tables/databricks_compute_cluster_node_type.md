# Table: databricks_compute_cluster_node_type

Spark node types can be used to launch a cluster.

## Examples

### Basic info

```sql
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

```sql
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

```sql
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

### List node types with I/O caching enabled

```sql
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

### List node types that support port forwarding

```sql
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

### Get node instance type details for each node type

```sql
select
  node_type_id,
  node_instance_type ->> 'instance_type_id' as instance_type_id,
  node_instance_type ->> 'local_disk_size_gb' as local_disk_size_gb,
  node_instance_type ->> 'local_disks' as local_disks,
  account_id
from
  databricks_compute_cluster_node_type;
```

### List hidden node types

```sql
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

### List gravition node types

```sql
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

### List all non-deprecated node types

```sql
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

### List node types having more than one GPUs

```sql
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

### List node types that support EBS volumes

```sql
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

### List node types in order of available memory

```sql
select
  node_type_id,
  category,
  memory_mb
from
  databricks_compute_cluster_node_type
order by
  memory_mb desc;
```