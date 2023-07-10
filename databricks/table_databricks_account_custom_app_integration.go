package databricks

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksAccountCustomAppIntegration(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_account_custom_app_integration",
		Description: "Get the list of custom oauth app integrations for the specified Databricks account.",
		List: &plugin.ListConfig{
			Hydrate: listAccountCustomAppIntegrations,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"integration_id"}),
			Hydrate:    getAccountCustomAppIntegration,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "integration_id",
				Description: "ID of this custom app.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "client_id",
				Description: "OAuth client ID of the custom OAuth app.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "confidential",
				Description: "Indicates if an OAuth client-secret should be generated.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "name",
				Description: "Human-readable name of the budget.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "redirect_urls",
				Description: "List of OAuth redirect URLs.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "token_access_policy",
				Description: "Token access policy.",
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

func listAccountCustomAppIntegrations(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectDatabricksAccount(ctx, d)
	if err != nil {
		logger.Error("databricks_account_custom_app_integration.listAccountCustomAppIntegrations", "connection_error", err)
		return nil, err
	}

	integrations, err := client.CustomAppIntegration.ListAll(ctx)
	if err != nil {
		logger.Error("databricks_account_custom_app_integration.listAccountCustomAppIntegrations", "api_error", err)
		return nil, err
	}

	for _, item := range integrations {
		d.StreamListItem(ctx, &item)

		// Context can be cancelled due to manual cancellation or if the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getAccountCustomAppIntegration(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("integration_id")

	// Return nil, if no input provided
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksAccount(ctx, d)
	if err != nil {
		logger.Error("databricks_account_custom_app_integration.getAccountCustomAppIntegration", "connection_error", err)
		return nil, err
	}

	integration, err := client.CustomAppIntegration.GetByIntegrationId(ctx, id)
	if err != nil {
		logger.Error("databricks_account_custom_app_integration.getAccountCustomAppIntegration", "api_error", err)
		return nil, err
	}
	return integration, nil
}
