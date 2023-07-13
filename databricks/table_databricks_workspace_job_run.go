package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksWorkspaceJobRun(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_workspace_job_run",
		Description: "Gets details for all the job runs.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.OptionalColumns([]string{"job_id", "run_type"}),
			Hydrate:    listWorkspaceJobRuns,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("run_id"),
			Hydrate:    getWorkspaceJobRun,
		},
		Columns: []*plugin.Column{
			{
				Name:        "run_id",
				Description: "The canonical identifier of the run.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "run_name",
				Description: "An optional name for the run.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "attempt_number",
				Description: "The sequence number of this run attempt for a triggered job run.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "cleanup_duration",
				Description: "The time in milliseconds it took to terminate the cluster and clean up any associated artifacts.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "creator_user_name",
				Description: "The user who created this job run.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "end_time",
				Description: "The time at which this run ended.",
				Transform:   transform.FromGo().Transform(convertTimestamp),
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "execution_duration",
				Description: "The time in milliseconds it took to execute the commands in the JAR or notebook.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "job_id",
				Description: "The canonical identifier of the job that contains this run.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "number_in_job",
				Description: "A unique identifier for this job run. This is set to the same value as `run_id`.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "original_attempt_run_id",
				Description: "If this run is a retry of a prior run attempt, this field contains the run_id of the original attempt.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "run_duration",
				Description: "The time in milliseconds it took the job run and all of its repairs to finish.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "run_page_url",
				Description: "The URL to the detail page of the run.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "run_type",
				Description: "The type of run.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "setup_duration",
				Description: "The time in milliseconds it took to set up the cluster.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "start_time",
				Description: "The time at which this run started.",
				Transform:   transform.FromGo().Transform(convertTimestamp),
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "trigger",
				Description: "The trigger that triggered this run.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "cluster_instance",
				Description: "The cluster instance that was used to run this job.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "cluster_spec",
				Description: "A snapshot of the job's cluster specification when this run was created.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "continuous",
				Description: "The continuous trigger that triggered this run.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "git_source",
				Description: "An optional specification for a remote repository containing the notebooks used by this job's notebook tasks.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "job_clusters",
				Description: "A list of job cluster specifications that can be shared and reused by tasks of this job.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "job_parameters",
				Description: "Job-level parameters used in the run.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "overriding_parameters",
				Description: "The parameters used for this run.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "repair_history",
				Description: "The repair history of this job run.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "schedule",
				Description: "The schedule that triggered this run.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "state",
				Description: "The current state of the run.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "tasks",
				Description: "The list of tasks performed by the run.",
				Type:        proto.ColumnType_JSON,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the resource.",
				Transform:   transform.FromField("RunName"),
				Type:        proto.ColumnType_STRING,
			},
		},
	}
}

//// LIST FUNCTION

func listWorkspaceJobRuns(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Limiting the results
	maxLimit := int32(25)
	if d.QueryContext.Limit != nil {
		limit := int32(*d.QueryContext.Limit)
		if limit < maxLimit {
			maxLimit = limit
		}
	}

	request := jobs.ListRunsRequest{
		Limit:       int(maxLimit),
		ExpandTasks: true,
	}
	if d.EqualsQuals["job_id"] != nil {
		request.JobId = d.EqualsQuals["job_id"].GetInt64Value()
	}
	if d.EqualsQuals["run_type"] != nil {
		request.RunType = jobs.ListRunsRunType(d.EqualsQualString("run_type"))
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_workspace_job_run.listWorkspaceJobRuns", "connection_error", err)
		return nil, err
	}

	for {
		response, err := client.Jobs.Impl().ListRuns(ctx, request)
		if err != nil {
			logger.Error("databricks_workspace_job_run.listWorkspaceJobRuns", "api_error", err)
			return nil, err
		}

		for _, item := range response.Runs {
			d.StreamListItem(ctx, item)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		if response.HasMore {
			request.PageToken = response.NextPageToken
		} else {
			return nil, nil
		}
	}
}

//// HYDRATE FUNCTIONS

func getWorkspaceJobRun(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQuals["run_id"].GetInt64Value()

	// Return nil, if no input provided
	if id == 0 {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_workspace_job_run.getWorkspaceJobRun", "connection_error", err)
		return nil, err
	}

	request := jobs.GetRunRequest{
		RunId: id,
	}

	run, err := client.Jobs.GetRun(ctx, request)
	if err != nil {
		logger.Error("databricks_workspace_job_run.getWorkspaceJobRun", "api_error", err)
		return nil, err
	}

	return *run, nil
}
