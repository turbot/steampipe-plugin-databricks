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
			"databricks_workspace_user":                 tableDatabricksWorkspaceUser(ctx),
		},
	}

	return p
}
