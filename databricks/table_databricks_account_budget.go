package databricks

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksAccountBudget(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_account_budget",
		Description: "Gets budget details associated with a Databricks account.",
		List: &plugin.ListConfig{
			Hydrate: listAccountBudgets,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"budget_id", "name"}),
			Hydrate:    getAccountBudget,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "budget_id",
				Description: "Databricks budget ID.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "Human-readable name of the budget.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "creation_time",
				Description: "Time when the budget was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "end_date",
				Description: "End date of the budget.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "filter",
				Description: "SQL-like filter expression with workspaceId, SKU and tag. Usage in your account that matches this expression will be counted in this budget.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "period",
				Description: "Period length in years, months, weeks and/or days.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "start_date",
				Description: "Start date of the budget period calculation.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "target_amount",
				Description: "Target amount of the budget per period in USD.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "update_time",
				Description: "Time when the budget was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
			},

			// JSON fields
			{
				Name:        "alerts",
				Description: "All the alerts associated with the Databricks budget.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "status_daily",
				Description: "Amount used in the budget for each day (noncumulative).",
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

func listAccountBudgets(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectDatabricksAccount(ctx, d)
	if err != nil {
		logger.Error("databricks_account_budget.listAccountBudgets", "connection_error", err)
		return nil, err
	}

	budgets, err := client.Budgets.ListAll(ctx)
	if err != nil {
		logger.Error("databricks_account_budget.listAccountBudgets", "api_error", err)
		return nil, err
	}

	for _, item := range budgets {
		d.StreamListItem(ctx, item)

		// Context can be cancelled due to manual cancellation or if the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getAccountBudget(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("budget_id")

	// Return nil, if no input provided
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksAccount(ctx, d)
	if err != nil {
		logger.Error("databricks_account_budget.getAccountBudget", "connection_error", err)
		return nil, err
	}

	budget, err := client.Budgets.GetByBudgetId(ctx, id)
	if err != nil {
		logger.Error("databricks_account_budget.getAccountBudget", "api_error", err)
		return nil, err
	}
	return budget, nil

}
