package databricks

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksWorkspaceMetastore(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_workspace_metastore",
		Description: "Gets an array of the available metastores.",
		List: &plugin.ListConfig{
			Hydrate: listWorkspaceMetastores,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("metastore_id"),
			Hydrate:    getWorkspaceMetastore,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "metastore_id",
				Description: "Unique identifier of metastore.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "The user-specified name of the metastore.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "cloud",
				Description: "Cloud vendor of the metastore home shard.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "Time at which this metastore was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "default_data_access_config_id",
				Description: "Unique identifier of the metastore's (Default) Data Access Configuration.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "delta_sharing_organization_name",
				Description: "The organization name of a Delta Sharing entity, to be used in Databricks-to-Databricks Delta Sharing as the official name.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "delta_sharing_recipient_token_lifetime_in_seconds",
				Description: "The lifetime of a delta sharing recipient token in seconds.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "delta_sharing_scope",
				Description: "The scope of Delta Sharing enabled for the metastore.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "global_metastore_id",
				Description: "Globally unique metastore ID across clouds and regions, of the form `cloud:region:metastore_id`.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "owner",
				Description: "The owner of the metastore.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "privilege_model_version",
				Description: "The privilege model version of the metastore.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "region",
				Description: "Cloud region which the metastore serves.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "storage_root",
				Description: "The storage root URL for metastore.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "storage_root_credential_id",
				Description: "UUID of storage credential to access the metastore storage_root.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "storage_root_credential_name",
				Description: "Name of storage credential to access the metastore storage_root.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "updated_at",
				Description: "Time at which this metastore was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
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

func listWorkspaceMetastores(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_workspace_metastore.listWorkspaceMetastores", "connection_error", err)
		return nil, err
	}

	metastores, err := client.Metastores.ListAll(ctx)
	if err != nil {
		logger.Error("databricks_workspace_metastore.listWorkspaceMetastores", "api_error", err)
		return nil, err
	}

	for _, item := range metastores {
		d.StreamListItem(ctx, item)

		// Context can be cancelled due to manual cancellation or if the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getWorkspaceMetastore(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("metastore_id")

	// Return nil, if no input provided
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_workspace_metastore.getWorkspaceMetastore", "connection_error", err)
		return nil, err
	}

	metastore, err := client.Metastores.GetById(ctx, id)
	if err != nil {
		logger.Error("databricks_workspace_metastore.getWorkspaceMetastore", "api_error", err)
		return nil, err
	}
	return *metastore, nil
}
