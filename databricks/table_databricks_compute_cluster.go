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

func tableDatabricksComputeCluster(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_compute_cluster",
		Description: "Gets a list of clusters.",
		List: &plugin.ListConfig{
			Hydrate: listComputeClusters,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"cluster_id"}),
			Hydrate:    getComputeCluster,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "cluster_id",
				Description: "Canonical identifier for the cluster.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "cluster_name",
				Description: "Cluster name requested by the user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "autotermination_minutes",
				Description: "The number of minutes of inactivity after which Databricks automatically terminates this cluster.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "cluster_cores",
				Description: "Number of CPU cores available for this cluster.",
				Type:        proto.ColumnType_DOUBLE,
			},
			{
				Name:        "cluster_memory_mb",
				Description: "Total amount of cluster memory, in megabytes.",
				Type:        proto.ColumnType_DOUBLE,
			},
			{
				Name:        "cluster_source",
				Description: "Determines whether the cluster was created by a user through the UI, created by the Databricks Jobs Scheduler, or through an API request.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "creator_user_name",
				Description: "Creator user name. The field won't be included if the user has already been deleted.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "data_security_mode",
				Description: "The data security level of the cluster.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "driver_instance_pool_id",
				Description: "The optional ID of the instance pool for the driver of the cluster belongs.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "driver_node_type_id",
				Description: "The node type of the Spark driver.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "enable_elastic_disk",
				Description: "Autoscaling Local Storage: when enabled, this cluster will dynamically acquire additional disk space when its Spark workers are running low on disk space.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "enable_local_disk_encryption",
				Description: "Whether to enable LUKS on cluster VMs' local disks.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "instance_pool_id",
				Description: "The optional ID of the instance pool to which the cluster belongs.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "jdbc_port",
				Description: "Port on which Spark JDBC server is listening, in the driver node.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "last_restarted_time",
				Description: "The time when the cluster was started/restarted.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "last_state_loss_time",
				Description: "Time when the cluster driver last lost its state (due to a restart or driver failure).",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "node_type_id",
				Description: "This field encodes, through a single value, the resources available to each of the Spark nodes in this cluster.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "num_workers",
				Description: "Number of worker nodes that this cluster should have.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "policy_id",
				Description: "The ID of the cluster policy used to create the cluster if applicable.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "runtime_engine",
				Description: "Decides which runtime engine to be use.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "single_user_name",
				Description: "Single user name if data_security_mode is `SINGLE_USER`",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "spark_context_id",
				Description: "A canonical SparkContext identifier.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "spark_version",
				Description: "The Spark version of the cluster.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "start_time",
				Description: "The time when the cluster creation request was received.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "state",
				Description: "The current state of the cluster.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "state_message",
				Description: "The message associated with the most recent state transition.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "terminated_time",
				Description: "The time when the cluster was terminated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},

			// JSON fields
			{
				Name:        "autoscale",
				Description: "Parameters needed in order to automatically scale clusters up and down based on load.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "aws_attributes",
				Description: "Attributes related to clusters running on Amazon Web Services.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "azure_attributes",
				Description: "Attributes related to clusters running on Microsoft Azure.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "cluster_log_conf",
				Description: "The configuration for delivering spark logs to a long-term storage destination.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "cluster_log_status",
				Description: "The status of the cluster log delivery.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "custom_tags",
				Description: "Additional tags for cluster resources.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "default_tags",
				Description: "Tags that are added by Databricks regardless of any `custom_tags`.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "docker_image",
				Description: "The Docker image to use for every container in the cluster.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "driver",
				Description: "Node on which the Spark driver resides.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "executors",
				Description: "Nodes on which the Spark executors reside.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "gcp_attributes",
				Description: "Attributes related to clusters running on Google Cloud Platform.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "init_scripts",
				Description: "The configuration for storing init scripts.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "cluster_permissions",
				Description: "The permissions that the cluster has on the workspace.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getComputeClusterPermissions,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "spark_conf",
				Description: "An object containing a set of optional, user-specified Spark configuration key-value pairs.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "spark_env_vars",
				Description: "An object containing a set of optional, user-specified Spark environment variables.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "ssh_public_keys",
				Description: "SSH public key contents that will be added to each Spark node in this cluster.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "termination_reason",
				Description: "The reason why the cluster was terminated.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "workload_type_client",
				Description: "Defines what type of clients can use the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("WorkloadType.Client"),
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the resource.",
				Transform:   transform.FromField("ClusterName"),
				Type:        proto.ColumnType_STRING,
			},
		}),
	}
}

//// LIST FUNCTION

func listComputeClusters(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_compute_cluster.listComputeClusters", "connection_error", err)
		return nil, err
	}

	request := compute.ListClustersRequest{}

	clusters, err := client.Clusters.ListAll(ctx, request)
	if err != nil {
		logger.Error("databricks_compute_cluster.listComputeClusters", "api_error", err)
		return nil, err
	}

	for _, item := range clusters {
		d.StreamListItem(ctx, item)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getComputeCluster(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("cluster_id")

	// Return nil, if no input provided
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_compute_cluster.getComputeCluster", "connection_error", err)
		return nil, err
	}

	cluster, err := client.Clusters.GetByClusterId(ctx, id)
	if err != nil {
		logger.Error("databricks_compute_cluster.getComputeCluster", "api_error", err)
		return nil, err
	}
	return *cluster, nil
}

func getComputeClusterPermissions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := h.Item.(compute.ClusterDetails).ClusterId

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_compute_cluster.getComputeClusterPermission", "connection_error", err)
		return nil, err
	}

	request := iam.GetPermissionRequest{
		RequestObjectId:   id,
		RequestObjectType: "clusters",
	}

	permission, err := client.Permissions.Get(ctx, request)
	if err != nil {
		logger.Error("databricks_compute_cluster.getComputeClusterPermission", "api_error", err)
		return nil, err
	}
	return permission, nil
}
