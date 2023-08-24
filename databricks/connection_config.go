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
	DataUsername   *string `cty:"username"`
	DataPassword   *string `cty:"password"`
        ClientId       *string `cty:"client_id"`
        ClientSecret   *string `cty:"client_secret"`
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
        "client_id": {
                Required: false,
                Type:     schema.TypeString,
        },
        "client_secret": {
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
