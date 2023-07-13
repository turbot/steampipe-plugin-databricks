package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksAccountUser(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_account_user",
		Description: "Gets details for all the users associated with a Databricks account.",
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
			Hydrate: listAccountUsers,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getAccountUser,
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
			// {
			// 	Name:        "access_control_assignable_roles",
			// 	Description: "All the roles that can be assigned to the Databricks user.",
			// 	Hydrate:     getAccessControlAssignableRoles,
			// 	Transform:   transform.FromField("Roles"),
			// 	Type:        proto.ColumnType_JSON,
			// },
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

func listAccountUsers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
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
	client, err := connectDatabricksAccount(ctx, d)
	if err != nil {
		logger.Error("databricks_account_user.listAccountUsers", "connection_error", err)
		return nil, err
	}

	filterQuals := []filterQualMap{
		{"id", "id", "string"},
		{"user_name", "userName", "string"},
		{"display_name", "displayName", "string"},
	}

	filter := buildQueryFilterFromQuals(filterQuals, d.Quals)

	request := iam.ListAccountUsersRequest{
		Count:      int(maxLimit),
		StartIndex: 1,
		Filter:     filter,
	}

	for {
		users, err := client.Users.ListAll(ctx, request)
		if err != nil {
			logger.Error("databricks_account_user.listAccountUsers", "api_error", err)
			return nil, err
		}

		for _, item := range users {
			d.StreamListItem(ctx, item)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		if len(users) < request.Count {
			return nil, nil
		} else {
			request.StartIndex = request.StartIndex + request.Count
		}
	}
}

//// HYDRATE FUNCTIONS

func getAccountUser(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("id")

	// Return nil, if no input provided
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksAccount(ctx, d)
	if err != nil {
		logger.Error("databricks_account_user.getAccountUser", "connection_error", err)
		return nil, err
	}

	user, err := client.Users.GetById(ctx, id)
	if err != nil {
		logger.Error("databricks_account_user.getAccountUser", "api_error", err)
		return nil, err
	}

	return *user, nil
}

// func getAccessControlAssignableRoles(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
// 	logger := plugin.Logger(ctx)
// 	id := h.Item.(*iam.User).Id

// 	// Return nil, if no input provided
// 	if id == "" {
// 		return nil, nil
// 	}

// 	// Create client
// 	client, err := connectDatabricksAccount(ctx, d)
// 	if err != nil {
// 		logger.Error("databricks_account_user.getUser", "connection_error", err)
// 		return nil, err
// 	}

// 	request := iam.GetAssignableRolesForResourceRequest{
// 		Resource: id,
// 	}

// 	role, err := client.AccessControl.GetAssignableRolesForResource(ctx, request)
// 	if err != nil {
// 		logger.Error("databricks_account_user.getUser", "api_error", err)
// 		return nil, err
// 	}
// 	return role, nil
// }
