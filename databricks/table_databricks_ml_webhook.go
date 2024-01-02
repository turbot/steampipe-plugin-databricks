package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksMLWebhook(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_ml_webhook",
		Description: "List all registry webhooks.",
		List: &plugin.ListConfig{
			Hydrate:    listMLWebhooks,
			KeyColumns: plugin.OptionalColumns([]string{"events", "model_name"}),
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The ID of the webhook.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "model_name",
				Description: "Name of the model whose events would trigger this webhook.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "creation_timestamp",
				Description: "Timestamp recorded when this model was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "description",
				Description: "Description of the model.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "last_updated_timestamp",
				Description: "Timestamp recorded when this model was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "status",
				Description: "Status of the webhook.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "events",
				Description: "Events that can trigger a registry webhook.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "http_url_spec",
				Description: "The HTTP URL specification for the webhook.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "job_spec",
				Description: "The job specification for the webhook.",
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

//// LIST FUNCTION

func listMLWebhooks(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := getWorkspaceClient(ctx, d)
	if err != nil {
		logger.Error("databricks_ml_webhook.listMLWebhooks", "connection_error", err)
		return nil, err
	}

	request := ml.ListWebhooksRequest{}
	if d.EqualsQualString("model_name") != "" {
		request.ModelName = d.EqualsQualString("model_name")
	}

	var events []ml.RegistryWebhookEvent
	quals := d.Quals
	if quals["events"] != nil {
		for _, q := range quals["events"].Quals {
			events = append(events, ml.RegistryWebhookEvent(q.Value.GetStringValue()))
		}
		request.Events = events
	}

	for {
		response, err := client.ModelRegistry.Impl().ListWebhooks(ctx, request)
		if err != nil {
			logger.Error("databricks_ml_webhook.listMLWebhooks", "api_error", err)
			return nil, err
		}

		for _, item := range response.Webhooks {
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
