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
			"databricks_account_group":                  tableDatabricksAccountGroup(ctx),
			"databricks_account_user":                   tableDatabricksAccountUser(ctx),
			"databricks_workspace_alert":                tableDatabricksWorkspaceAlert(ctx),
			"databricks_workspace_catalog":              tableDatabricksWorkspaceCatalog(ctx),
			"databricks_workspace_cluster":              tableDatabricksWorkspaceCluster(ctx),
			"databricks_workspace_cluster_policy":       tableDatabricksWorkspaceClusterPolicy(ctx),
			"databricks_workspace_connection":           tableDatabricksWorkspaceConnection(ctx),
			"databricks_workspace_current_user":         tableDatabricksWorkspaceCurrentUser(ctx),
			"databricks_workspace_dashboard":            tableDatabricksWorkspaceDashboard(ctx),
			"databricks_workspace_data_source":          tableDatabricksWorkspaceDataSource(ctx),
			"databricks_workspace_dbfs":                 tableDatabricksWorkspaceDbfs(ctx),
			"databricks_workspace_experiment":           tableDatabricksWorkspaceExperiment(ctx),
			"databricks_workspace_external_location":    tableDatabricksWorkspaceExternalLocation(ctx),
			"databricks_workspace_function":             tableDatabricksWorkspaceFunction(ctx),
			"databricks_workspace_git_credential":       tableDatabricksWorkspaceGitCredential(ctx),
			"databricks_workspace_global_init_script":   tableDatabricksWorkspaceGlobalInitScript(ctx),
			"databricks_workspace_group":                tableDatabricksWorkspaceGroup(ctx),
			"databricks_workspace_instance_pool":        tableDatabricksWorkspaceInstancePool(ctx),
			"databricks_workspace_instance_profile":     tableDatabricksWorkspaceInstanceProfile(ctx),
			"databricks_workspace_ip_access_list":       tableDatabricksWorkspaceIpAccessList(ctx),
			"databricks_workspace_job":                  tableDatabricksWorkspaceJob(ctx),
			"databricks_workspace_job_run":              tableDatabricksWorkspaceJobRun(ctx),
			"databricks_workspace_metastore":            tableDatabricksWorkspaceMetastore(ctx),
			"databricks_workspace_model":                tableDatabricksWorkspaceModel(ctx),
			"databricks_workspace_pipeline":             tableDatabricksWorkspacePipeline(ctx),
			"databricks_workspace_pipeline_event":       tableDatabricksWorkspacePipelineEvent(ctx),
			"databricks_workspace_pipeline_update":      tableDatabricksWorkspacePipelineUpdate(ctx),
			"databricks_workspace_user":                 tableDatabricksWorkspaceUser(ctx),
			"databricks_workspace_webhook":              tableDatabricksWorkspaceWebhook(ctx),
		},
	}

	return p
}
