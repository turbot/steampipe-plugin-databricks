package databricks

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksWorkspaceGlobalInitScript(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_workspace_global_init_script",
		Description: "Get a list of all global init scripts for this workspace.",
		List: &plugin.ListConfig{
			Hydrate: listWorkspaceGlobalInitScripts,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("script_id"),
			Hydrate:    getWorkspaceGlobalInitScript,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "script_id",
				Description: "The global init script ID.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "The name of the script.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "The time the script was created.",
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "created_by",
				Description: "The user who created the script.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "enabled",
				Description: "Whether the script is enabled.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "position",
				Description: "The position of the script.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "script",
				Description: "The Base64-encoded content of the script.",
				Hydrate:     getWorkspaceGlobalInitScript,
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "updated_at",
				Description: "The time the script was last updated.",
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "updated_by",
				Description: "The user who last updated the script.",
				Type:        proto.ColumnType_STRING,
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

func listWorkspaceGlobalInitScripts(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_workspace_global_init_script.listWorkspaceGlobalInitScripts", "connection_error", err)
		return nil, err
	}

	scripts, err := client.GlobalInitScripts.ListAll(ctx)
	if err != nil {
		logger.Error("databricks_workspace_global_init_script.listWorkspaceGlobalInitScripts", "api_error", err)
		return nil, err
	}

	for _, item := range scripts {
		d.StreamListItem(ctx, item)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getWorkspaceGlobalInitScript(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("script_id")

	// Return nil, if no input provided
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_workspace_global_init_script.getWorkspaceGlobalInitScript", "connection_error", err)
		return nil, err
	}

	script, err := client.GlobalInitScripts.GetByScriptId(ctx, id)
	if err != nil {
		logger.Error("databricks_workspace_global_init_script.getWorkspaceGlobalInitScript", "api_error", err)
		return nil, err
	}
	return *script, nil
}
