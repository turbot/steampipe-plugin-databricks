package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksComputeClusterPolicy(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_compute_cluster_policy",
		Description: "Gets an array of cluster policies.",
		List: &plugin.ListConfig{
			Hydrate: listComputeClusterPolicies,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("policy_id"),
			Hydrate:    getComputeClusterPolicy,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "policy_id",
				Description: "Canonical unique identifier for the Cluster Policy.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "Cluster Policy name requested by the user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at_timestamp",
				Description: "The timestamp (in millisecond) when this Cluster Policy was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "creator_user_name",
				Description: "Creator user name. The field won't be included if the user has already been deleted.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "definition",
				Description: "Policy definition document expressed in Databricks Cluster Policy Definition Language.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "description",
				Description: "Additional human-readable description of the cluster policy.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "is_default",
				Description: "If true, policy is a default policy created and managed by Databricks.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "max_clusters_per_user",
				Description: "Max number of clusters per user that can be active using this policy. If not present, there is no max limit.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "policy_family_definition_overrides",
				Description: "Policy definition JSON document expressed in Databricks Policy Definition Language.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "policy_family_id",
				Description: "ID of the policy family.",
				Type:        proto.ColumnType_STRING,
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

func listComputeClusterPolicies(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_compute_cluster_policy.listComputeClusterPolicies", "connection_error", err)
		return nil, err
	}

	request := compute.ListClusterPoliciesRequest{}

	policies, err := client.ClusterPolicies.ListAll(ctx, request)
	if err != nil {
		logger.Error("databricks_compute_cluster_policy.listComputeClusterPolicies", "api_error", err)
		return nil, err
	}

	for _, item := range policies {
		d.StreamListItem(ctx, item)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getComputeClusterPolicy(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("policy_id")

	// Return nil, if no input provided
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_compute_cluster_policy.getComputeClusterPolicy", "connection_error", err)
		return nil, err
	}

	policy, err := client.ClusterPolicies.GetByPolicyId(ctx, id)
	if err != nil {
		logger.Error("databricks_compute_cluster_policy.getComputeClusterPolicy", "api_error", err)
		return nil, err
	}
	return policy, nil
}
