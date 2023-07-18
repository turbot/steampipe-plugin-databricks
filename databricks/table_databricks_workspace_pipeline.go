package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksWorkspacePipeline(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_workspace_pipeline",
		Description: "Lists pipelines defined in the Delta Live Tables system.",
		List: &plugin.ListConfig{
			Hydrate: listWorkspacePipelines,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("pipeline_id"),
			Hydrate:    getWorkspacePipeline,
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
				Name:        "cluster_id",
				Description: "The unique identifier of the cluster running the pipeline.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "creator_user_name",
				Description: "The user who created the pipeline.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "run_as_user_name",
				Description: "The username that the pipeline runs as.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "state",
				Description: "The current state of the pipeline.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "latest_updates",
				Description: "Status of the latest updates for the pipeline.",
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

func listWorkspacePipelines(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
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
		logger.Error("databricks_workspace_pipeline.listWorkspacePipelines", "connection_error", err)
		return nil, err
	}

	for {
		response, err := client.Pipelines.Impl().ListPipelines(ctx, request)
		if err != nil {
			logger.Error("databricks_workspace_pipeline.listWorkspacePipelines", "api_error", err)
			return nil, err
		}

		for _, item := range response.Statuses {
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

func getWorkspacePipeline(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("pipeline_id")

	// Return nil, if no input provided
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_workspace_pipeline.getWorkspacePipeline", "connection_error", err)
		return nil, err
	}

	pipeline, err := client.Pipelines.GetByPipelineId(ctx, id)
	if err != nil {
		logger.Error("databricks_workspace_pipeline.getWorkspacePipeline", "api_error", err)
		return nil, err
	}
	return *pipeline, nil
}
