package databricks

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type databricksConfig struct {
	AccountToken   *string `cty:"account_token"`
	AccountHost    *string `cty:"account_host"`
	WorkspaceToken *string `cty:"workspace_token"`
	WorkspaceHost  *string `cty:"workspace_host"`
	AccountId      *string `cty:"account_id"`
	Profile        *string `cty:"profile"`
	ConfigFilePath *string `cty:"config_file_path"`
	Username       *string `cty:"username"`
	Password       *string `cty:"password"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"account_token": {
		Required: false,
		Type:     schema.TypeString,
	},
	"account_host": {
		Required: false,
		Type:     schema.TypeString,
	},
	"workspace_token": {
		Required: false,
		Type:     schema.TypeString,
	},
	"workspace_host": {
		Required: false,
		Type:     schema.TypeString,
	},
	"account_id": {
		Required: true,
		Type:     schema.TypeString,
	},
	"profile": {
		Required: false,
		Type:     schema.TypeString,
	},
	"config_file_path": {
		Required: false,
		Type:     schema.TypeString,
	},
	"username": {
		Required: false,
		Type:     schema.TypeString,
	},
	"password": {
		Required: false,
		Type:     schema.TypeString,
	},
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
