# Table: databricks_sql_query

Query definitions include the target SQL warehouse, query text, name, description, tags, parameters, and visualizations.

## Examples

### Basic info

```sql
select
  id,
  name,
  created_at,
  description,
  query,
  account_id
from
  databricks_sql_query;
```

### List queries modified in the past 7 days

```sql
select
  id,
  name,
  created_at,
  description,
  last_modified_by,
  updated_at,
  query,
  account_id
from
  databricks_sql_query
where
  updated_at > now() - interval '7' day;
```

### List archived queries

```sql
select
  id,
  name,
  created_at,
  description,
  query,
  account_id
from
  databricks_sql_query
where
  is_archived;
```

### List queries marked as favourite

```sql
select
  id,
  name,
  created_at,
  description,
  query,
  account_id
from
  databricks_sql_query
where
  is_favorite;
```

### List queries that are in draft

```sql
select
  id,
  name,
  created_at,
  description,
  query,
  account_id
from
  databricks_sql_query
where
  is_draft;
```

### List queries that are safe from SQL injection

```sql
select
  id,
  name,
  created_at,
  description,
  query,
  account_id
from
  databricks_sql_query
where
  is_safe;
```

### List queries that can be managed by you

```sql
select
  id,
  name,
  created_at,
  description,
  query,
  account_id
from
  databricks_sql_query
where
  permission_tier = 'CAN_MANAGE';
```

### List parameters associated with each query

```sql
select
  id,
  name,
  created_at,
  description,
  query,
  p ->> 'name' as parameter_name,
  p ->> 'type' as parameter_type,
  p ->> 'value' as parameter_value,
  p ->> 'title' as parameter_title,
  account_id
from
  databricks_sql_query,
  jsonb_array_elements(options -> 'parameters') as p;
```

### List all queries that are not editable

```sql
select
  id,
  name,
  created_at,
  description,
  query,
  account_id
from
  databricks_sql_query
where
  not can_edit;
```

### List visualizations associated to the queries

```sql
select
  id,
  name,
  created_at,
  query,
  account_id,
  visualizations ->> 'CreatedAt' as visualization_create_time,
  visualizations ->> 'Id' as visualization_id,
  visualizations ->> 'Name' as visualization_name,
  visualizations ->> 'Type' as visualization_type,
  visualizations ->> 'UpdatedAt' as visualization_update_time
from
  databricks_sql_query
where
  visualizations is not null;
```