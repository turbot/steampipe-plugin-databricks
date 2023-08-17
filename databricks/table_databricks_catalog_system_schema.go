package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksCatalogSystemSchema(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_catalog_system_schema",
		Description: "Gets an array of system schemas for a metastore.",
		List: &plugin.ListConfig{
			ParentHydrate: listCatalogMetastores,
			Hydrate:       listCatalogSystemSchemas,
			KeyColumns:    plugin.OptionalColumns([]string{"metastore_id"}),
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "metastore_id",
				Description: "Unique identifier of parent metastore.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "schema",
				Description: "Name of the system schema.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "state",
				Description: "The current state of enablement for the system schema. An empty string means the system schema is available and ready for opt-in.",
				Type:        proto.ColumnType_STRING,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Schema"),
			},
		}),
	}
}

type systemSchemaInfo struct {
	catalog.SystemSchemaInfo
	MetastoreId string
}

//// LIST FUNCTION

func listCatalogSystemSchemas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	metastoreId := h.Item.(catalog.MetastoreInfo).MetastoreId

	// Return if input metastore_id is not matching with metastore_id of parent
	if d.EqualsQualString("metastore_id") != "" && d.EqualsQualString("metastore_id") != metastoreId {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_catalog_system_schema.listCatalogSystemSchemas", "connection_error", err)
		return nil, err
	}

	response, err := client.SystemSchemas.ListByMetastoreId(ctx, metastoreId)
	if err != nil {
		logger.Error("databricks_catalog_system_schema.listCatalogSystemSchemas", "api_error", err)
		return nil, err
	}

	for _, item := range response.Schemas {
		d.StreamListItem(ctx, systemSchemaInfo{item, metastoreId})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}
