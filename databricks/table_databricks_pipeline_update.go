package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksPipelineUpdate(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_pipeline_update",
		Description: "List updates for an active pipeline.",
		List: &plugin.ListConfig{
			ParentHydrate: listPipelines,
			Hydrate:       listPipelineUpdates,
			KeyColumns:    plugin.OptionalColumns([]string{"pipeline_id"}),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"pipeline_id", "update_id"}),
			Hydrate:    getPipelineUpdate,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "update_id",
				Description: "Unique identifier of the update.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "pipeline_id",
				Description: "Unique identifier of pipeline.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "cause",
				Description: "What triggered this update.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "cluster_id",
				Description: "The ID of the cluster that the update is running on.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "creation_time",
				Description: "The time when this update was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "full_refresh",
				Description: "Whether to reset all tables before running the pipeline.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "state",
				Description: "The current state of the pipeline.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "config",
				Description: "The pipeline configuration with system defaults applied where unspecified by the user.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "full_refresh_selection",
				Description: "A list of tables to update with full refresh.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "refresh_selection",
				Description: "A list of tables to update without full refresh.",
				Type:        proto.ColumnType_JSON,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("UpdateId"),
			},
		}),
	}
}

//// LIST FUNCTION

func listPipelineUpdates(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	pipelineId := h.Item.(pipelines.PipelineStateInfo).PipelineId

	if d.EqualsQualString("pipeline_id") != "" && d.EqualsQualString("pipeline_id") != pipelineId {
		return nil, nil
	}

	// Limiting the results
	maxLimit := int32(100)
	if d.QueryContext.Limit != nil {
		limit := int32(*d.QueryContext.Limit)
		if limit < maxLimit {
			maxLimit = limit
		}
	}

	request := pipelines.ListUpdatesRequest{
		MaxResults: int(maxLimit),
		PipelineId: pipelineId,
	}

	// Create client
	client, err := getWorkspaceClient(ctx, d)
	if err != nil {
		logger.Error("databricks_pipeline_update.listPipelineUpdates", "connection_error", err)
		return nil, err
	}

	for {
		response, err := client.Pipelines.Impl().ListUpdates(ctx, request)
		if err != nil {
			logger.Error("databricks_pipeline_update.listPipelineUpdates", "api_error", err)
			return nil, err
		}

		for _, item := range response.Updates {
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

func getPipelineUpdate(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	pipelineId := d.EqualsQualString("pipeline_id")
	updateId := d.EqualsQualString("update_id")

	// Return nil, if no input provided
	if pipelineId == "" || updateId == "" {
		return nil, nil
	}

	// Create client
	client, err := getWorkspaceClient(ctx, d)
	if err != nil {
		logger.Error("databricks_pipeline_update.getPipelineUpdate", "connection_error", err)
		return nil, err
	}

	update, err := client.Pipelines.GetUpdateByPipelineIdAndUpdateId(ctx, pipelineId, updateId)
	if err != nil {
		logger.Error("databricks_pipeline_update.getPipelineUpdate", "api_error", err)
		return nil, err
	}
	return *update.Update, nil
}
