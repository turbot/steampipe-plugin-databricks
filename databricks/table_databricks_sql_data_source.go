package databricks

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksSQLDataSource(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_sql_data_source",
		Description: "Retrieves a full list of SQL warehouses available in this workspace.",
		List: &plugin.ListConfig{
			Hydrate: listSQLDataSources,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The unique identifier for this data source / SQL warehouse.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "The string name of this data source / SQL warehouse as it appears in the Databricks SQL web application.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "pause_reason",
				Description: "The reason why the data source is paused.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "paused",
				Description: "Whether the data source is paused.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "supports_auto_limit",
				Description: "Whether the data source supports auto limit.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "syntax",
				Description: "The syntax of the data source.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "type",
				Description: "The type of the data source.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "view_only",
				Description: "Whether the data source is view only.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "warehouse_id",
				Description: "The ID of the associated SQL warehouse, if this data source is backed by a SQL warehouse.",
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

func listSQLDataSources(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_sql_data_source.listSQLDataSources", "connection_error", err)
		return nil, err
	}

	dataSources, err := client.DataSources.List(ctx)
	if err != nil {
		logger.Error("databricks_sql_data_source.listSQLDataSources", "api_error", err)
		return nil, err
	}

	for _, item := range dataSources {
		d.StreamListItem(ctx, item)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
