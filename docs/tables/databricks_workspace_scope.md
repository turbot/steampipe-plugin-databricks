# Table: databricks_workspace_scope

Workspace manages the notebooks and folders in databricks. A notebook is a web-based interface to a document that contains runnable code, visualizations, and explanatory text

## Examples

### Basic info

```sql
select
  name,
  backend_type,
  account_id
from
  databricks_workspace_scope;
```

### List scopes using the azure keyvault

```sql
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

### List total secrets per scope

```sql
select
  name,
  count(*) as total_secrets
from
  databricks_workspace_scope_secret
group by
  name;
```

### List acls for each scope

```sql
select
  name as scope_name,
  backend_type,
  acl ->> 'principal' as principal,
  acl ->> 'permission' as permission
from
  databricks_workspace_scope,
  jsonb_array_elements(acls) as acl;
```