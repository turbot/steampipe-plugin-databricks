package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksIAMAccountGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_iam_account_group",
		Description: "Gets group details associated with a Databricks account.",
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:      "display_name",
					Require:   plugin.Optional,
					Operators: []string{"=", "<>"},
				},
			},
			Hydrate: listIAMAccountGroups,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("id"),
			ShouldIgnoreError: isNotFoundError([]string{"SCIM_404"}),
			Hydrate:           getIAMAccountGroup,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "Databricks group id.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "display_name",
				Description: "Human-readable name of the group.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "external_id",
				Description: "External id of the group.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "entitlements",
				Description: "All the entitlements associated with the group.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "groups",
				Description: "All the groups the group belongs to.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "members",
				Description: "Members of the group.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "meta",
				Description: "Container for the group identifier. Workspace local versus account.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "roles",
				Description: "All the roles associated with the group.",
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

func listIAMAccountGroups(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
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
		logger.Error("databricks_iam_account_group.listIAMAccountGroups", "connection_error", err)
		return nil, err
	}

	filterQuals := []filterQualMap{
		{"display_name", "displayName", "string"},
	}
	filter := buildQueryFilterFromQuals(filterQuals, d.Quals)

	request := iam.ListAccountGroupsRequest{
		Count:      int(maxLimit),
		StartIndex: 1,
		Filter:     filter,
	}

	for {
		groups, err := client.Groups.ListAll(ctx, request)
		if err != nil {
			logger.Error("databricks_iam_account_group.listIAMAccountGroups", "api_error", err)
			return nil, err
		}

		for _, item := range groups {
			d.StreamListItem(ctx, item)

			// Context can be cancelled due to manual cancellation or if the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		if len(groups) < request.Count {
			return nil, nil
		} else {
			request.StartIndex = request.StartIndex + request.Count
		}
	}
}

//// HYDRATE FUNCTIONS

func getIAMAccountGroup(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("id")

	// Return nil, if no input provided
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksAccount(ctx, d)
	if err != nil {
		logger.Error("databricks_iam_account_group.getIAMAccountGroup", "connection_error", err)
		return nil, err
	}

	group, err := client.Groups.GetById(ctx, id)
	if err != nil {
		logger.Error("databricks_iam_account_group.getIAMAccountGroup", "api_error", err)
		return nil, err
	}
	return *group, nil
}
