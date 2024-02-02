package databricks

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type databricksConfig struct {
	AccountToken   *string `hcl:"account_token"`
	AccountHost    *string `hcl:"account_host"`
	WorkspaceToken *string `hcl:"workspace_token"`
	WorkspaceHost  *string `hcl:"workspace_host"`
	AccountId      *string `hcl:"account_id"`
	Profile        *string `hcl:"profile"`
	ConfigFilePath *string `hcl:"config_file_path"`
	Username       *string `hcl:"username"`
	Password       *string `hcl:"password"`
        ClientID       *string `hcl:"client_id"`
        ClientSecret   *string `hcl:"client_secret"`
}

func ConfigInstance() interface{} {
	return &databricksConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) databricksConfig {
	if connection == nil || connection.Config == nil {
		return databricksConfig{}
	}
	config, _ := connection.Config.(databricksConfig)
	return config
}
