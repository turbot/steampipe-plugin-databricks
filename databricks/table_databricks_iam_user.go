package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksIAMUser(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_iam_user",
		Description: "List details for all the users associated with a Databricks workspace.",
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:      "id",
					Require:   plugin.Optional,
					Operators: []string{"=", "<>"},
				},
				{
					Name:      "user_name",
					Require:   plugin.Optional,
					Operators: []string{"=", "<>"},
				},
				{
					Name:      "display_name",
					Require:   plugin.Optional,
					Operators: []string{"=", "<>"},
				},
			},
			Hydrate: listIAMUsers,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("id"),
			ShouldIgnoreError: isNotFoundError([]string{"SCIM_404"}),
			Hydrate:           getIAMUser,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "Databricks user ID.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "user_name",
				Description: "Email address of the Databricks user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "active",
				Description: "Whether the user is active.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "display_name",
				Description: "String that represents a concatenation of given and family names.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "external_id",
				Description: "External ID of the user.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "emails",
				Description: "All the emails associated with the Databricks user.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "entitlements",
				Description: "All the entitlements associated with the Databricks user.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "groups",
				Description: "All the groups the user belongs to.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "name",
				Description: "Name of the user.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "roles",
				Description: "All the roles associated with the Databricks user.",
				Type:        proto.ColumnType_JSON,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DisplayName"),
			},
		}),
	}
}

//// LIST FUNCTION

func listIAMUsers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
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
		logger.Error("databricks_iam_user.listIAMUsers", "connection_error", err)
		return nil, err
	}

	filterQuals := []filterQualMap{
		{"id", "id", "string"},
		{"user_name", "userName", "string"},
		{"display_name", "displayName", "string"},
	}

	filter := buildQueryFilterFromQuals(filterQuals, d.Quals)

	request := iam.ListUsersRequest{
		Count:      int(maxLimit),
		StartIndex: 1,
		Filter:     filter,
	}

	for {
		response, err := client.Users.Impl().List(ctx, request)
		if err != nil {
			logger.Error("databricks_iam_user.listIAMUsers", "api_error", err)
			return nil, err
		}

		for _, item := range response.Resources {
			d.StreamListItem(ctx, item)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		if response.ItemsPerPage < int64(request.Count) {
			return nil, nil
		} else {
			request.StartIndex = request.StartIndex + request.Count
		}
	}
}

//// HYDRATE FUNCTIONS

func getIAMUser(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("id")

	// Return nil, if no input provided
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_iam_user.getIAMUser", "connection_error", err)
		return nil, err
	}

	user, err := client.Users.GetById(ctx, id)
	if err != nil {
		logger.Error("databricks_iam_user.getIAMUser", "api_error", err)
		return nil, err
	}

	return *user, nil
}
