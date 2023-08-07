package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksSQLQuery(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_sql_query",
		Description: "Gets a list of queries.",
		List: &plugin.ListConfig{
			Hydrate: listSQLQueries,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"id"}),
			Hydrate:    getSQLQuery,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "Databricks query ID.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "Name of the query.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "can_edit",
				Description: "Whether the authenticated user is allowed to edit the definition of this query.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "created_at",
				Description: "Timestamp when the query was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "data_source_id",
				Description: "The UUID that uniquely identifies this data source / SQL warehouse across the API.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "General description that conveys additional information about this query such as usage notes.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "is_archived",
				Description: "Whether the query is trashed.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "is_draft",
				Description: "Whether the query is a draft.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "is_favorite",
				Description: "Whether the query appears in the current user's favorites list.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "is_safe",
				Description: "Indicates if a query either does not use any text type parameters or uses a data source type where text type parameters are handled safely.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "last_modified_by_id",
				Description: "The user ID of the user who last modified the query.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "latest_query_data_id",
				Description: "If there is a cached result for this query and user, this field includes the query result ID.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "parent",
				Description: "The identifier of the parent folder containing the query.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "permission_tier",
				Description: "The permission tier of the query.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "query",
				Description: "The text of the query to be run.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "query_hash",
				Description: "A SHA-256 hash of the query text along with the authenticated user ID.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "updated_at",
				Description: "Timestamp when the query was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "user_id",
				Description: "The user ID of the user who created the query.",
				Type:        proto.ColumnType_INT,
			},

			// JSON fields
			{
				Name:        "last_modified_by",
				Description: "The user who last modified the query.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "options",
				Description: "The options for the query.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "tags",
				Description: "The tags associated with the query.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "user",
				Description: "The user who created the query.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "visualizations",
				Description: "The visualizations associated with the query.",
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

func listSQLQueries(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Limiting the results
	maxLimit := 100
	if d.QueryContext.Limit != nil {
		limit := int(*d.QueryContext.Limit)
		if limit < maxLimit {
			maxLimit = limit
		}
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_sql_query.listSQLQueries", "connection_error", err)
		return nil, err
	}

	request := sql.ListQueriesRequest{
		PageSize: maxLimit,
		Page:     0,
	}

	for {
		response, err := client.Queries.Impl().List(ctx, request)
		if err != nil {
			logger.Error("databricks_sql_query.listSQLQueries", "api_error", err)
			return nil, err
		}

		for _, item := range response.Results {
			d.StreamListItem(ctx, item)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		if len(response.Results) < int(maxLimit) {
			return nil, nil
		} else {
			request.Page += 1
		}
	}
}

//// HYDRATE FUNCTIONS

func getSQLQuery(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("id")

	// Return nil, if no input provided
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_sql_query.getSQLQuery", "connection_error", err)
		return nil, err
	}

	query, err := client.Queries.GetByQueryId(ctx, id)
	if err != nil {
		logger.Error("databricks_sql_query.getSQLQuery", "api_error", err)
		return nil, err
	}
	return *query, nil
}
