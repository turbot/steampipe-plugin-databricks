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

	// Default to using env vars (#2)
	// But prefer the config (#1)

	if databricksConfig.Profile != nil {
		os.Setenv("DATABRICKS_CONFIG_PROFILE", *databricksConfig.Profile)
		if databricksConfig.ConfigFilePath != nil {
			os.Setenv("DATABRICKS_CONFIG_FILE", *databricksConfig.ConfigFilePath)
		}
	} else if os.Getenv("DATABRICKS_CONFIG_PROFILE") == "" {
		if databricksConfig.AccountToken != nil {
			os.Setenv("DATABRICKS_TOKEN", *databricksConfig.AccountToken)
		} else if os.Getenv("DATABRICKS_TOKEN") == "" {
                        if databricksConfig.ClientId != nil {
                                os.Setenv("DATABRICKS_CLIENT_ID", *databricksConfig.ClientId)
                                os.Setenv("DATABRICKS_CLIENT_SECRET", *databricksConfig.ClientSecret)
                        } else {
                                if databricksConfig.DataUsername != nil {
                                        os.Setenv("DATABRICKS_USERNAME", *databricksConfig.DataUsername)
                                }
                                if databricksConfig.DataPassword != nil {
                                        os.Setenv("DATABRICKS_PASSWORD", *databricksConfig.DataPassword)
                                }
                        }
		}

		if databricksConfig.AccountHost != nil {
			os.Setenv("DATABRICKS_HOST", *databricksConfig.AccountHost)
		}

		if databricksConfig.AccountId != nil {
			os.Setenv("DATABRICKS_ACCOUNT_ID", *databricksConfig.AccountId)
		} else if os.Getenv("DATABRICKS_ACCOUNT_ID") == "" {
			return nil, errors.New("account_id must be configured")
		}
	}

	client, err := databricks.NewAccountClient()
	if err != nil {
		fmt.Println("Unable to initialize client:", err)
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

	// Default to using env vars (#2)
	// But prefer the config (#1)

	if databricksConfig.Profile != nil {
		os.Setenv("DATABRICKS_CONFIG_PROFILE", *databricksConfig.Profile)
		if databricksConfig.ConfigFilePath != nil {
			os.Setenv("DATABRICKS_CONFIG_FILE", *databricksConfig.ConfigFilePath)
		}
	} else if os.Getenv("DATABRICKS_CONFIG_PROFILE") == "" {
		if databricksConfig.WorkspaceToken != nil {
			os.Setenv("DATABRICKS_TOKEN", *databricksConfig.WorkspaceToken)
		} else if os.Getenv("DATABRICKS_TOKEN") == "" {
                        if databricksConfig.ClientId != nil {
                                os.Setenv("DATABRICKS_CLIENT_ID", *databricksConfig.ClientId)
                                os.Setenv("DATABRICKS_CLIENT_SECRET", *databricksConfig.ClientSecret)
                        } else {
                                if databricksConfig.DataUsername != nil {
                                        os.Setenv("DATABRICKS_USERNAME", *databricksConfig.DataUsername)
                                }
                                if databricksConfig.DataPassword != nil {
                                        os.Setenv("DATABRICKS_PASSWORD", *databricksConfig.DataPassword)
                                }
                        }
		}

		if databricksConfig.WorkspaceHost != nil {
			os.Setenv("DATABRICKS_HOST", *databricksConfig.WorkspaceHost)
		}

		if databricksConfig.AccountId != nil {
			os.Setenv("DATABRICKS_ACCOUNT_ID", *databricksConfig.AccountId)
		} else if os.Getenv("DATABRICKS_ACCOUNT_ID") == "" {
			return nil, errors.New("account_id must be configured")
		}
	}

	client, err := databricks.NewWorkspaceClient()
	if err != nil {
		fmt.Println("Unable to initialize client:", err)
		return nil, err
	}

	d.ConnectionManager.Cache.Set(cacheKey, client)

	return client, nil
}
