package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksCatalogTable(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_catalog_table",
		Description: "Gets an array of the available tables.",
		List: &plugin.ListConfig{
			Hydrate:    listCatalogTables,
			KeyColumns: plugin.OptionalColumns([]string{"catalog_name", "schema_name"}),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("full_name"),
			Hydrate:    getCatalogTable,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "table_id",
				Description: "Name of table, relative to parent schema.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "full_name",
				Description: "Full name of table, in form of __catalog_name__.__schema_name__.__table_name__.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "catalog_name",
				Description: "Name of parent catalog.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "schema_name",
				Description: "Name of parent schema relative to its parent catalog.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "comment",
				Description: "User-provided free-form text description.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "Time at which this table was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "created_by",
				Description: "The user who created the table.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "data_access_configuration_id",
				Description: "Unique ID of the Data Access Configuration to use with the table data.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "data_source_format",
				Description: "Data source format.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "deleted_at",
				Description: "Time at which this table was deleted.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "metastore_id",
				Description: "Unique identifier of parent metastore.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "Name of table, relative to parent schema.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "owner",
				Description: "Username of current owner of table.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "sql_path",
				Description: "List of schemes whose objects can be referenced without qualification.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "storage_credential_name",
				Description: "Name of the storage credential, when a storage credential is configured for use with this table.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "storage_location",
				Description: "Storage root URL for table.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "table_type",
				Description: "The type of the table.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "updated_at",
				Description: "Time at which this table was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "updated_by",
				Description: "The user who last updated the table.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "view_definition",
				Description: "The SQL text defining the view.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON Columns
			{
				Name:        "columns",
				Description: "The array of __ColumnInfo__ definitions of the table's columns.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "delta_runtime_properties_kvpairs",
				Description: "Information pertaining to current state of the delta table.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "properties",
				Description: "A map of key-value properties attached to the securable.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "row_filter",
				Description: "The row filter associated with the table.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "table_constraints",
				Description: "Table constraints associated with the table.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "view_dependencies",
				Description: "View dependencies associated with the table.",
				Type:        proto.ColumnType_JSON,
			},
		}),
	}
}

//// LIST FUNCTION

func listCatalogTables(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Limiting the results
	maxLimit := 1000
	if d.QueryContext.Limit != nil {
		limit := int(*d.QueryContext.Limit)
		if limit < maxLimit {
			maxLimit = limit
		}
	}

	request := catalog.ListTablesRequest{
		MaxResults:           maxLimit,
		IncludeDeltaMetadata: true,
	}

	if d.EqualsQualString("catalog_name") != "" {
		request.CatalogName = d.EqualsQualString("catalog_name")
	}
	if d.EqualsQualString("schema_name") != "" {
		request.SchemaName = d.EqualsQualString("schema_name")
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_catalog_table.listCatalogTables", "connection_error", err)
		return nil, err
	}

	for {
		response, err := client.Tables.Impl().List(ctx, request)
		if err != nil {
			logger.Error("databricks_catalog_table.listCatalogTables", "api_error", err)
			return nil, err
		}

		for _, item := range response.Tables {
			d.StreamListItem(ctx, item)

			// Context can be cancelled due to manual cancellation or if the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		if response.NextPageToken == "" {
			return nil, nil
		}
		request.PageToken = response.NextPageToken
	}
}

//// HYDRATE FUNCTIONS

func getCatalogTable(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	name := d.EqualsQualString("full_name")

	// Return nil, if no input provided
	if name == "" {
		return nil, nil
	}

	request := catalog.GetTableRequest{
		FullName:             name,
		IncludeDeltaMetadata: true,
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_catalog_table.getCatalogTable", "connection_error", err)
		return nil, err
	}

	table, err := client.Tables.Get(ctx, request)
	if err != nil {
		logger.Error("databricks_catalog_table.getCatalogTable", "api_error", err)
		return nil, err
	}
	return *table, nil
}
