package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksMLExperiment(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_ml_experiment",
		Description: "Gets details for all the experiments associated with a Databricks workspace.",
		List: &plugin.ListConfig{
			Hydrate:    listMLExperiments,
			KeyColumns: plugin.OptionalColumns([]string{"lifecycle_stage"}),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("experiment_id"),
			Hydrate:    getMLExperiment,
		},
		Columns: []*plugin.Column{
			{
				Name:        "experiment_id",
				Description: "Unique identifier for the experiment.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "Human readable name that identifies the experiment.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "artifact_location",
				Description: "Location where experiment artifacts are stored.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "creation_time",
				Description: "Time when the experiment was created.",
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "last_update_time",
				Description: "Time when the experiment was last updated.",
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "lifecycle_stage",
				Description: "Current life cycle stage of the experiment.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "tags",
				Description: "Additional metadata key-value pairs.",
				Type:        proto.ColumnType_JSON,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

//// LIST FUNCTION

func listMLExperiments(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Limiting the results
	maxLimit := int32(1000)
	if d.QueryContext.Limit != nil {
		limit := int32(*d.QueryContext.Limit)
		if limit < maxLimit {
			maxLimit = limit
		}
	}

	request := ml.ListExperimentsRequest{
		MaxResults: int(maxLimit),
		ViewType:   string(ml.SearchExperimentsViewTypeAll),
	}
	if d.EqualsQualString("lifecycle_stage") != "" {
		if d.EqualsQualString("lifecycle_stage") == "active" {
			request.ViewType = string(ml.SearchExperimentsViewTypeActiveOnly)
		} else if d.EqualsQualString("lifecycle_stage") == "deleted" {
			request.ViewType = string(ml.SearchExperimentsViewTypeDeletedOnly)
		}
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_ml_experiment.listMLExperiments", "connection_error", err)
		return nil, err
	}

	for {
		response, err := client.Experiments.Impl().ListExperiments(ctx, request)
		if err != nil {
			logger.Error("databricks_ml_experiment.listMLExperiments", "api_error", err)
			return nil, err
		}

		for _, item := range response.Experiments {
			d.StreamListItem(ctx, item)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		if response.NextPageToken == "" {
			return nil, nil
		} else {
			request.PageToken = response.NextPageToken
		}
	}
}

//// HYDRATE FUNCTIONS

func getMLExperiment(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	name := d.EqualsQualString("name")

	// Return nil, if no input provided
	if name == "" {
		return nil, nil
	}

	request := ml.GetByNameRequest{
		ExperimentName: name,
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_ml_experiment.getMLExperiment", "connection_error", err)
		return nil, err
	}

	experiment, err := client.Experiments.GetByName(ctx, request)
	if err != nil {
		logger.Error("databricks_ml_experiment.getMLExperiment", "api_error", err)
		return nil, err
	}

	return *experiment, nil
}
