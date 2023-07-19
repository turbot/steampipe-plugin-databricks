package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksPipelinesPipelineEvent(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_pipelines_pipeline_event",
		Description: "Retrieves events for a pipeline.",
		List: &plugin.ListConfig{
			ParentHydrate: listPipelinesPipelines,
			Hydrate:       listPipelinesPipelineEvents,
			KeyColumns:    plugin.OptionalColumns([]string{"pipeline_id"}),
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "A time-based, globally unique id.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "pipeline_id",
				Description: "Unique identifier of pipeline.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "event_type",
				Description: "The type of event.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "level",
				Description: "The severity level of the event.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "maturity_level",
				Description: "The maturity level of the event.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "message",
				Description: "The display message associated with the event.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "timestamp",
				Description: "The time of the event.",
				Type:        proto.ColumnType_TIMESTAMP,
			},

			// JSON fields
			{
				Name:        "error",
				Description: "Information about an error captured by the event.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "origin",
				Description: "Describes where the event originates from.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "sequence",
				Description: "A sequencing object to identify and order events.",
				Type:        proto.ColumnType_JSON,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Id"),
			},
		}),
	}
}

type pipelineEventInfo struct {
	pipelines.PipelineEvent
	PipelineId string
}

//// LIST FUNCTION

func listPipelinesPipelineEvents(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	pipelineId := h.Item.(pipelines.PipelineStateInfo).PipelineId

	if d.EqualsQualString("pipeline_id") != "" && d.EqualsQualString("pipeline_id") != pipelineId {
		return nil, nil
	}

	// Limiting the results
	maxLimit := int32(1000)
	if d.QueryContext.Limit != nil {
		limit := int32(*d.QueryContext.Limit)
		if limit < maxLimit {
			maxLimit = limit
		}
	}

	request := pipelines.ListPipelineEventsRequest{
		MaxResults: int(maxLimit),
		PipelineId: pipelineId,
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_pipelines_pipeline_event.listPipelinesPipelineEvents", "connection_error", err)
		return nil, err
	}

	for {
		response, err := client.Pipelines.Impl().ListPipelineEvents(ctx, request)
		if err != nil {
			logger.Error("databricks_pipelines_pipeline_event.listPipelinesPipelineEvents", "api_error", err)
			return nil, err
		}

		for _, item := range response.Events {
			d.StreamListItem(ctx, pipelineEventInfo{item, pipelineId})

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
