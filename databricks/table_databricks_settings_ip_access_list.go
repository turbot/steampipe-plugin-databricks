package databricks

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksSettingsIpAccessList(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_settings_ip_access_list",
		Description: "Gets all IP access lists for the specified workspace.",
		List: &plugin.ListConfig{
			Hydrate: listSettingsIpAccessLists,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("list_id"),
			Hydrate:    getSettingsIpAccessList,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "list_id",
				Description: "Universally unique identifier (UUID) of the IP access list.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "label",
				Description: "Label for the IP access list.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "address_count",
				Description: "Total number of IP or CIDR values.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "created_at",
				Description: "Time at which the IP access list was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "created_by",
				Description: "User ID of the user who created this list.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "enabled",
				Description: "Whether this IP access list is enabled.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "list_type",
				Description: "The list type.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "updated_at",
				Description: "Time at which the IP access list was updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "updated_by",
				Description: "User ID of the user who updated this list.",
				Type:        proto.ColumnType_INT,
			},

			// JSON fields
			{
				Name:        "ip_addresses",
				Description: "Array of IP addresses or CIDR values to be added to the IP access list.",
				Type:        proto.ColumnType_JSON,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Label"),
			},
		}),
	}
}

//// LIST FUNCTION

func listSettingsIpAccessLists(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := getWorkspaceClient(ctx, d)
	if err != nil {
		logger.Error("databricks_settings_ip_access_list.listSettingsIpAccessLists", "connection_error", err)
		return nil, err
	}

	lists, err := client.IpAccessLists.ListAll(ctx)
	if err != nil {
		logger.Error("databricks_settings_ip_access_list.listSettingsIpAccessLists", "api_error", err)
		return nil, err
	}

	for _, item := range lists {
		d.StreamListItem(ctx, item)

		// Context can be cancelled due to manual cancellation or if the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getSettingsIpAccessList(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("list_id")

	// Return nil, if no input provided
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := getWorkspaceClient(ctx, d)
	if err != nil {
		logger.Error("databricks_settings_ip_access_list.getSettingsIpAccessList", "connection_error", err)
		return nil, err
	}

	list, err := client.IpAccessLists.GetByIpAccessListId(ctx, id)
	if err != nil {
		logger.Error("databricks_settings_ip_access_list.getSettingsIpAccessList", "api_error", err)
		return nil, err
	}
	return *list.IpAccessList, nil
}
