package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksWorkspaceSecret(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_workspace_secret",
		Description: "Lists the secret keys that are stored.",
		List: &plugin.ListConfig{
			ParentHydrate: listWorkspaceScopes,
			Hydrate:       listWorkspaceSecrets,
			KeyColumns:    plugin.OptionalColumns([]string{"scope_name"}),
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "scope_name",
				Description: "The name of the secret scope.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "key",
				Description: "A unique name to identify the secret.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "last_updated_timestamp",
				Description: "The last updated timestamp (in milliseconds) for the secret.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Key"),
			},
		}),
	}
}

type secretScopeInfo struct {
	ScopeName string
	workspace.SecretMetadata
}

//// LIST FUNCTION

func listWorkspaceSecrets(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	scope := h.Item.(workspace.SecretScope).Name

	if d.EqualsQualString("scope_name") != "" && d.EqualsQualString("scope_name") != scope {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_workspace_secret.listWorkspaceSecrets", "connection_error", err)
		return nil, err
	}

	secrets, err := client.Secrets.ListSecretsByScope(ctx, scope)
	if err != nil {
		logger.Error("databricks_workspace_secret.listWorkspaceSecrets", "api_error", err)
		return nil, err
	}

	for _, item := range secrets.Secrets {
		d.StreamListItem(ctx, secretScopeInfo{scope, item})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
