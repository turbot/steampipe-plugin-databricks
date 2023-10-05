## v0.2.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#19](https://github.com/turbot/steampipe-plugin-databricks/pull/19))

## v0.2.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#16](https://github.com/turbot/steampipe-plugin-databricks/pull/16))
- Recompiled plugin with Go version `1.21`. ([#16](https://github.com/turbot/steampipe-plugin-databricks/pull/16))

## v0.1.0 [2023-08-28]

_Breaking changes_

- The following configuration arguments have been renamed for better consistency: ([#4](https://github.com/turbot/steampipe-plugin-databricks/pull/4))
  - `config_file` to `config_file_path`
  - `config_profile` to `profile`
- The following tables have been renamed to remove redundant wording: ([#4](https://github.com/turbot/steampipe-plugin-databricks/pull/4))
  - `databricks_catalog_catalog` -> `databricks_catalog`
  - `databricks_jobs_job` -> `databricks_job`
  - `databricks_jobs_job_run` -> `databricks_job_run`
  - `databricks_pipelines_pipeline` -> `databricks_pipeline`
  - `databricks_pipelines_pipeline_event` -> `databricks_pipeline_event`
  - `databricks_pipelines_pipeline_update` -> `databricks_pipeline_update`
  - `databricks_workspace_workspace` -> `databricks_workspace`

_Bug fixes_

- Fixed the plugin to correctly return an empty row instead of an error in cases where the tables lack resources available for querying. ([#4](https://github.com/turbot/steampipe-plugin-databricks/pull/4))
- Fixed the configuration argument validation by eliminating the necessity of setting account_host, account_token/username/password, and workspace host configuration arguments/environment variables in certain scenarios. ([#4](https://github.com/turbot/steampipe-plugin-databricks/pull/4))

## v0.0.1 [2023-08-17]

_What's new?_

- New tables added
  - [databricks_catalog_catalog](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_catalog_catalog)
  - [databricks_catalog_connection](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_catalog_connection)
  - [databricks_catalog_external_location](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_catalog_external_location)
  - [databricks_catalog_function](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_catalog_function)
  - [databricks_catalog_metastore](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_catalog_metastore)
  - [databricks_catalog_schema](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_catalog_schema)
  - [databricks_catalog_storage_credential](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_catalog_storage_credential)
  - [databricks_catalog_system_schema](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_catalog_system_schema)
  - [databricks_catalog_table](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_catalog_table)
  - [databricks_catalog_volume](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_catalog_volume)
  - [databricks_compute_cluster](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_compute_cluster)
  - [databricks_compute_cluster_node_type](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_compute_cluster_node_type)
  - [databricks_compute_cluster_policy](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_compute_cluster_policy)
  - [databricks_compute_global_init_script](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_compute_global_init_script)
  - [databricks_compute_instance_pool](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_compute_instance_pool)
  - [databricks_compute_instance_profile](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_compute_instance_profile)
  - [databricks_compute_policy_family](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_compute_policy_family)
  - [databricks_files_dbfs](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_files_dbfs)
  - [databricks_iam_account_group](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_iam_account_group)
  - [databricks_iam_account_user](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_iam_account_user)
  - [databricks_iam_current_user](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_iam_current_user)
  - [databricks_iam_group](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_iam_group)
  - [databricks_iam_service_principal](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_iam_service_principal)
  - [databricks_iam_user](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_iam_user)
  - [databricks_jobs_job](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_jobs_job)
  - [databricks_jobs_job_run](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_jobs_job_run)
  - [databricks_ml_experiment](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_ml_experiment)
  - [databricks_ml_model](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_ml_model)
  - [databricks_ml_webhook](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_ml_webhook)
  - [databricks_pipelines_pipeline](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_pipelines_pipeline)
  - [databricks_pipelines_pipeline_event](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_pipelines_pipeline_event)
  - [databricks_pipelines_pipeline_update](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_pipelines_pipeline_update)
  - [databricks_serving_serving_endpoint](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_serving_serving_endpoint)
  - [databricks_settings_ip_access_list](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_settings_ip_access_list)
  - [databricks_settings_token](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_settings_token)
  - [databricks_settings_token_management](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_settings_token_management)
  - [databricks_sharing_provider](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_sharing_provider)
  - [databricks_sharing_recipient](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_sharing_recipient)
  - [databricks_sharing_share](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_sharing_share)
  - [databricks_sql_alert](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_sql_alert)
  - [databricks_sql_dashboard](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_sql_dashboard)
  - [databricks_sql_data_source](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_sql_data_source)
  - [databricks_sql_query](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_sql_query)
  - [databricks_sql_query_history](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_sql_query_history)
  - [databricks_sql_warehouse](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_sql_warehouse)
  - [databricks_sql_warehouse_config](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_sql_warehouse_config)
  - [databricks_workspace_git_credential](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_workspace_git_credential)
  - [databricks_workspace_repo](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_workspace_repo)
  - [databricks_workspace_scope](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_workspace_scope)
  - [databricks_workspace_secret](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_workspace_secret)
  - [databricks_workspace_workspace](https://hub.steampipe.io/plugins/turbot/databricks/tables/databricks_workspace_workspace)
