package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksComputePolicyFamily(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_compute_policy_family",
		Description: "Retrieve a list of policy families.",
		List: &plugin.ListConfig{
			Hydrate: listComputePolicyFamilies,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("policy_family_id"),
			Hydrate:    getComputePolicyFamily,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "policy_family_id",
				Description: "ID of the policy family.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "Name of the policy family.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "Human-readable description of the purpose of the policy family.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "definition",
				Description: "Policy definition document expressed in Databricks Cluster Policy Definition Language.",
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

func listComputePolicyFamilies(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Limiting the results
	maxLimit := int64(1000)
	if d.QueryContext.Limit != nil {
		limit := int64(*d.QueryContext.Limit)
		if limit < maxLimit {
			maxLimit = limit
		}
	}

	request := compute.ListPolicyFamiliesRequest{
		MaxResults: maxLimit,
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_compute_policy_family.listComputePolicyFamilies", "connection_error", err)
		return nil, err
	}

	for {
		response, err := client.PolicyFamilies.Impl().List(ctx, request)
		if err != nil {
			logger.Error("databricks_compute_policy_family.listComputePolicyFamilies", "api_error", err)
			return nil, err
		}

		for _, item := range response.PolicyFamilies {
			d.StreamListItem(ctx, item)

			// Context can be cancelled due to manual cancellation or if the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}

			if response.NextPageToken == "" {
				return nil, nil
			}
			request.PageToken = response.NextPageToken
		}
	}
}

//// HYDRATE FUNCTIONS

func getComputePolicyFamily(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("policy_family_id")

	// Return nil, if no input provided
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_compute_policy_family.getComputePolicyFamily", "connection_error", err)
		return nil, err
	}

	policyFamily, err := client.PolicyFamilies.GetByPolicyFamilyId(ctx, id)
	if err != nil {
		logger.Error("databricks_compute_policy_family.getComputePolicyFamily", "api_error", err)
		return nil, err
	}
	return *policyFamily, nil
}
