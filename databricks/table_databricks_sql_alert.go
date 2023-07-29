package databricks

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksSQLAlert(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_sql_alert",
		Description: "Gets a list of alerts.",
		List: &plugin.ListConfig{
			Hydrate: listSQLAlerts,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.AnyColumn([]string{"id"}),
			ShouldIgnoreError: isNotFoundError([]string{"400"}),
			Hydrate:           getSQLAlert,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "Databricks alert ID.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "Name of the alert.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "Timestamp when the alert was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "last_triggered_at",
				Description: "Timestamp when the alert was last triggered.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "parent",
				Description: "The identifier of the workspace folder containing the object.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "rearm",
				Description: "Number of seconds after being triggered before the alert rearms itself and can be triggered again. If `null`, alert will never be triggered again.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "updated_at",
				Description: "Timestamp when the alert was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
			},

			// JSON fields
			{
				Name:        "options",
				Description: "Alert configuration options",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "query",
				Description: "Query associated with the alert.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "state",
				Description: "State of the alert.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "user",
				Description: "User associated with the alert.",
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

func listSQLAlerts(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_sql_alert.listSQLAlerts", "connection_error", err)
		return nil, err
	}

	alerts, err := client.Alerts.List(ctx)
	if err != nil {
		logger.Error("databricks_sql_alert.listSQLAlerts", "api_error", err)
		return nil, err
	}

	for _, item := range alerts {
		d.StreamListItem(ctx, item)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getSQLAlert(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("id")

	// Return nil, if no input provided
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_sql_alert.getSQLAlert", "connection_error", err)
		return nil, err
	}

	alert, err := client.Alerts.GetByAlertId(ctx, id)
	if err != nil {
		logger.Error("databricks_sql_alert.getSQLAlert", "api_error", err)
		return nil, err
	}
	return alert, nil
}
