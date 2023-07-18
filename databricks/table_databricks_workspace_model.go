package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksWorkspaceModel(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_workspace_model",
		Description: "Lists all available registered models.",
		List: &plugin.ListConfig{
			Hydrate: listWorkspaceModels,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getWorkspaceModel,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "Unique name for the model.",
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
				Name:        "user_id",
				Description: "User ID of the user who created this model.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "latest_versions",
				Description: "Collection of latest model versions for each stage.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "tags",
				Description: "Additional metadata key-value pairs for this model.",
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

func listWorkspaceModels(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Limiting the results
	maxLimit := int32(1000)
	if d.QueryContext.Limit != nil {
		limit := int32(*d.QueryContext.Limit)
		if limit < maxLimit {
			maxLimit = limit
		}
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_workspace_model.listWorkspaceModels", "connection_error", err)
		return nil, err
	}

	request := ml.ListModelsRequest{
		MaxResults: int(maxLimit),
	}

	for {
		response, err := client.ModelRegistry.Impl().ListModels(ctx, request)
		if err != nil {
			logger.Error("databricks_workspace_model.listWorkspaceModels", "api_error", err)
			return nil, err
		}

		for _, item := range response.RegisteredModels {
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

func getWorkspaceModel(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	name := d.EqualsQualString("name")

	// Return nil, if no input provided
	if name == "" {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_workspace_model.getWorkspaceModel", "connection_error", err)
		return nil, err
	}

	request := ml.GetModelRequest{
		Name: name,
	}

	model, err := client.ModelRegistry.GetModel(ctx, request)
	if err != nil {
		logger.Error("databricks_workspace_model.getWorkspaceModel", "api_error", err)
		return nil, err
	}
	return *model, nil
}
