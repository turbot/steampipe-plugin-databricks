package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksCatalog(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_catalog",
		Description: "Gets an array of catalogs in the metastore.",
		List: &plugin.ListConfig{
			Hydrate: listCatalogs,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("name"),
			ShouldIgnoreError: isNotFoundError([]string{"CATALOG_DOES_NOT_EXIST"}),
			Hydrate:           getCatalog,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "Name of the catalog.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "catalog_type",
				Description: "Type of the catalog.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "comment",
				Description: "User-provided free-form text description.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "connection_name",
				Description: "The name of the connection to an external data source.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "Time at which this catalog was created, in epoch milliseconds.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "created_by",
				Description: "Username of catalog creator.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "enable_auto_maintenance",
				Description: "Whether auto maintenance should be enabled for this object and objects under it.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "isolation_mode",
				Description: "Whether the current securable is accessible from all workspaces or a specific set of workspaces.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "metastore_id",
				Description: "Unique identifier of parent metastore.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "owner",
				Description: "Username of current owner of catalog.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "provider_name",
				Description: "The name of delta sharing provider.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "share_name",
				Description: "The name of the share under the share provider.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "storage_location",
				Description: "Storage Location URL (full path) for managed tables within catalog.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "storage_root",
				Description: "Storage root URL for managed tables within catalog.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "updated_at",
				Description: "Time at which this catalog was last updated, in epoch milliseconds.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "updated_by",
				Description: "Username of user who last modified catalog.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "catalog_permissions",
				Description: "Permissions for the catalog.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getCatalogPermissions,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "catalog_effective_permissions",
				Description: "Effective permissions for the catalog.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getCatalogEffectivePermissions,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "effective_auto_maintenance_flag",
				Description: "Effective auto maintenance flag.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "options",
				Description: "Catalog options - A map of key-value properties attached to the securable.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "properties",
				Description: "Catalog properties - A map of key-value properties attached to the securable.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "workspace_bindings",
				Description: "Array of workspace bindings.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getCatalogWorkspaceBindings,
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

func listCatalogs(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_catalog.listCatalogs", "connection_error", err)
		return nil, err
	}

	catalogs, err := client.Catalogs.ListAll(ctx)
	if err != nil {
		logger.Error("databricks_catalog.listCatalogs", "api_error", err)
		return nil, err
	}

	for _, item := range catalogs {
		d.StreamListItem(ctx, item)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getCatalog(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	name := d.EqualsQualString("name")

	// Return nil, if no input provided
	if name == "" {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_catalog.getCatalog", "connection_error", err)
		return nil, err
	}

	catalog, err := client.Catalogs.GetByName(ctx, name)
	if err != nil {
		logger.Error("databricks_catalog.getCatalog", "api_error", err)
		return nil, err
	}

	return *catalog, nil
}

func getCatalogWorkspaceBindings(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	name := h.Item.(catalog.CatalogInfo).Name

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_catalog.getCatalogWorkspaceBindings", "connection_error", err)
		return nil, err
	}

	bindings, err := client.WorkspaceBindings.GetByName(ctx, name)
	if err != nil {
		logger.Error("databricks_catalog.getCatalogWorkspaceBindings", "api_error", err)
		return nil, err
	}

	return *bindings, nil
}

func getCatalogPermissions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	name := h.Item.(catalog.CatalogInfo).Name

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_catalog.getCatalogPermissions", "connection_error", err)
		return nil, err
	}

	permission, err := client.Grants.GetBySecurableTypeAndFullName(ctx, catalog.SecurableTypeCatalog, name)
	if err != nil {
		logger.Error("databricks_catalog.getCatalogPermissions", "api_error", err)
		return nil, err
	}
	return permission.PrivilegeAssignments, nil
}

func getCatalogEffectivePermissions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	name := h.Item.(catalog.CatalogInfo).Name

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_catalog.getCatalogEffectivePermissions", "connection_error", err)
		return nil, err
	}

	permission, err := client.Grants.GetEffectiveBySecurableTypeAndFullName(ctx, catalog.SecurableTypeCatalog, name)
	if err != nil {
		logger.Error("databricks_catalog.getCatalogEffectivePermissions", "api_error", err)
		return nil, err
	}
	return permission.PrivilegeAssignments, nil
}
