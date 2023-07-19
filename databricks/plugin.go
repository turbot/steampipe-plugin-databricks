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
		Name:             pluginName,
		DefaultTransform: transform.FromCamel().Transform(transform.NullIfZeroValue),
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundError([]string{"does not exist"}),
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"databricks_account_budget":                 tableDatabricksAccountBudget(ctx),
			"databricks_account_credential":             tableDatabricksAccountCredential(ctx),
			"databricks_account_custom_app_integration": tableDatabricksAccountCustomAppIntegration(ctx),
			"databricks_account_encryption_key":         tableDatabricksAccountEncryptionKey(ctx),

			"databricks_catalog_catalog":            tableDatabricksCatalogCatalog(ctx),
			"databricks_catalog_connection":         tableDatabricksCatalogConnection(ctx),
			"databricks_catalog_external_location":  tableDatabricksCatalogExternalLocation(ctx),
			"databricks_catalog_function":           tableDatabricksCatalogFunction(ctx),
			"databricks_catalog_metastore":          tableDatabricksCatalogMetastore(ctx),
			"databricks_compute_cluster":            tableDatabricksComputeCluster(ctx),
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
			"databricks_iam_user":                   tableDatabricksIAMUser(ctx),
			"databricks_jobs_job":                   tableDatabricksJobsJob(ctx),
			"databricks_jobs_job_run":               tableDatabricksJobsJobRun(ctx),
			"databricks_ml_experiment":              tableDatabricksMLExperiment(ctx),
			"databricks_ml_model":                   tableDatabricksMLModel(ctx),
			"databricks_ml_webhook":                 tableDatabricksMLWebhook(ctx),
			"databricks_pipelines_pipeline":         tableDatabricksPipelinesPipeline(ctx),
			"databricks_pipelines_pipeline_event":   tableDatabricksPipelinesPipelineEvent(ctx),
			"databricks_pipelines_pipeline_update":  tableDatabricksPipelinesPipelineUpdate(ctx),
			"databricks_settings_ip_access_list":    tableDatabricksSettingsIpAccessList(ctx),
			"databricks_sharing_provider":           tableDatabricksSharingProvider(ctx),
			"databricks_sql_alert":                  tableDatabricksSQLAlert(ctx),
			"databricks_sql_dashboard":              tableDatabricksSQLDashboard(ctx),
			"databricks_sql_data_source":            tableDatabricksSQLDataSource(ctx),
			"databricks_workspace_git_credential":   tableDatabricksWorkspaceGitCredential(ctx),
		},
	}

	return p
}
