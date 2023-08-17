package databricks

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksComputeInstanceProfile(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_compute_instance_profile",
		Description: "List the instance profiles that the calling user can use to launch a cluster.",
		List: &plugin.ListConfig{
			Hydrate: listComputeInstanceProfiles,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "instance_profile_arn",
				Description: "The AWS ARN of the instance profile to register with Databricks.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "iam_role_arn",
				Description: "The AWS IAM role ARN of the role associated with the instance profile.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "is_meta_instance_profile",
				Description: "This validation uses AWS dry-run mode for the RunInstances API to determine whether the instance profile is valid.",
				Type:        proto.ColumnType_BOOL,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("InstanceProfileName"),
			},
		}),
	}
}

//// LIST FUNCTION

func listComputeInstanceProfiles(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_compute_instance_profile.listComputeInstanceProfiles", "connection_error", err)
		return nil, err
	}

	instanceProfiles, err := client.InstanceProfiles.ListAll(ctx)
	if err != nil {
		logger.Error("databricks_compute_instance_profile.listComputeInstanceProfiles", "api_error", err)
		return nil, err
	}

	for _, item := range instanceProfiles {
		d.StreamListItem(ctx, item)

		// Context can be cancelled due to manual cancellation or if the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}
