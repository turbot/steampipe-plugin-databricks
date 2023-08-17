package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksCatalogFunction(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_catalog_function",
		Description: "List functions within the specified parent catalog and schema.",
		List: &plugin.ListConfig{
			Hydrate:           listCatalogFunctions,
			ShouldIgnoreError: isNotFoundError([]string{"CATALOG_DOES_NOT_EXIST", "SCHEMA_DOES_NOT_EXIST"}),
			KeyColumns:        plugin.AllColumns([]string{"catalog_name", "schema_name"}),
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("full_name"),
			ShouldIgnoreError: isNotFoundError([]string{"FUNCTION_DOES_NOT_EXIST", "CATALOG_DOES_NOT_EXIST", "SCHEMA_DOES_NOT_EXIST"}),
			Hydrate:           getCatalogFunction,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "function_id",
				Description: "Id of Function, relative to parent schema.",
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
				Name:        "name",
				Description: "Name of function, relative to parent schema.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "comment",
				Description: "User-provided free-form text description.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "Time at which this function was created.",
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "created_by",
				Description: "The user who created this function.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "data_type",
				Description: "Scalar function return data type.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "external_language",
				Description: "External function language.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "external_name",
				Description: "External function name.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "full_data_type",
				Description: "Pretty printed function data type.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "full_name",
				Description: "Full name of function, in form of __catalog_name__.__schema_name__.__function__name__.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "is_deterministic",
				Description: "Whether the function is deterministic.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "is_null_call",
				Description: "Whether the function is a null call.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "metastore_id",
				Description: "Unique identifier of parent metastore.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "owner",
				Description: "Owner of the function.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "parameter_style",
				Description: "Parameter style of the function.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "routine_body",
				Description: "Routine body of the function.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "routine_definition",
				Description: "Routine definition of the function.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "security_type",
				Description: "Security type of the function.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "specific_name",
				Description: "Specific name of the function.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "sql_data_access",
				Description: "SQL data access of the function.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "sql_path",
				Description: "List of schemes whose objects can be referenced without qualification.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "updated_at",
				Description: "Time at which this function was last updated.",
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "updated_by",
				Description: "The user who last updated this function.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "function_permissions",
				Description: "Function permissions.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getCatalogFunctionPermissions,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "function_effective_permissions",
				Description: "Function effective permissions.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getCatalogFunctionEffectivePermissions,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "input_params",
				Description: "The array of __FunctionParameterInfo__ definitions of the function's parameters.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "properties",
				Description: "A map of key-value properties attached to the securable.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "return_params",
				Description: "Table function return parameters.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "routine_dependencies",
				Description: "Routine dependencies of the function.",
				Type:        proto.ColumnType_JSON,
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

func listCatalogFunctions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	if d.EqualsQualString("catalog_name") == "" || d.EqualsQualString("schema_name") == "" {
		return nil, nil
	}

	request := catalog.ListFunctionsRequest{
		SchemaName:  d.EqualsQualString("schema_name"),
		CatalogName: d.EqualsQualString("catalog_name"),
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_catalog_function.listCatalogFunctions", "connection_error", err)
		return nil, err
	}

	response, err := client.Functions.ListAll(ctx, request)
	if err != nil {
		logger.Error("databricks_catalog_function.listCatalogFunctions", "api_error", err)
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

func getCatalogFunction(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	name := d.EqualsQualString("full_name")

	// Return nil, if no input provided
	if name == "" {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_catalog_function.getCatalogFunction", "connection_error", err)
		return nil, err
	}

	function, err := client.Functions.GetByName(ctx, name)
	if err != nil {
		logger.Error("databricks_catalog_function.getCatalogFunction", "api_error", err)
		return nil, err
	}

	return *function, nil
}

func getCatalogFunctionPermissions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	name := h.Item.(catalog.FunctionInfo).FullName

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_catalog_function.getCatalogFunctionPermissions", "connection_error", err)
		return nil, err
	}

	permission, err := client.Grants.GetBySecurableTypeAndFullName(ctx, catalog.SecurableTypeFunction, name)
	if err != nil {
		logger.Error("databricks_catalog_function.getCatalogFunctionPermissions", "api_error", err)
		return nil, err
	}
	return permission.PrivilegeAssignments, nil
}

func getCatalogFunctionEffectivePermissions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	name := h.Item.(catalog.FunctionInfo).FullName

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_catalog_function.getCatalogFunctionEffectivePermissions", "connection_error", err)
		return nil, err
	}

	permission, err := client.Grants.GetEffectiveBySecurableTypeAndFullName(ctx, catalog.SecurableTypeFunction, name)
	if err != nil {
		logger.Error("databricks_catalog_function.getCatalogFunctionEffectivePermissions", "api_error", err)
		return nil, err
	}
	return permission.PrivilegeAssignments, nil
}
