package databricks

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksCatalogConnection(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_catalog_connection",
		Description: "Gets an array of connections for the workspace.",
		List: &plugin.ListConfig{
			Hydrate: listCatalogConnections,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("name"),
			ShouldIgnoreError: isNotFoundError([]string{"CONNECTION_DOES_NOT_EXIST"}),
			Hydrate:           getCatalogConnection,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "Name of the connection.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "connection_id",
				Description: "Unique identifier of the Connection.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "comment",
				Description: "User-provided free-form text description.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "connection_type",
				Description: "Type of the connection.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "The creation time of the connection.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "created_by",
				Description: "The user who created the connection.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "credential_type",
				Description: "Type of the credential.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "full_name",
				Description: "The full name of the connection.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "metastore_id",
				Description: "Unique identifier of parent metastore.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "owner",
				Description: "The user who owns the connection.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "read_only",
				Description: "Whether the connection is read-only.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "updated_at",
				Description: "The last time the connection was updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "updated_by",
				Description: "The user who last updated the connection.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "url",
				Description: "URL of the remote data source, extracted from options_kvpairs.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "options_kvpairs",
				Description: "A map of key-value properties attached to the securable.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "properties_kvpairs",
				Description: "An object containing map of key-value properties attached to the connection.",
				Type:        proto.ColumnType_JSON,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("FullName"),
			},
		}),
	}
}

//// LIST FUNCTION

func listCatalogConnections(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_catalog_connection.listCatalogConnections", "connection_error", err)
		return nil, err
	}

	connections, err := client.Connections.ListAll(ctx)
	if err != nil {
		logger.Error("databricks_catalog_connection.listCatalogConnections", "api_error", err)
		return nil, err
	}

	for _, item := range connections {
		d.StreamListItem(ctx, item)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getCatalogConnection(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	name := d.EqualsQualString("name")

	// Return nil, if no input provided
	if name == "" {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_catalog_connection.getCatalogConnection", "connection_error", err)
		return nil, err
	}

	connection, err := client.Connections.GetByNameArg(ctx, name)
	if err != nil {
		logger.Error("databricks_catalog_connection.getCatalogConnection", "api_error", err)
		return nil, err
	}

	return *connection, nil
}
