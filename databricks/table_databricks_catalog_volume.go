package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksCatalogVolume(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_catalog_volume",
		Description: "Gets an array of the available volumes.",
		List: &plugin.ListConfig{
			Hydrate:           listCatalogVolumes,
			ShouldIgnoreError: isNotFoundError([]string{"CATALOG_DOES_NOT_EXIST", "SCHEMA_DOES_NOT_EXIST"}),
			KeyColumns:        plugin.AllColumns([]string{"catalog_name", "schema_name"}),
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("full_name"),
			ShouldIgnoreError: isNotFoundError([]string{"VOLUME_DOES_NOT_EXIST", "CATALOG_DOES_NOT_EXIST", "SCHEMA_DOES_NOT_EXIST"}),
			Hydrate:           getCatalogVolume,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "volume_id",
				Description: "The unique identifier of the volume.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "full_name",
				Description: "The three-level (fully qualified) name of the volume.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "catalog_name",
				Description: "The name of the catalog where the schema and the volume are.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "schema_name",
				Description: "The name of the schema where the volume is.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "comment",
				Description: "The comment attached to the volume.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "Time at which this volume was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "created_by",
				Description: "The user who created the volume.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "metastore_id",
				Description: "Unique identifier of parent metastore.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "Name of volume.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "owner",
				Description: "The identifier of the user who owns the volume.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "storage_location",
				Description: "The storage location on the cloud.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "updated_at",
				Description: "Time at which this volume was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "updated_by",
				Description: "The user who last updated the volume.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "volume_type",
				Description: "The type of the volume.",
				Type:        proto.ColumnType_STRING,
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

func listCatalogVolumes(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	catalogName := d.EqualsQualString("catalog_name")
	schemaName := d.EqualsQualString("schema_name")

	// Return nil, if no input provided
	if catalogName == "" || schemaName == "" {
		return nil, nil
	}

	request := catalog.ListVolumesRequest{
		CatalogName: catalogName,
		SchemaName:  schemaName,
	}

	// Create client
	client, err := getWorkspaceClient(ctx, d)
	if err != nil {
		logger.Error("databricks_catalog_volume.listCatalogVolumes", "connection_error", err)
		return nil, err
	}

	volumes, err := client.Volumes.ListAll(ctx, request)
	if err != nil {
		logger.Error("databricks_catalog_volume.listCatalogVolumes", "api_error", err)
		return nil, err
	}

	for _, item := range volumes {
		d.StreamListItem(ctx, item)

		// Context can be cancelled due to manual cancellation or if the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getCatalogVolume(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	name := d.EqualsQualString("full_name")

	// Return nil, if no input provided
	if name == "" {
		return nil, nil
	}

	request := catalog.ReadVolumeRequest{
		FullNameArg: name,
	}

	// Create client
	client, err := getWorkspaceClient(ctx, d)
	if err != nil {
		logger.Error("databricks_catalog_volume.getCatalogVolume", "connection_error", err)
		return nil, err
	}

	volume, err := client.Volumes.Read(ctx, request)
	if err != nil {
		logger.Error("databricks_catalog_volume.getCatalogVolume", "api_error", err)
		return nil, err
	}
	return volume, nil
}
