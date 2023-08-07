package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksPipelinesPipeline(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_pipelines_pipeline",
		Description: "List pipelines defined in the Delta Live Tables system.",
		List: &plugin.ListConfig{
			Hydrate: listPipelinesPipelines,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("pipeline_id"),
			Hydrate:    getPipelinesPipeline,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "pipeline_id",
				Description: "Unique identifier of pipeline.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "The user-friendly name of the pipeline.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "catalog",
				Description: "A catalog in Unity Catalog to publish data from this pipeline to.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getPipelinesPipeline,
				Transform:   transform.FromField("Spec.Catalog"),
			},
			{
				Name:        "cause",
				Description: "An optional message detailing the cause of the pipeline state.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getPipelinesPipeline,
			},
			{
				Name:        "channel",
				Description: "DLT Release Channel that specifies which version to use.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getPipelinesPipeline,
				Transform:   transform.FromField("Spec.Channel"),
			},
			{
				Name:        "cluster_id",
				Description: "The unique identifier of the cluster running the pipeline.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "continuous",
				Description: "Whether the pipeline is continuous or triggered.",
				Type:        proto.ColumnType_BOOL,
				Hydrate:     getPipelinesPipeline,
				Transform:   transform.FromField("Spec.Continuous"),
			},
			{
				Name:        "creator_user_name",
				Description: "The user who created the pipeline.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "development",
				Description: "Whether the pipeline is in Development mode.",
				Type:        proto.ColumnType_BOOL,
				Hydrate:     getPipelinesPipeline,
				Transform:   transform.FromField("Spec.Development"),
			},
			{
				Name:        "edition",
				Description: "Pipeline product edition.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getPipelinesPipeline,
				Transform:   transform.FromField("Spec.Edition"),
			},
			{
				Name:        "health",
				Description: "The health of the pipeline.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getPipelinesPipeline,
			},
			{
				Name:        "last_modified",
				Description: "The last time the pipeline settings were modified or created.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getPipelinesPipeline,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "photon",
				Description: "Whether photon is enabled for this pipeline.",
				Type:        proto.ColumnType_BOOL,
				Hydrate:     getPipelinesPipeline,
				Transform:   transform.FromField("Spec.Photon"),
			},
			{
				Name:        "run_as_user_name",
				Description: "The username that the pipeline runs as.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "serverless",
				Description: "Whether serverless compute is enabled for this pipeline.",
				Type:        proto.ColumnType_BOOL,
				Hydrate:     getPipelinesPipeline,
				Transform:   transform.FromField("Spec.Serverless"),
			},
			{
				Name:        "state",
				Description: "The current state of the pipeline.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "storage",
				Description: "DBFS root directory for storing checkpoints and tables.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getPipelinesPipeline,
				Transform:   transform.FromField("Spec.Storage"),
			},
			{
				Name:        "target",
				Description: "Target schema (database) to add tables in this pipeline to.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getPipelinesPipeline,
				Transform:   transform.FromField("Spec.Target"),
			},

			// JSON fields
			{
				Name:        "clusters",
				Description: "Cluster settings for this pipeline deployment.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPipelinesPipeline,
				Transform:   transform.FromField("Spec.Clusters"),
			},
			{
				Name:        "configuration",
				Description: "String-String configuration for this pipeline execution.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPipelinesPipeline,
				Transform:   transform.FromField("Spec.Configuration"),
			},
			{
				Name:        "filters",
				Description: "Filters on which Pipeline packages to include in the deployed graph.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPipelinesPipeline,
				Transform:   transform.FromField("Spec.Filters"),
			},
			{
				Name:        "latest_updates",
				Description: "Status of the latest updates for the pipeline.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "libraries",
				Description: "Libraries or code needed by this deployment.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPipelinesPipeline,
				Transform:   transform.FromField("Spec.Libraries"),
			},
			{
				Name:        "pipeline_permissions",
				Description: "Permissions for this pipeline.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPipelinesPipelinePermissions,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "trigger",
				Description: "Which pipeline trigger to use.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPipelinesPipeline,
				Transform:   transform.FromField("Spec.Trigger"),
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

func listPipelinesPipelines(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Limiting the results
	maxLimit := int32(100)
	if d.QueryContext.Limit != nil {
		limit := int32(*d.QueryContext.Limit)
		if limit < maxLimit {
			maxLimit = limit
		}
	}

	request := pipelines.ListPipelinesRequest{
		MaxResults: int(maxLimit),
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_pipelines_pipeline.listPipelinesPipelines", "connection_error", err)
		return nil, err
	}

	for {
		response, err := client.Pipelines.Impl().ListPipelines(ctx, request)
		if err != nil {
			logger.Error("databricks_pipelines_pipeline.listPipelinesPipelines", "api_error", err)
			return nil, err
		}

		for _, item := range response.Statuses {
			d.StreamListItem(ctx, item)

			// Context can be cancelled due to manual cancellation or if the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		if response.NextPageToken == "" {
			return nil, nil
		}
		request.PageToken = response.NextPageToken
	}
}

//// HYDRATE FUNCTIONS

func getPipelinesPipeline(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	var id string
	if h.Item != nil {
		id = getPipelineId(h.Item)
	} else {
		id = d.EqualsQualString("pipeline_id")
	}

	// Return nil, if no input provided
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_pipelines_pipeline.getPipelinesPipeline", "connection_error", err)
		return nil, err
	}

	pipeline, err := client.Pipelines.GetByPipelineId(ctx, id)
	if err != nil {
		logger.Error("databricks_pipelines_pipeline.getPipelinesPipeline", "api_error", err)
		return nil, err
	}
	return *pipeline, nil
}

func getPipelinesPipelinePermissions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := getPipelineId(h.Item)

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_pipelines_pipeline.getPipelinesPipelinePermissions", "connection_error", err)
		return nil, err
	}

	request := iam.GetPermissionRequest{
		RequestObjectId:   id,
		RequestObjectType: "pipelines",
	}

	permission, err := client.Permissions.Get(ctx, request)
	if err != nil {
		logger.Error("databricks_pipelines_pipeline.getPipelinesPipelinePermissions", "api_error", err)
		return nil, err
	}
	return permission, nil
}

func getPipelineId(item interface{}) string {
	switch item := item.(type) {
	case pipelines.PipelineStateInfo:
		return item.PipelineId
	case pipelines.GetPipelineResponse:
		return item.PipelineId
	}
	return ""
}
