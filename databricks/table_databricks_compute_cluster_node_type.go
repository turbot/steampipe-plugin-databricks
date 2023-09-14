package databricks

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksComputeClusterNodeType(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_compute_cluster_node_type",
		Description: "Returns a list of supported Spark node types.",
		List: &plugin.ListConfig{
			Hydrate: listComputeClusterNodeTypes,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "node_type_id",
				Description: "Unique identifier for this node type.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "category",
				Description: "Category of the node type.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "A string description associated with this node type.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "display_order",
				Description: "Display order of the node type.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "instance_type_id",
				Description: "An identifier for the type of hardware that this node runs on.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "is_deprecated",
				Description: "Whether the node type is deprecated.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "is_encrypted_in_transit",
				Description: "AWS specific, whether this instance supports encryption in transit, used for hipaa and pci workloads.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "is_graviton",
				Description: "Whether this instance is a graviton instance.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "is_hidden",
				Description: "Whether the node type is hidden.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "is_io_cache_enabled",
				Description: "Flag indicating whether I/O cache is enabled for the node type.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "memory_mb",
				Description: "Memory (in MB) available for this node type.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "num_cores",
				Description: "Number of cores for the node type.",
				Type:        proto.ColumnType_DOUBLE,
			},
			{
				Name:        "num_gpus",
				Description: "Number of GPUs for the node type.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "photon_driver_capable",
				Description: "Indicates whether this node type is capable of being a Photon driver.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "photon_worker_capable",
				Description: "Indicates whether this node type is capable of being a Photon worker.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "support_cluster_tags",
				Description: "Flag indicating whether the node type supports cluster tags.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "support_ebs_volumes",
				Description: "Flag indicating whether the node type supports EBS volumes.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "support_port_forwarding",
				Description: "Flag indicating whether the node type supports port forwarding.",
				Type:        proto.ColumnType_BOOL,
			},

			// JSON fields
			{
				Name:        "node_info_status",
				Description: "Node info status information.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("NodeInfo.Status"),
			},
			{
				Name:        "node_instance_type",
				Description: "Node instance type information.",
				Type:        proto.ColumnType_JSON,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("NodeTypeId"),
			},
		}),
	}
}

//// LIST FUNCTION

func listComputeClusterNodeTypes(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := getWorkspaceClient(ctx, d)
	if err != nil {
		logger.Error("databricks_compute_cluster_node_type.listComputeClusterNodeTypes", "connection_error", err)
		return nil, err
	}

	nodeTypes, err := client.Clusters.ListNodeTypes(ctx)
	if err != nil {
		logger.Error("databricks_compute_cluster_node_type.listComputeClusterNodeTypes", "api_error", err)
		return nil, err
	}

	for _, item := range nodeTypes.NodeTypes {
		d.StreamListItem(ctx, item)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}
