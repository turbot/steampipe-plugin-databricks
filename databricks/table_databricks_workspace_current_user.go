package databricks

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksWorkspaceCurrentUser(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_workspace_current_user",
		Description: "Gets details for the current user of the workspace.",
		List: &plugin.ListConfig{
			Hydrate: getWorkspaceCurrentUser,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "Databricks user ID.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "user_name",
				Description: "Email address of the Databricks user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "active",
				Description: "Whether the user is active.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "display_name",
				Description: "String that represents a concatenation of given and family names.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "external_id",
				Description: "External ID of the user.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "emails",
				Description: "All the emails associated with the Databricks user.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "entitlements",
				Description: "All the entitlements associated with the Databricks user.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "groups",
				Description: "All the groups the user belongs to.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "name",
				Description: "Name of the user.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "roles",
				Description: "All the roles associated with the Databricks user.",
				Type:        proto.ColumnType_JSON,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DisplayName"),
			},
		},
	}
}

//// LIST FUNCTION

func getWorkspaceCurrentUser(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_workspace_current_user.getWorkspaceCurrentUser", "connection_error", err)
		return nil, err
	}

	user, err := client.CurrentUser.Me(ctx)
	if err != nil {
		logger.Error("databricks_workspace_current_user.getWorkspaceCurrentUser", "api_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, user)

	return nil, nil
}
