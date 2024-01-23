package databricks

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableDatabricksSettingsToken(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_settings_token",
		Description: "List all the valid tokens for a user-workspace pair.",
		List: &plugin.ListConfig{
			Hydrate: listSettingsToken,
		},
		Columns: getTokenInfoColumns(),
	}
}

//// LIST FUNCTION

func listSettingsToken(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := getWorkspaceClient(ctx, d)
	if err != nil {
		logger.Error("databricks_settings_token.listSettingsToken", "connection_error", err)
		return nil, err
	}

	tokens, err := client.Tokens.ListAll(ctx)
	if err != nil {
		logger.Error("databricks_settings_token.listSettingsToken", "api_error", err)
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
