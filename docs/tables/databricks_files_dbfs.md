---
title: "Steampipe Table: databricks_files_dbfs - Query Databricks DBFS Files using SQL"
description: "Allows users to query DBFS Files in Databricks, specifically providing information about file paths, file sizes, and file types."
---

# Table: databricks_files_dbfs - Query Databricks DBFS Files using SQL

Databricks DBFS (Databricks File System) is a distributed file system installed on Databricks clusters. It allows users to interact with object storage via standard file system operations, and is crucial for storing all types of data such as ETL outputs, machine learning models, etc. DBFS provides the interface to access and manage data across all Databricks workspaces and to persist objects across cluster lifetimes.

## Table Usage Guide

The `databricks_files_dbfs` table provides insights into DBFS Files within Databricks. As a data scientist or data engineer, explore file-specific details through this table, including file paths, sizes, and types. Utilize it to manage and organize your data in Databricks, ensuring efficient data processing and analytics.

## Examples

### Basic info
Explore the basic information of files stored in Databricks, including file size, modification time, and content. This can be particularly beneficial for understanding the file structure, tracking changes, and managing storage effectively.

```sql+postgres
select
  path,
  file_size,
  is_dir,
  modification_time,
  content
from
  databricks_files_dbfs
where
  path_prefix = '/';
```

```sql+sqlite
select
  path,
  file_size,
  is_dir,
  modification_time,
  content
from
  databricks_files_dbfs
where
  path_prefix = '/';
```

### List all the directories in DBFS
Explore all directories in DBFS to gain insights into their modification times, which can be useful for understanding file system changes and data modifications.

```sql+postgres
select
  path,
  modification_time
from
  databricks_files_dbfs
where
  path_prefix = '/'
  and is_dir;
```

```sql+sqlite
select
  path,
  modification_time
from
  databricks_files_dbfs
where
  path_prefix = '/'
  and is_dir = 1;
```

### List all the files in DBFS
Explore which files are stored in your DBFS by assessing their paths, sizes, and modification times. This could be useful in instances where you need to manage your storage space or track changes to files over time.

```sql+postgres
select
  path,
  file_size,
  modification_time
from
  databricks_files_dbfs
where
  path_prefix = '/'
  and not is_dir;
```

```sql+sqlite
select
  path,
  file_size,
  modification_time
from
  databricks_files_dbfs
where
  path_prefix = '/'
  and not is_dir;
```

### List all the files in DBFS that are larger than 1MB
Explore which files in your Databricks File System (DBFS) are larger than 1MB. This can be useful for managing your storage and identifying files that might be taking up more space than necessary.

```sql+postgres
select
  path,
  file_size,
  modification_time
from
  databricks_files_dbfs
where
  path_prefix = '/'
  and not is_dir
  and file_size > 1000000;
```

```sql+sqlite
select
  path,
  file_size,
  modification_time
from
  databricks_files_dbfs
where
  path_prefix = '/'
  and not is_dir
  and file_size > 1000000;
```

### List all the files in DBFS that were modified in the past 7 days
Discover the segments that have seen recent changes by pinpointing the specific locations where files have been modified in the past week. This allows you to keep track of updates and changes, ensuring you're always working with the most recent data.

```sql+postgres
select
  path,
  file_size,
  is_dir,
  modification_time
from
  databricks_files_dbfs
where
  path_prefix = '/'
  and modification_time > now() - interval '7' day;
```

```sql+sqlite
select
  path,
  file_size,
  is_dir,
  modification_time
from
  databricks_files_dbfs
where
  path_prefix = '/'
  and modification_time > datetime('now', '-7 day');
```

### Get contents of a particular file/directory
Explore the contents of a specific file or directory to understand its size, modification time, and data. This can be useful for auditing file changes, monitoring data usage, or troubleshooting issues related to file content.

```sql+postgres
select
  path,
  file_size,
  modification_time,
  content ->> 'bytes_read' as bytes_read,
  content ->> 'data' as data
from
  databricks_files_dbfs
where
  path = '/path/to/file/directory';
```

```sql+sqlite
select
  path,
  file_size,
  modification_time,
  json_extract(content, '$.bytes_read') as bytes_read,
  json_extract(content, '$.data') as data
from
  databricks_files_dbfs
where
  path = '/path/to/file/directory';
```