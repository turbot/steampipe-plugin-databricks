package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/databricks-sdk-go/service/serving"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksServingServingEndpoint(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_serving_serving_endpoint",
		Description: "List all serving endpoints.",
		List: &plugin.ListConfig{
			Hydrate: listServingServingEndpoints,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getServingServingEndpoint,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "System-generated ID of the endpoint.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "The name of the serving endpoint.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "creation_timestamp",
				Description: "Timestamp when the endpoint was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "creator",
				Description: "The email of the user who created the serving endpoint.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "last_updated_timestamp",
				Description: "The timestamp when the endpoint was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "permission_level",
				Description: "The permission level of the principal making the request.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getServingServingEndpoint,
			},

			// JSON fields
			{
				Name:        "config",
				Description: "The config that is currently being served by the endpoint.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "pending_config",
				Description: "The config that the endpoint is attempting to update to.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getServingServingEndpoint,
			},
			{
				Name:        "serving_endpoint_permissions",
				Description: "The permissions of the principal making the request.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getServingServingEndpointPermissions,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "state",
				Description: "Information corresponding to the state of the serving endpoint.",
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

func listServingServingEndpoints(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_serving_serving_endpoint.listServingServingEndpoints", "connection_error", err)
		return nil, err
	}

	endpoints, err := client.ServingEndpoints.ListAll(ctx)
	if err != nil {
		logger.Error("databricks_serving_serving_endpoint.listServingServingEndpoints", "api_error", err)
		return nil, err
	}

	for _, item := range endpoints {
		d.StreamListItem(ctx, item)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getServingServingEndpoint(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	var name string
	if h.Item != nil {
		name = getEndpoint(h.Item).Name
	} else {
		name = d.EqualsQualString("name")
	}

	// Return nil, if no input provided
	if name == "" {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_serving_serving_endpoint.getServingServingEndpoint", "connection_error", err)
		return nil, err
	}

	endpoint, err := client.ServingEndpoints.GetByName(ctx, name)
	if err != nil {
		logger.Error("databricks_serving_serving_endpoint.getServingServingEndpoint", "api_error", err)
		return nil, err
	}
	return endpoint, nil
}

func getServingServingEndpointPermissions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := getEndpoint(h.Item).Id

	request := iam.GetPermissionRequest{
		RequestObjectId:   id,
		RequestObjectType: "serving-endpoints",
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_serving_serving_endpoint.getServingServingEndpointPermissions", "connection_error", err)
		return nil, err
	}

	permissions, err := client.Permissions.Get(ctx, request)
	if err != nil {
		logger.Error("databricks_serving_serving_endpoint.getServingServingEndpointPermissions", "api_error", err)
		return nil, err
	}
	return permissions, nil
}

func getEndpoint(item interface{}) *endpointInfo {
	switch item := item.(type) {
	case serving.ServingEndpoint:
		return &endpointInfo{item.Id, item.Name}
	case serving.ServingEndpointDetailed:
		return &endpointInfo{item.Id, item.Name}
	}
	return nil
}

type endpointInfo struct {
	Id   string
	Name string
}
