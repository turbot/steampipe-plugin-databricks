package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksWorkspaceScope(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_workspace_scope",
		Description: "List all secret scopes available in the workspace.",
		List: &plugin.ListConfig{
			Hydrate: listWorkspaceScopes,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "A unique name to identify the secret scope.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "backend_type",
				Description: "The type of secret scope backend.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "acls",
				Description: "The access control list for the secret scope.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getWorkspaceScopeAcls,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "keyvault_metadata",
				Description: "The metadata for the secret scope if the type is `AZURE_KEYVAULT`.",
				Type:        proto.ColumnType_JSON,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		}),
	}
}

//// LIST FUNCTION

func listWorkspaceScopes(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := getWorkspaceClient(ctx, d)
	if err != nil {
		logger.Error("databricks_workspace_scope.listWorkspaceScopes", "connection_error", err)
		return nil, err
	}

	scopes, err := client.Secrets.ListScopesAll(ctx)
	if err != nil {
		logger.Error("databricks_workspace_scope.listWorkspaceScopes", "api_error", err)
		return nil, err
	}

	for _, item := range scopes {
		d.StreamListItem(ctx, item)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getWorkspaceScopeAcls(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	scope := h.Item.(workspace.SecretScope).Name

	// Create client
	client, err := getWorkspaceClient(ctx, d)
	if err != nil {
		logger.Error("databricks_workspace_scope.getWorkspaceScopeAcls", "connection_error", err)
		return nil, err
	}

	acls, err := client.Secrets.ListAclsByScope(ctx, scope)
	if err != nil {
		logger.Error("databricks_workspace_scope.getWorkspaceScopeAcls", "api_error", err)
		return nil, err
	}

	return acls.Items, nil
}
