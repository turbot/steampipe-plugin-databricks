package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksWorkspaceDashboard(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_workspace_dashboard",
		Description: "Gets details for all the dashboards associated with a Databricks workspace.",
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:      "id",
					Require:   plugin.Optional,
					Operators: []string{"=", "<>"},
				},
				{
					Name:      "dashboard_name",
					Require:   plugin.Optional,
					Operators: []string{"=", "<>"},
				},
				{
					Name:      "display_name",
					Require:   plugin.Optional,
					Operators: []string{"=", "<>"},
				},
			},
			Hydrate: listWorkspaceDashboards,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getWorkspaceDashboard,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "Databricks dashboard ID.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "Email address of the Databricks dashboard.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "can_edit",
				Description: "Whether the authenticated user can edit the query definition.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "created_at",
				Description: "Timestamp when this dashboard was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "dashboard_filters_enabled",
				Description: "In the web application, query filters that share a name are coupled to a single selection box if this value is `true`.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "is_archived",
				Description: "Whether the dashboard is trashed.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "is_draft",
				Description: "Whether the dashboard is a draft.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "is_favorite",
				Description: "Whether the dashboard is a favorite.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "parent",
				Description: "The identifier of the parent folder containing the dashboard.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "permission_tier",
				Description: "The permission level of the dashboard.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "slug",
				Description: "The URL slug of the dashboard.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "updated_at",
				Description: "Timestamp when this dashboard was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "user_id",
				Description: "The ID of the user that created and owns this dashboard.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "options",
				Description: "The options for the dashboard.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "tags",
				Description: "The tags for the dashboard.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "user",
				Description: "The user that created and owns this dashboard.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "widgets",
				Description: "The widgets for the dashboard.",
				Type:        proto.ColumnType_JSON,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

//// LIST FUNCTION

func listWorkspaceDashboards(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Limiting the results
	maxLimit := int32(10000)
	if d.QueryContext.Limit != nil {
		limit := int32(*d.QueryContext.Limit)
		if limit < maxLimit {
			maxLimit = limit
		}
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_workspace_dashboard.listWorkspaceDashboards", "connection_error", err)
		return nil, err
	}

	filterQuals := []filterQualMap{
		{"name", "Name", "string"},
	}

	filter := buildQueryFilterFromQuals(filterQuals, d.Quals)

	request := sql.ListDashboardsRequest{
		PageSize: int(maxLimit),
		Page:     1,
		Q:        filter,
	}

	count := 0
	for {
		response, err := client.Dashboards.Impl().List(ctx, request)
		if err != nil {
			logger.Error("databricks_workspace_dashboard.listWorkspaceDashboards", "api_error", err)
			return nil, err
		}

		for _, item := range response.Results {
			d.StreamListItem(ctx, &item)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		count += len(response.Results)
		if count < response.Count {
			request.Page += 1
		} else {
			return nil, nil
		}
	}
}

//// HYDRATE FUNCTIONS

func getWorkspaceDashboard(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("id")

	// Return nil, if no input provided
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_workspace_dashboard.getWorkspaceDashboard", "connection_error", err)
		return nil, err
	}

	dashboard, err := client.Dashboards.GetByDashboardId(ctx, id)
	if err != nil {
		logger.Error("databricks_workspace_dashboard.getWorkspaceDashboard", "api_error", err)
		return nil, err
	}

	return dashboard, nil
}
