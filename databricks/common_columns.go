package databricks

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// Columns defined on every account-level resource
func commonColumnsForAccountResource() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "account_id",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getCommonColumns,
			Transform:   transform.FromCamel(),
			Description: "The Databricks Account ID in which the resource is located.",
		},
	}
}

func databricksAccountColumns(columns []*plugin.Column) []*plugin.Column {
	return append(columns, commonColumnsForAccountResource()...)
}

func getCommonColumns(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	config := GetConfig(d.Connection)
	if config.AccountId != nil {
		return databricksCommonColumnData{
			AccountId: *config.AccountId,
		}, nil
	}
	return nil, nil
}

type databricksCommonColumnData struct {
	AccountId string
}
