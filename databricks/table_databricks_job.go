package databricks

import (
	"context"
	"strconv"

	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksJob(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_job",
		Description: "Get details for all the jobs associated with a Databricks workspace.",
		List: &plugin.ListConfig{
			Hydrate:    listJobs,
			KeyColumns: plugin.OptionalColumns([]string{"name"}),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("job_id"),
			Hydrate:    getJob,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "job_id",
				Description: "The canonical identifier for this job.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "name",
				Description: "The name of this job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Settings.Name"),
			},
			{
				Name:        "created_time",
				Description: "The time at which this job was created in epoch milliseconds.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "creator_user_name",
				Description: "The creator user name.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "format",
				Description: "The format of this job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Settings.Format"),
			},
			{
				Name:        "max_concurrent_runs",
				Description: "The maximum number of concurrent runs for this job.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Settings.MaxConcurrentRuns"),
			},
			{
				Name:        "run_as_user_name",
				Description: "The email of an active workspace user or the application ID of a service principal that the job runs as.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "timeout_seconds",
				Description: "An optional timeout applied to each run of this job.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Settings.TimeoutSeconds"),
			},

			// JSON fields
			{
				Name:        "compute",
				Description: "A list of compute requirements that can be referenced by tasks of this job.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Settings.Compute"),
			},
			{
				Name:        "continuous",
				Description: "An optional continuous property for this job.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Settings.Continuous"),
			},
			{
				Name:        "email_notifications",
				Description: "An optional set of email addresses that is notified when runs of this job begin or complete as well as when this job is deleted.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Settings.EmailNotifications"),
			},
			{
				Name:        "git_source",
				Description: "An optional specification for a remote repository containing the notebooks used by this job's notebook tasks.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Settings.GitSource"),
			},
			{
				Name:        "job_clusters",
				Description: "A list of job cluster specifications that can be shared and reused by tasks of this job.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Settings.JobClusters"),
			},
			{
				Name:        "job_permissions",
				Description: "A list of job-level permissions.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getJobPermissions,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "notification_settings",
				Description: "Optional notification settings that are used when sending notifications to each of the `email_notifications` and `webhook_notifications` for this job.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Settings.NotificationSettings"),
			},
			{
				Name:        "parameters",
				Description: "Job-level parameter definitions.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Settings.Parameters"),
			},
			{
				Name:        "run_as",
				Description: "Specifies the user or service principal that the job runs as.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Settings.RunAs"),
			},
			{
				Name:        "schedule",
				Description: "An optional periodic schedule for this job.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Settings.Schedule"),
			},
			{
				Name:        "tags",
				Description: "A map of tags associated with the job.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Settings.Tags"),
			},
			{
				Name:        "tasks",
				Description: "A list of tasks that this job executes.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Settings.Tasks"),
			},
			{
				Name:        "trigger",
				Description: "Trigger settings for this job.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Settings.Trigger"),
			},
			{
				Name:        "trigger_history",
				Description: "History of the file arrival trigger associated with the job.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getJob,
			},
			{
				Name:        "webhook_notifications",
				Description: "A collection of system notification IDs to notify when the run begins or completes.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Settings.WebhookNotifications"),
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Settings.Name"),
			},
		}),
	}
}

//// LIST FUNCTION

func listJobs(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Limiting the results
	maxLimit := int32(100)
	if d.QueryContext.Limit != nil {
		limit := int32(*d.QueryContext.Limit)
		if limit < maxLimit {
			maxLimit = limit
		}
	}

	request := jobs.ListJobsRequest{
		ExpandTasks: true,
		Limit:       int(maxLimit),
	}
	if d.EqualsQualString("name") != "" {
		request.Name = d.EqualsQualString("name")
	}

	// Create client
	client, err := getWorkspaceClient(ctx, d)
	if err != nil {
		logger.Error("databricks_job.listJobs", "connection_error", err)
		return nil, err
	}

	for {
		response, err := client.Jobs.Impl().List(ctx, request)
		if err != nil {
			logger.Error("databricks_job.listJobs", "api_error", err)
			return nil, err
		}

		for _, item := range response.Jobs {
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

func getJob(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	var id int64
	if h.Item != nil {
		id = getJobId(h.Item)
	} else {
		id = d.EqualsQuals["job_id"].GetInt64Value()
	}

	// Return nil, if no input provided
	if id == 0 {
		return nil, nil
	}

	// Create client
	client, err := getWorkspaceClient(ctx, d)
	if err != nil {
		logger.Error("databricks_job.getJob", "connection_error", err)
		return nil, err
	}

	job, err := client.Jobs.GetByJobId(ctx, id)
	if err != nil {
		logger.Error("databricks_job.getJob", "api_error", err)
		return nil, err
	}

	return *job, nil
}

func getJobPermissions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := getJobId(h.Item)

	// Create client
	client, err := getWorkspaceClient(ctx, d)
	if err != nil {
		logger.Error("databricks_job.getJobPermissions", "connection_error", err)
		return nil, err
	}

	request := iam.GetPermissionRequest{
		RequestObjectId:   strconv.Itoa(int(id)),
		RequestObjectType: "jobs",
	}

	permission, err := client.Permissions.Get(ctx, request)
	if err != nil {
		logger.Error("databricks_job.getJobPermissions", "api_error", err)
		return nil, err
	}
	return permission, nil
}

func getJobId(item interface{}) int64 {
	switch item := item.(type) {
	case jobs.Job:
		return item.JobId
	case jobs.BaseJob:
		return item.JobId
	}
	return 0
}
