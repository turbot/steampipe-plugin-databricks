package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksSettingsTokenManagement(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_settings_token_management",
		Description: "Lists all tokens associated with the specified workspace or user.",
		List: &plugin.ListConfig{
			Hydrate:    listSettingsTokenManagement,
			KeyColumns: plugin.OptionalColumns([]string{"created_by_id", "created_by_username"}),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("token_id"),
			Hydrate:    getSettingsTokenManagement,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "token_id",
				Description: "ID of the token.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_by_id",
				Description: "User id of the user that created the token.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "created_by_username",
				Description: "Username of the user that created the token.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "comment",
				Description: "Comment that describes the purpose of the token, specified by the token creator.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "creation_time",
				Description: "Timestamp when the token was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "expiry_time",
				Description: "Timestamp when the token expires.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "owner_id",
				Description: "User id of the user that owns the token.",
				Type:        proto.ColumnType_INT,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("TokenId"),
			},
		}),
	}
}

//// LIST FUNCTION

func listSettingsTokenManagement(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	request := settings.ListTokenManagementRequest{}
	if d.EqualsQualString("created_by_id") != "" {
		request.CreatedById = d.EqualsQualString("created_by_id")
	}
	if d.EqualsQualString("created_by_username") != "" {
		request.CreatedByUsername = d.EqualsQualString("created_by_user_name")
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_settings_token_management.listSettingsTokenManagement", "connection_error", err)
		return nil, err
	}

	tokens, err := client.TokenManagement.ListAll(ctx, request)
	if err != nil {
		logger.Error("databricks_settings_token_management.listSettingsTokenManagement", "api_error", err)
		return nil, err
	}

	for _, item := range tokens {
		d.StreamListItem(ctx, item)

		// Context can be cancelled due to manual cancellation or if the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getSettingsTokenManagement(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("token_id")

	// Return nil, if no input provided
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_settings_token_management.getSettingsTokenManagement", "connection_error", err)
		return nil, err
	}

	token, err := client.TokenManagement.GetByTokenId(ctx, id)
	if err != nil {
		logger.Error("databricks_settings_token_management.getSettingsTokenManagement", "api_error", err)
		return nil, err
	}
	return *token, nil
}
