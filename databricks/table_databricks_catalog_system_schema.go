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
			Hydrate: listCatalogSystemSchemas,
			// KeyColumns: plugin.OptionalColumns([]string{"metastore_id"}),
		},
		Columns: databricksAccountColumns([]*plugin.Column{
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

//// LIST FUNCTION

func listCatalogSystemSchemas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	request := catalog.ListSystemSchemasRequest{}
	// if d.EqualsQualString("metastore_id") != "" {
	// 	request.MetastoreId = d.EqualsQualString("metastore_id")
	// }

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_catalog_system_schema.listCatalogSystemSchemas", "connection_error", err)
		return nil, err
	}

	schemas, err := client.SystemSchemas.ListAll(ctx, request)
	if err != nil {
		logger.Error("databricks_catalog_system_schema.listCatalogSystemSchemas", "api_error", err)
		return nil, err
	}

	for _, item := range schemas {
		d.StreamListItem(ctx, item)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}
