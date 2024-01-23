package databricks

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableDatabricksSQLWarehouseConfig(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_sql_warehouse_config",
		Description: "Gets the workspace level configuration that is shared by all SQL warehouses in a workspace.",
		List: &plugin.ListConfig{
			Hydrate: getSQLWarehouseConfig,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "google_service_account",
				Description: "Google Service Account used to pass to cluster to access Google Cloud Storage.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "instance_profile_arn",
				Description: "Instance profile used to pass IAM role to the cluster.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "security_policy",
				Description: "Security policy for warehouses.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "channel",
				Description: "Channel selection details.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "data_access_config",
				Description: "Spark confs for external hive metastore configuration JSON serialized.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "enabled_warehouse_types",
				Description: "List of Warehouse Types allowed in this workspace.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "sql_configuration_parameters",
				Description: "SQL configuration parameters.",
				Type:        proto.ColumnType_JSON,
			},
		}),
	}
}

//// LIST FUNCTION

func getSQLWarehouseConfig(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := getWorkspaceClient(ctx, d)
	if err != nil {
		logger.Error("databricks_sql_warehouse_config.getSQLWarehouseConfig", "connection_error", err)
		return nil, err
	}

	config, err := client.Warehouses.GetWorkspaceWarehouseConfig(ctx)
	if err != nil {
		logger.Error("databricks_sql_warehouse_config.getSQLWarehouseConfig", "api_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, config)

	return nil, nil
}
