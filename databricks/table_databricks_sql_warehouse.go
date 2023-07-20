package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksSQLWarehouse(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_sql_warehouse",
		Description: "Gets a list of warehouses.",
		List: &plugin.ListConfig{
			Hydrate: listSQLWarehouses,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"id"}),
			Hydrate:    getSQLWarehouse,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "Unique identifier for warehouse.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "auto_stop_mins",
				Description: "The amount of time in minutes that a SQL warehouse must be idle before it is automatically stopped.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "cluster_size",
				Description: "Size of the clusters allocated for this warehouse.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "creator_name",
				Description: "Name of the creator of the warehouse.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "enable_photon",
				Description: "Whether the warehouse should use Photon optimized clusters.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "enable_serverless_compute",
				Description: "Whether the warehouse should use serverless compute.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "jdbc_url",
				Description: "JDBC URL for the warehouse.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "max_num_clusters",
				Description: "Maximum number of clusters that the autoscaler will create to handle concurrent queries.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "min_num_clusters",
				Description: "Minimum number of available clusters that will be maintained for this SQL warehouse.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "Logical name for the cluster.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "num_active_sessions",
				Description: "Current number of active sessions for the warehouse.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "num_clusters",
				Description: "Current number of clusters running for the service.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "spot_instance_policy",
				Description: "Configurations whether the warehouse should use spot instances.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "state",
				Description: "State of the warehouse.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "warehouse_type",
				Description: "Type of the warehouse.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "channel",
				Description: "Channel details for the warehouse.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "health",
				Description: "Health status of the warehouse.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "odbc_params",
				Description: "ODBC parameters for the warehouse.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "tags",
				Description: "A set of key-value pairs that will be tagged on all resources associated with the warehouse.",
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

func listSQLWarehouses(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	request := sql.ListWarehousesRequest{}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_sql_warehouse.listSQLWarehouses", "connection_error", err)
		return nil, err
	}

	warehouses, err := client.Warehouses.ListAll(ctx, request)
	if err != nil {
		logger.Error("databricks_sql_warehouse.listSQLWarehouses", "api_error", err)
		return nil, err
	}

	for _, item := range warehouses {
		d.StreamListItem(ctx, item)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getSQLWarehouse(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("id")

	// Return nil, if no input provided
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_sql_warehouse.getSQLWarehouse", "connection_error", err)
		return nil, err
	}

	warehouse, err := client.Warehouses.GetById(ctx, id)
	if err != nil {
		logger.Error("databricks_sql_warehouse.getSQLWarehouse", "api_error", err)
		return nil, err
	}
	return warehouse, nil
}
