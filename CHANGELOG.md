## v1.1.1 [2025-04-18]

_Bug fixes_

- Fixed Linux AMD64 plugin build failures for `Postgres 14 FDW`, `Postgres 15 FDW`, and `SQLite Extension` by upgrading GitHub Actions runners from `ubuntu-20.04` to `ubuntu-22.04`.

## v1.1.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#85](https://github.com/turbot/steampipe-plugin-databricks/pull/85))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#85](https://github.com/turbot/steampipe-plugin-databricks/pull/85))

## v0.4.0 [2024-02-02]

_What's new?_

- Added `OAuth` config support to provide users the ability to set `OAuth secret client ID` and `OAuth secret value` of a service principal. For more information, please see [Databricks plugin configuration](https://hub.steampipe.io/plugins/turbot/databricks#configuration). ([#6](https://github.com/turbot/steampipe-plugin-databricks/pull/6)) (Thanks [@rinzool](https://github.com/rinzool) for the contribution!)
- Added `Config` object to directly pass credentials to the client. ([#10](https://github.com/turbot/steampipe-plugin-databricks/pull/10))

## v0.3.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#32](https://github.com/turbot/steampipe-plugin-databricks/pull/32))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#32](https://github.com/turbot/steampipe-plugin-databricks/pull/32))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-databricks/blob/main/docs/LICENSE). ([#32](https://github.com/turbot/steampipe-plugin-databricks/pull/32))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#31](https://github.com/turbot/steampipe-plugin-databricks/pull/31))

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
