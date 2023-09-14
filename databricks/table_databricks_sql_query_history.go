package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableDatabricksSQLQueryHistory(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_sql_query_history",
		Description: "List the history of queries through SQL warehouses.",
		List: &plugin.ListConfig{
			Hydrate:    listSQLQueryHistory,
			KeyColumns: plugin.OptionalColumns([]string{"warehouse_id", "user_id", "status"}),
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "query_id",
				Description: "Databricks query ID.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "warehouse_id",
				Description: "The UUID that uniquely identifies this data source/SQL warehouse across the API.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "user_id",
				Description: "The ID of the user who ran the query.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "duration",
				Description: "Total execution time of the query from the client's point of view.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "endpoint_id",
				Description: "Alias for warehouse id.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "error_message",
				Description: "Message describing why the query could not complete.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "executed_as_user_id",
				Description: "The ID of the user whose credentials were used to run the query.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "executed_as_user_name",
				Description: "The email address or username of the user whose credentials were used to run the query.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "execution_end_time_ms",
				Description: "The time execution of the query ended.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "is_final",
				Description: "Whether more updates for the query are expected.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "lookup_key",
				Description: "A key that can be used to look up query details.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "plans_state",
				Description: "Whether plans exist for the execution, or the reason why they are missing.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "query_end_time_ms",
				Description: "The time the query ended.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "query_start_time_ms",
				Description: "The time the query started.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "query_text",
				Description: "The text of the query.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "rows_produced",
				Description: "The number of results returned by the query.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "spark_ui_url",
				Description: "URL to the query plan.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "statement_type",
				Description: "Type of statement for this query.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status",
				Description: "The status of the query.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "user_name",
				Description: "The email address or username of the user who ran the query.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "channel_used",
				Description: "Channel information for the SQL warehouse at the time of query execution.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "metrics",
				Description: "Metrics about query execution.",
				Type:        proto.ColumnType_JSON,
			},
		}),
	}
}

//// LIST FUNCTION

func listSQLQueryHistory(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Limiting the results
	maxLimit := 100
	if d.QueryContext.Limit != nil {
		limit := int(*d.QueryContext.Limit)
		if limit < maxLimit {
			maxLimit = limit
		}
	}

	request := sql.ListQueryHistoryRequest{
		MaxResults: maxLimit,
		FilterBy:   &sql.QueryFilter{},
	}
	if d.EqualsQualString("warehouse_id") != "" {
		if request.FilterBy.WarehouseIds == nil {
			request.FilterBy.WarehouseIds = make([]string, 0)
		}
		request.FilterBy.WarehouseIds = append(request.FilterBy.WarehouseIds, d.EqualsQualString("warehouse_id"))
	}
	if d.EqualsQuals["user_id"] != nil {
		if request.FilterBy.UserIds == nil {
			request.FilterBy.UserIds = make([]int, 0)
		}
		request.FilterBy.UserIds = append(request.FilterBy.UserIds, int(d.EqualsQuals["user_id"].GetInt64Value()))
	}
	if d.EqualsQualString("status") != "" {
		if request.FilterBy.WarehouseIds == nil {
			request.FilterBy.Statuses = make([]sql.QueryStatus, 0)
		}
		request.FilterBy.Statuses = append(request.FilterBy.Statuses, sql.QueryStatus(d.EqualsQualString("status")))
	}

	// Create client
	client, err := getWorkspaceClient(ctx, d)
	if err != nil {
		logger.Error("databricks_sql_query_history.listSQLQueryHistory", "connection_error", err)
		return nil, err
	}

	for {
		response, err := client.QueryHistory.Impl().List(ctx, request)
		if err != nil {
			logger.Error("databricks_sql_query_history.listSQLQueryHistory", "api_error", err)
			return nil, err
		}

		for _, item := range response.Res {
			d.StreamListItem(ctx, item)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		if response.HasNextPage {
			request.PageToken = response.NextPageToken
		} else {
			return nil, nil
		}
	}
}
