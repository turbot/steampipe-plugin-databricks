package databricks

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksWorkspaceGitCredential(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_workspace_git_credential",
		Description: "Lists the calling user's Git credentials.",
		List: &plugin.ListConfig{
			Hydrate: listWorkspaceGitCredentials,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"credential_id", "git_provider"}),
			Hydrate:    getWorkspaceGitCredential,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "credential_id",
				Description: "ID of the credential object in the workspace.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "git_provider",
				Description: "The git provider.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "git_username",
				Description: "The git username.",
				Type:        proto.ColumnType_STRING,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("GitUsername"),
			},
		}),
	}
}

//// LIST FUNCTION

func listWorkspaceGitCredentials(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_workspace_git_credential.listWorkspaceGitCredentials", "connection_error", err)
		return nil, err
	}

	creds, err := client.GitCredentials.ListAll(ctx)
	if err != nil {
		logger.Error("databricks_workspace_git_credential.listWorkspaceGitCredentials", "api_error", err)
		return nil, err
	}

	for _, item := range creds {
		d.StreamListItem(ctx, item)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getWorkspaceGitCredential(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_workspace_git_credential.getWorkspaceGitCredential", "connection_error", err)
		return nil, err
	}

	// Get by id if id provided as input
	if d.EqualsQuals["credential_id"] != nil {
		id := d.EqualsQuals["credential_id"].GetInt64Value()

		cred, err := client.GitCredentials.GetByCredentialId(ctx, id)
		if err != nil {
			logger.Error("databricks_workspace_git_credential.getWorkspaceGitCredential", "api_error", err)
			return nil, err
		}
		return *cred, nil
	}

	// Get by name if name provided as input
	if d.EqualsQuals["git_provider"] != nil {
		provider := d.EqualsQualString("git_provider")

		cred, err := client.GitCredentials.GetByGitProvider(ctx, provider)
		if err != nil {
			logger.Error("databricks_workspace_git_credential.getWorkspaceGitCredential", "api_error", err)
			return nil, err
		}
		return *cred, nil
	}

	return nil, nil
}
