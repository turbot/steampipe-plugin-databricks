# Table: databricks_files_dbfs

DBFS makes it simple to interact with various data sources without having to include a users credentials every time to read a file.

## Examples

### Basic info

```sql
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

```sql
select
  path,
  modification_time
from
  databricks_files_dbfs
where
  path_prefix = '/'
  and is_dir;
```

### List all the files in DBFS

```sql
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

```sql
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

```sql
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

### Get contents of a particular file/directory

```sql
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