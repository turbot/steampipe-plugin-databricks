package databricks

import (
	"context"
	"errors"
	"os"

	"github.com/databricks/databricks-sdk-go"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func getAccountClient(ctx context.Context, d *plugin.QueryData) (*databricks.AccountClient, error) {
	i, err := getAccountClientCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	return i.(*databricks.AccountClient), nil
}

// Cached form of getClient, using the per-connection and parallel safe
// Memoize() method.
var getAccountClientCached = plugin.HydrateFunc(getAccountClientUncached).Memoize()

func getAccountClientUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	databricksConfig := GetConfig(d.Connection)
	config := &databricks.Config{}

	// Account ID is required for a common column
	if databricksConfig.AccountId != nil {
		config.AccountID = *databricksConfig.AccountId
	}

	if config.AccountID == "" && os.Getenv("DATABRICKS_ACCOUNT_ID") == "" {
		return nil, errors.New("account_id must be configured")
	}

	// Check for profile and config file first
	if databricksConfig.Profile != nil {
		config.Profile = *databricksConfig.Profile
	}

	if databricksConfig.ConfigFilePath != nil {
		config.ConfigFile = *databricksConfig.ConfigFilePath
	}

	// If not using a profile and config file, check for OAuth config or token
	if config.ConfigFile == "" && os.Getenv("DATABRICKS_CONFIG_PROFILE") == "" {

		// Account host is required but can be set in the profile config
		if databricksConfig.AccountHost != nil {
			config.Host = *databricksConfig.AccountHost
		}

		if databricksConfig.ClientID != nil && databricksConfig.ClientSecret != nil {
			config.ClientID = *databricksConfig.ClientID
			config.ClientSecret = *databricksConfig.ClientSecret
		}

		empty_oauth_config := config.ClientID == "" && config.ClientSecret == "" && os.Getenv("DATABRICKS_CLIENT_ID") == "" && os.Getenv("DATABRICKS_CLIENT_SECRET") == ""

                // If not using OAuth config, check for token
		if empty_oauth_config && databricksConfig.WorkspaceToken != nil {
			config.Token = *databricksConfig.WorkspaceToken
		}

		// Finally, check for a username and password
		if empty_oauth_config && config.Token == "" && os.Getenv("DATABRICKS_TOKEN") == "" {
			if databricksConfig.Username != nil {
				config.Username = *databricksConfig.Username
			}
			if databricksConfig.Password != nil {
				config.Password = *databricksConfig.Password
			}
		}
	}

	client, err := databricks.NewAccountClient(config)
	if err != nil {
		plugin.Logger(ctx).Error("Unable to initialize account client:", err.Error())
		return nil, err
	}

	return client, nil
}

func getWorkspaceClient(ctx context.Context, d *plugin.QueryData) (*databricks.WorkspaceClient, error) {
	i, err := getWorkspacetClientCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	return i.(*databricks.WorkspaceClient), nil
}

// Cached form of getClient, using the per-connection and parallel safe
// Memoize() method.
var getWorkspacetClientCached = plugin.HydrateFunc(getWorkspacetClientUncached).Memoize()

func getWorkspacetClientUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	databricksConfig := GetConfig(d.Connection)
	config := &databricks.Config{}

	// Account ID is required for a common column
	if databricksConfig.AccountId != nil {
		config.AccountID = *databricksConfig.AccountId
	}

	if config.AccountID == "" && os.Getenv("DATABRICKS_ACCOUNT_ID") == "" {
		return nil, errors.New("account_id must be configured")
	}

	// Check for profile and config file first
	if databricksConfig.Profile != nil {
		config.Profile = *databricksConfig.Profile
	}

	if databricksConfig.ConfigFilePath != nil {
		config.ConfigFile = *databricksConfig.ConfigFilePath
	}

        // If not using a profile and config file, check for OAuth config or token
	if config.ConfigFile == "" && os.Getenv("DATABRICKS_CONFIG_PROFILE") == "" {

		// Workspace host is required but can be set in the profile config
		if databricksConfig.WorkspaceHost != nil {
			config.Host = *databricksConfig.WorkspaceHost
		}

		if databricksConfig.ClientID != nil && databricksConfig.ClientSecret != nil {
			config.ClientID = *databricksConfig.ClientID
			config.ClientSecret = *databricksConfig.ClientSecret
		}

		empty_oauth_config := config.ClientID == "" && config.ClientSecret == "" && os.Getenv("DATABRICKS_CLIENT_ID") == "" && os.Getenv("DATABRICKS_CLIENT_SECRET") == ""

                // If not using OAuth config, check for token
		if empty_oauth_config && databricksConfig.WorkspaceToken != nil {
			config.Token = *databricksConfig.WorkspaceToken
		}

		// Finally, check for a username and password
		if empty_oauth_config && config.Token == "" && os.Getenv("DATABRICKS_TOKEN") == "" {
			if databricksConfig.Username != nil {
				config.Username = *databricksConfig.Username
			}
			if databricksConfig.Password != nil {
				config.Password = *databricksConfig.Password
			}
		}
	}

	client, err := databricks.NewWorkspaceClient(config)
	if err != nil {
		plugin.Logger(ctx).Error("Unable to initialize workspace client:", err.Error())
		return nil, err
	}

	return client, nil
}
