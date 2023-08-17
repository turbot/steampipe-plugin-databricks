package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksIAMServicePrincipal(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_iam_service_principal",
		Description: "List the set of service principals associated with a Databricks workspace.",
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:      "display_name",
					Require:   plugin.Optional,
					Operators: []string{"=", "<>"},
				},
			},
			Hydrate: listIAMServicePrincipals,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("id"),
			ShouldIgnoreError: isNotFoundError([]string{"SCIM_404"}),
			Hydrate:           getIAMServicePrincipal,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "Databricks service principal ID.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "display_name",
				Description: "String that represents a concatenation of given and family names.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "active",
				Description: "Whether the service principal is active.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "application_id",
				Description: "UUID relating to the service principal.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "external_id",
				Description: "External id of the service principal.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "entitlements",
				Description: "All the entitlements associated with the service principal.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "groups",
				Description: "All the groups associated with the service principal.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "roles",
				Description: "All the roles associated with the service principal.",
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

func listIAMServicePrincipals(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
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
		logger.Error("databricks_iam_service_principal.listIAMServicePrincipals", "connection_error", err)
		return nil, err
	}

	filterQuals := []filterQualMap{
		{"display_name", "displayName", "string"},
	}
	filter := buildQueryFilterFromQuals(filterQuals, d.Quals)

	request := iam.ListServicePrincipalsRequest{
		Count:      int(maxLimit),
		StartIndex: 1,
		Filter:     filter,
	}

	for {
		principals, err := client.ServicePrincipals.ListAll(ctx, request)
		if err != nil {
			logger.Error("databricks_iam_service_principal.listIAMServicePrincipals", "api_error", err)
			return nil, err
		}

		for _, item := range principals {
			d.StreamListItem(ctx, item)

			// Context can be cancelled due to manual cancellation or if the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		if len(principals) < request.Count {
			return nil, nil
		} else {
			request.StartIndex += request.Count
		}
	}
}

//// HYDRATE FUNCTIONS

func getIAMServicePrincipal(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("id")

	// Return nil, if no input provided
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_iam_service_principal.getIAMServicePrincipal", "connection_error", err)
		return nil, err
	}

	principal, err := client.ServicePrincipals.GetById(ctx, id)
	if err != nil {
		logger.Error("databricks_iam_service_principal.getIAMServicePrincipal", "api_error", err)
		return nil, err
	}
	return *principal, nil
}
