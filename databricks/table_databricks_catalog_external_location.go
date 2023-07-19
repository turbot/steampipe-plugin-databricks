package databricks

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksCatalogExternalLocation(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_catalog_external_location",
		Description: "Gets an array of external locations from the metastore.",
		List: &plugin.ListConfig{
			Hydrate: listCatalogExternalLocations,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getCatalogExternalLocation,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "Human readable name that identifies the experiment.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "comment",
				Description: "User-provided free-form text description.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "Time at which this external location was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "created_by",
				Description: "The user who created this external location.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "credential_id",
				Description: "Unique ID of the location's storage credential.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "credential_name",
				Description: "Name of the storage credential used with this location.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "metastore_id",
				Description: "Unique identifier of metastore hosting the external location.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "owner",
				Description: "The user who owns this external location.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "read_only",
				Description: "Whether this external location is read-only.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "updated_at",
				Description: "Time at which this external location was last updated.",
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "updated_by",
				Description: "The user who last updated this external location.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "url",
				Description: "The Path URL of the external location.",
				Type:        proto.ColumnType_STRING,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

//// LIST FUNCTION

func listCatalogExternalLocations(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_catalog_external_location.listCatalogExternalLocations", "connection_error", err)
		return nil, err
	}

	response, err := client.ExternalLocations.ListAll(ctx)
	if err != nil {
		logger.Error("databricks_catalog_external_location.listCatalogExternalLocations", "api_error", err)
		return nil, err
	}

	for _, item := range response {
		d.StreamListItem(ctx, item)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getCatalogExternalLocation(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	name := d.EqualsQualString("name")

	// Return nil, if no input provided
	if name == "" {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_catalog_external_location.getCatalogExternalLocation", "connection_error", err)
		return nil, err
	}

	el, err := client.ExternalLocations.GetByName(ctx, name)
	if err != nil {
		logger.Error("databricks_catalog_external_location.getCatalogExternalLocation", "api_error", err)
		return nil, err
	}

	return *el, nil
}
