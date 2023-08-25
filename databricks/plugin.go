package databricks

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const pluginName = "steampipe-plugin-databricks"

// Plugin creates this (databricks) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:               pluginName,
		DefaultTransform:   transform.FromCamel().Transform(transform.NullIfZeroValue),
		DefaultRetryConfig: &plugin.RetryConfig{ShouldRetryErrorFunc: shouldRetryError([]string{"429"})},
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundError([]string{"INVALID_PARAMETER_VALUE", "RESOURCE_DOES_NOT_EXIST", "DOES_NOT_EXIST", "404"}),
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"databricks_catalog":                    tableDatabricksCatalog(ctx),
			"databricks_catalog_connection":         tableDatabricksCatalogConnection(ctx),
			"databricks_catalog_external_location":  tableDatabricksCatalogExternalLocation(ctx),
			"databricks_catalog_function":           tableDatabricksCatalogFunction(ctx),
			"databricks_catalog_metastore":          tableDatabricksCatalogMetastore(ctx),
			"databricks_catalog_schema":             tableDatabricksCatalogSchema(ctx),
			"databricks_catalog_storage_credential": tableDatabricksCatalogStorageCredential(ctx),
			"databricks_catalog_system_schema":      tableDatabricksCatalogSystemSchema(ctx),
			"databricks_catalog_table":              tableDatabricksCatalogTable(ctx),
			"databricks_catalog_volume":             tableDatabricksCatalogVolume(ctx),
			"databricks_compute_cluster":            tableDatabricksComputeCluster(ctx),
			"databricks_compute_cluster_node_type":  tableDatabricksComputeClusterNodeType(ctx),
			"databricks_compute_cluster_policy":     tableDatabricksComputeClusterPolicy(ctx),
			"databricks_compute_global_init_script": tableDatabricksComputeGlobalInitScript(ctx),
			"databricks_compute_instance_pool":      tableDatabricksComputeInstancePool(ctx),
			"databricks_compute_instance_profile":   tableDatabricksComputeInstanceProfile(ctx),
			"databricks_compute_policy_family":      tableDatabricksComputePolicyFamily(ctx),
			"databricks_files_dbfs":                 tableDatabricksFilesDbfs(ctx),
			"databricks_iam_account_group":          tableDatabricksIAMAccountGroup(ctx),
			"databricks_iam_account_user":           tableDatabricksIAMAccountUser(ctx),
			"databricks_iam_current_user":           tableDatabricksIAMCurrentUser(ctx),
			"databricks_iam_group":                  tableDatabricksIAMGroup(ctx),
			"databricks_iam_service_principal":      tableDatabricksIAMServicePrincipal(ctx),
			"databricks_iam_user":                   tableDatabricksIAMUser(ctx),
			"databricks_job":                        tableDatabricksJob(ctx),
			"databricks_job_run":                    tableDatabricksJobRun(ctx),
			"databricks_ml_experiment":              tableDatabricksMLExperiment(ctx),
			"databricks_ml_model":                   tableDatabricksMLModel(ctx),
			"databricks_ml_webhook":                 tableDatabricksMLWebhook(ctx),
			"databricks_pipeline":                   tableDatabricksPipeline(ctx),
			"databricks_pipeline_event":             tableDatabricksPipelineEvent(ctx),
			"databricks_pipeline_update":            tableDatabricksPipelineUpdate(ctx),
			"databricks_serving_serving_endpoint":   tableDatabricksServingServingEndpoint(ctx),
			"databricks_settings_ip_access_list":    tableDatabricksSettingsIpAccessList(ctx),
			"databricks_settings_token":             tableDatabricksSettingsToken(ctx),
			"databricks_settings_token_management":  tableDatabricksSettingsTokenManagement(ctx),
			"databricks_sharing_provider":           tableDatabricksSharingProvider(ctx),
			"databricks_sharing_recipient":          tableDatabricksSharingRecipient(ctx),
			"databricks_sharing_share":              tableDatabricksSharingShare(ctx),
			"databricks_sql_alert":                  tableDatabricksSQLAlert(ctx),
			"databricks_sql_dashboard":              tableDatabricksSQLDashboard(ctx),
			"databricks_sql_data_source":            tableDatabricksSQLDataSource(ctx),
			"databricks_sql_query":                  tableDatabricksSQLQuery(ctx),
			"databricks_sql_query_history":          tableDatabricksSQLQueryHistory(ctx),
			"databricks_sql_warehouse":              tableDatabricksSQLWarehouse(ctx),
			"databricks_sql_warehouse_config":       tableDatabricksSQLWarehouseConfig(ctx),
			"databricks_workspace_git_credential":   tableDatabricksWorkspaceGitCredential(ctx),
			"databricks_workspace_repo":             tableDatabricksWorkspaceRepo(ctx),
			"databricks_workspace_scope":            tableDatabricksWorkspaceScope(ctx),
			"databricks_workspace_secret":           tableDatabricksWorkspaceSecret(ctx),
			"databricks_workspace":                  tableDatabricksWorkspace(ctx),
		},
	}

	return p
}
