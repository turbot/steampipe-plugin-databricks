package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksWorkspaceRepo(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_workspace_repo",
		Description: "Returns repos that the calling user has Manage permissions on.",
		List: &plugin.ListConfig{
			Hydrate:    listWorkspaceRepos,
			KeyColumns: plugin.AnyColumn([]string{"path"}),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"id"}),
			Hydrate:    getWorkspaceRepo,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "ID of the repo object in the workspace.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "path",
				Description: "Desired path for the repo in the workspace.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "branch",
				Description: "Branch that the local version of the repo is checked out to.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "head_commit_id",
				Description: "SHA-1 hash representing the commit ID of the current HEAD of the repo.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "provider",
				Description: "Git provider.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "url",
				Description: "URL of the Git repo to be linked.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "sparse_checkout_patterns",
				Description: "List of patterns to include for sparse checkout.",
				Type:        proto.ColumnType_JSON,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Path"),
			},
		}),
	}
}

//// LIST FUNCTION

func listWorkspaceRepos(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	request := workspace.ListReposRequest{}
	if d.EqualsQualString("path") != "" {
		request.PathPrefix = d.EqualsQualString("path")
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_workspace_repo.listWorkspaceRepos", "connection_error", err)
		return nil, err
	}

	for {
		response, err := client.Repos.Impl().List(ctx, request)
		if err != nil {
			logger.Error("databricks_workspace_repo.listWorkspaceRepos", "api_error", err)
			return nil, err
		}

		for _, item := range response.Repos {
			d.StreamListItem(ctx, item)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		if response.NextPageToken == "" {
			return nil, nil
		}
		request.NextPageToken = response.NextPageToken
	}
}

//// HYDRATE FUNCTIONS

func getWorkspaceRepo(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQuals["id"].GetInt64Value()

	// Return nil, if no input provided
	if id == 0 {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_workspace_repo.getWorkspaceRepo", "connection_error", err)
		return nil, err
	}

	repo, err := client.Repos.GetByRepoId(ctx, id)
	if err != nil {
		logger.Error("databricks_workspace_repo.getWorkspaceRepo", "api_error", err)
		return nil, err
	}
	return repo, nil
}
