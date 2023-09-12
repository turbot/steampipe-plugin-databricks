package databricks

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/databricks/databricks-sdk-go"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func connectDatabricksAccount(ctx context.Context, d *plugin.QueryData) (*databricks.AccountClient, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "databricks_account_client"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*databricks.AccountClient), nil
	}

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

	// Next check for a token
	if config.ConfigFile == "" && os.Getenv("DATABRICKS_CONFIG_PROFILE") == "" {
		// Account host is required but can be set in the profile config
		if databricksConfig.AccountHost != nil {
			config.Host = *databricksConfig.AccountHost
		}

		if databricksConfig.AccountToken != nil {
			config.Token = *databricksConfig.AccountToken
		}

		// Finally, check for a username and password
		if config.Token == "" && os.Getenv("DATABRICKS_TOKEN") == "" {
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
		fmt.Println("Unable to initialize account client:", err.Error())
		return nil, err
	}

	d.ConnectionManager.Cache.Set(cacheKey, client)

	return client, nil
}

func connectDatabricksWorkspace(ctx context.Context, d *plugin.QueryData) (*databricks.WorkspaceClient, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "databricks_workspace_client"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*databricks.WorkspaceClient), nil
	}

	databricksConfig := GetConfig(d.Connection)
	config := &databricks.Config{}

	// Default to using env vars (#2)
	// But prefer the config (#1)

	if databricksConfig.Profile != nil {
		config.Profile = *databricksConfig.Profile
		if databricksConfig.ConfigFilePath != nil {
			config.ConfigFile = *databricksConfig.ConfigFilePath
		}
	} else if os.Getenv("DATABRICKS_CONFIG_PROFILE") == "" {
		if databricksConfig.WorkspaceToken != nil {
			config.Token = *databricksConfig.WorkspaceToken
		} else if os.Getenv("DATABRICKS_TOKEN") == "" {
			if databricksConfig.Username != nil {
				config.Username = *databricksConfig.Username
			}
			if databricksConfig.Password != nil {
				config.Password = *databricksConfig.Password
			}
		}
		if databricksConfig.WorkspaceHost != nil {
			config.Host = *databricksConfig.WorkspaceHost
		}
	}

	if databricksConfig.AccountId != nil {
		config.AccountID = *databricksConfig.AccountId
	} else if os.Getenv("DATABRICKS_ACCOUNT_ID") == "" {
		return nil, errors.New("account_id must be configured")
	}

	client, err := databricks.NewWorkspaceClient(config)
	if err != nil {
		fmt.Println("Unable to initialize workspace client:", err.Error())
		return nil, err
	}

	d.ConnectionManager.Cache.Set(cacheKey, client)

	return client, nil
}
