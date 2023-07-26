package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksComputeInstancePool(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_compute_instance_pool",
		Description: "Gets a list of instance pools with their statistics.",
		List: &plugin.ListConfig{
			Hydrate: listComputeInstancePools,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("instance_pool_id"),
			Hydrate:    getComputeInstancePool,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "instance_pool_id",
				Description: "Canonical unique identifier for the pool.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "instance_pool_name",
				Description: "Pool name requested by the user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "enable_elastic_disk",
				Description: "Autoscaling Local Storage: when enabled, this instances in this pool will dynamically acquire additional disk space when its Spark workers are running low on disk space.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "idle_instance_autotermination_minutes",
				Description: "Automatically terminates the extra instances in the pool cache after they are inactive for this time in minutes if min_idle_instances requirement is already met.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "max_capacity",
				Description: "Maximum number of instances that can be launched in this pool.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "min_idle_instances",
				Description: "Minimum number of idle instances to keep in the instance pool.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "node_type_id",
				Description: "This field encodes, through a single value, the resources available to each of the Spark nodes in this cluster.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "state",
				Description: "The current state of the instance pool.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "aws_attributes",
				Description: "Attributes related to instance pools running on Amazon Web Services.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "azure_attributes",
				Description: "Attributes related to instance pools running on Azure.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "custom_tags",
				Description: "Additional tags for pool resources.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "default_tags",
				Description: "Tags that are added by Databricks.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "disk_spec",
				Description: "Defines the specification of the disks that will be attached to all spark containers.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "gcp_attributes",
				Description: "Attributes related to instance pools running on Google Cloud Platform.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "instance_pool_fleet_attributes",
				Description: "The fleet related setting to power the instance pool.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "instance_pool_permission",
				Description: "The permission of the instance pool.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getComputeInstancePoolPermission,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "pending_instance_errors",
				Description: "List of error messages for the failed pending instances.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Status.PendingInstanceErrors"),
			},
			{
				Name:        "preloaded_docker_images",
				Description: "Custom Docker Image BYOC.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "preloaded_spark_versions",
				Description: "A list of preloaded Spark image versions for the pool.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "stats",
				Description: "Usage statistics about the instance pool.",
				Type:        proto.ColumnType_JSON,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("InstancePoolName"),
			},
		}),
	}
}

//// LIST FUNCTION

func listComputeInstancePools(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_compute_instance_pool.listComputeInstancePools", "connection_error", err)
		return nil, err
	}

	instancePools, err := client.InstancePools.ListAll(ctx)
	if err != nil {
		logger.Error("databricks_compute_instance_pool.listComputeInstancePools", "api_error", err)
		return nil, err
	}

	for _, item := range instancePools {
		d.StreamListItem(ctx, item)

		// Context can be cancelled due to manual cancellation or if the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getComputeInstancePool(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("instance_pool_id")

	// Return nil, if no input provided
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_compute_instance_pool.getComputeInstancePool", "connection_error", err)
		return nil, err
	}

	instancePool, err := client.InstancePools.GetByInstancePoolId(ctx, id)
	if err != nil {
		logger.Error("databricks_compute_instance_pool.getComputeInstancePool", "api_error", err)
		return nil, err
	}
	return *instancePool, nil
}

func getComputeInstancePoolPermission(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := getInstancePoolId(h.Item)

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_compute_instance_pool.getComputeInstancePoolPermission", "connection_error", err)
		return nil, err
	}

	request := iam.GetPermissionRequest{
		RequestObjectId:   id,
		RequestObjectType: "instance-pools",
	}

	permission, err := client.Permissions.Get(ctx, request)
	if err != nil {
		logger.Error("databricks_compute_instance_pool.getComputeInstancePoolPermission", "api_error", err)
		return nil, err
	}
	return permission, nil
}

func getInstancePoolId(item interface{}) string {
	switch item := item.(type) {
	case compute.InstancePoolAndStats:
		return item.InstancePoolId
	case compute.GetInstancePool:
		return item.InstancePoolId
	}
	return ""
}
