package databricks

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
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

var getCommonColumnsMemoized = plugin.HydrateFunc(getCommonColumnsUncached).Memoize(memoize.WithCacheKeyFunction(getCommonColumnsCacheKey))

// Build a cache key for the call to getCommonColumnsCacheKey.
func getCommonColumnsCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "getCommonColumns"
	return key, nil
}

// declare a wrapper hydrate function to call the memoized function
// - this is required when a memoized function is used for a column definition
func getCommonColumns(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return getCommonColumnsMemoized(ctx, d, h)
}

func getCommonColumnsUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	config := GetConfig(d.Connection)
	if config.AccountId != nil {
		return databricksCommonColumnData{
			AccountId: *config.AccountId,
		}, nil
	}
	return nil, nil
}

// declare a wrapper hydrate function to call the memoized function
// - this is required when a memoized function is used for a column definition
func getAccountIdForConnection(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	commonColumnData, err := getCommonColumnsMemoized(ctx, d, h)
	if err != nil {
		return nil, err
	}

	return commonColumnData.(databricksCommonColumnData).AccountId, nil
}

type databricksCommonColumnData struct {
	AccountId string
}
