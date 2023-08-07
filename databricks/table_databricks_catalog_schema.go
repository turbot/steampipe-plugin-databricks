package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksCatalogSchema(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_catalog_schema",
		Description: "List schemas for a catalog in the metastore.",
		List: &plugin.ListConfig{
			ParentHydrate: listCatalogCatalogs,
			Hydrate:       listCatalogSchemas,
			KeyColumns:    plugin.OptionalColumns([]string{"catalog_name"}),
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("full_name"),
			ShouldIgnoreError: isNotFoundError([]string{"SCHEMA_DOES_NOT_EXIST"}),
			Hydrate:           getCatalogSchema,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "full_name",
				Description: "Full name of schema, in form of __catalog_name__.__schema_name__.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "Name of schema, relative to parent catalog.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "catalog_name",
				Description: "Name of parent catalog.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "catalog_type",
				Description: "The type of the parent catalog.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "comment",
				Description: "User-provided free-form text description.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "Time at which this schema was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "created_by",
				Description: "The user who created this schema.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "enable_auto_maintenance",
				Description: "Whether auto maintenance should be enabled for this object and objects under it.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "metastore_id",
				Description: "Unique identifier of parent metastore.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "owner",
				Description: "Owner of the schema.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "storage_location",
				Description: "Storage location for managed tables within schema.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "storage_root",
				Description: "Storage root URL for managed tables within schema.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "updated_at",
				Description: "Time at which this schema was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "updated_by",
				Description: "The user who last updated this schema.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "effective_auto_maintenance_flag",
				Description: "Effective auto maintenance flag of the schema.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "properties",
				Description: "A map of key-value properties attached to the securable.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "schema_permissions",
				Description: "Permissions of the schema.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getCatalogSchemaPermissions,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "schema_effective_permissions",
				Description: "Effective permissions of the schema.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getCatalogSchemaEffectivePermissions,
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

func listCatalogSchemas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	name := h.Item.(catalog.CatalogInfo).Name

	if d.EqualsQualString("catalog_name") != "" && d.EqualsQualString("catalog_name") != name {
		return nil, nil
	}

	request := catalog.ListSchemasRequest{
		CatalogName: name,
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_catalog_schema.listCatalogSchemas", "connection_error", err)
		return nil, err
	}

	schemas, err := client.Schemas.ListAll(ctx, request)
	if err != nil {
		logger.Error("databricks_catalog_schema.listCatalogSchemas", "api_error", err)
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

//// HYDRATE FUNCTIONS

func getCatalogSchema(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	name := d.EqualsQualString("full_name")

	// Return nil, if no input provided
	if name == "" {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_catalog_schema.getCatalogSchema", "connection_error", err)
		return nil, err
	}

	schema, err := client.Schemas.GetByFullName(ctx, name)
	if err != nil {
		logger.Error("databricks_catalog_schema.getCatalogSchema", "api_error", err)
		return nil, err
	}

	return *schema, nil
}

func getCatalogSchemaPermissions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	name := h.Item.(catalog.SchemaInfo).FullName

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_catalog_schema.getCatalogSchemaPermissions", "connection_error", err)
		return nil, err
	}

	permission, err := client.Grants.GetBySecurableTypeAndFullName(ctx, catalog.SecurableTypeSchema, name)
	if err != nil {
		logger.Error("databricks_catalog_schema.getCatalogSchemaPermissions", "api_error", err)
		return nil, err
	}
	return permission.PrivilegeAssignments, nil
}

func getCatalogSchemaEffectivePermissions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	name := h.Item.(catalog.SchemaInfo).FullName

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_catalog_schema.getCatalogSchemaEffectivePermissions", "connection_error", err)
		return nil, err
	}

	permission, err := client.Grants.GetEffectiveBySecurableTypeAndFullName(ctx, catalog.SecurableTypeSchema, name)
	if err != nil {
		logger.Error("databricks_catalog_schema.getCatalogSchemaEffectivePermissions", "api_error", err)
		return nil, err
	}
	return permission.PrivilegeAssignments, nil
}
