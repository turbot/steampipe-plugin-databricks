package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksSharingProvider(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_sharing_provider",
		Description: "Gets an array of available authentication providers.",
		List: &plugin.ListConfig{
			Hydrate:           listSharingProviders,
			ShouldIgnoreError: isNotFoundError([]string{"INVALID_PARAMETER_VALUE"}),
			KeyColumns:        plugin.OptionalColumns([]string{"data_provider_global_metastore_id"}),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"name"}),
			Hydrate:    getSharingProvider,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "Name of the provider.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "data_provider_global_metastore_id",
				Description: "The global UC metastore id of the data provider.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "authentication_type",
				Description: "The delta sharing authentication type.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "cloud",
				Description: "Cloud vendor of the provider's UC metastore.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "comment",
				Description: "Description about the provider.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "Timestamp when the provider was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "created_by",
				Description: "User who created the provider.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "metastore_id",
				Description: "UUID of the provider's UC metastore.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "owner",
				Description: "User who owns the provider.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "recipient_profile_str",
				Description: "The recipient profile as a string.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "region",
				Description: "Cloud region of the provider's UC metastore.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "updated_at",
				Description: "Timestamp when the provider was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "updated_by",
				Description: "User who last modified the provider.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "recipient_profile",
				Description: "The recipient profile description.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "shares",
				Description: "An array of provider's shares.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getSharingProviderShares,
				Transform:   transform.FromValue(),
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		}),
	}
}

//// LIST FUNCTION

func listSharingProviders(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	request := sharing.ListProvidersRequest{}
	if d.EqualsQualString("data_provider_global_metastore_id") != "" {
		request.DataProviderGlobalMetastoreId = d.EqualsQualString("data_provider_global_metastore_id")
	}

	// Create client
	client, err := getWorkspaceClient(ctx, d)
	if err != nil {
		logger.Error("databricks_sharing_provider.listSharingProviders", "connection_error", err)
		return nil, err
	}

	providers, err := client.Providers.ListAll(ctx, request)
	if err != nil {
		logger.Error("databricks_sharing_provider.listSharingProviders", "api_error", err)
		return nil, err
	}

	for _, item := range providers {
		d.StreamListItem(ctx, item)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getSharingProvider(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	name := d.EqualsQualString("name")

	// Return nil, if no input provided
	if name == "" {
		return nil, nil
	}

	// Create client
	client, err := getWorkspaceClient(ctx, d)
	if err != nil {
		logger.Error("databricks_sharing_provider.getSharingProvider", "connection_error", err)
		return nil, err
	}

	provider, err := client.Providers.GetByName(ctx, name)
	if err != nil {
		logger.Error("databricks_sharing_provider.getSharingProvider", "api_error", err)
		return nil, err
	}
	return provider, nil
}

func getSharingProviderShares(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	name := h.Item.(sharing.ProviderInfo).Name

	// Create client
	client, err := getWorkspaceClient(ctx, d)
	if err != nil {
		logger.Error("databricks_sharing_provider.getSharingProviderShares", "connection_error", err)
		return nil, err
	}

	shares, err := client.Providers.ListSharesByName(ctx, name)
	if err != nil {
		logger.Error("databricks_sharing_provider.getSharingProviderShares", "api_error", err)
		return nil, err
	}
	return shares.Shares, nil
}
