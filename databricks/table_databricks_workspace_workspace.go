package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksWorkspaceWorkspace(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_workspace_workspace",
		Description: "Lists all secret workspaces available in the workspace.",
		List: &plugin.ListConfig{
			Hydrate: listWorkspaceWorkspaces,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "object_id",
				Description: "Unique identifier for the object.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "created_at",
				Description: "The creation time of the workspace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "language",
				Description: "The language of the workspace.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "modified_at",
				Description: "The last modified time of the workspace.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "object_type",
				Description: "The type of the object in workspace.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "path",
				Description: "The absolute path of the workspace.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "size",
				Description: "The file size in bytes.",
				Type:        proto.ColumnType_INT,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Path"),
			},
		}),
	}
}

//// LIST FUNCTION

func listWorkspaceWorkspaces(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	request := workspace.ListWorkspaceRequest{}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_workspace_workspace.listWorkspaceWorkspaces", "connection_error", err)
		return nil, err
	}

	workspaces, err := client.Workspace.ListAll(ctx, request)
	if err != nil {
		logger.Error("databricks_workspace_workspace.listWorkspaceWorkspaces", "api_error", err)
		return nil, err
	}

	for _, item := range workspaces {
		d.StreamListItem(ctx, item)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
