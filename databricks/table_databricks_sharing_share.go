package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksSharingShare(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_sharing_share",
		Description: "Lists all data object shares from the metastore.",
		List: &plugin.ListConfig{
			Hydrate: listSharingShares,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("name"),
			ShouldIgnoreError: isNotFoundError([]string{"SHARE_DOES_NOT_EXIST"}),
			Hydrate:           getSharingShare,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "Name of the share.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "comment",
				Description: "User-provided free-form text description.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "Timestamp when the share was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "created_by",
				Description: "User who created the share.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "owner",
				Description: "Username of current owner of share.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "updated_at",
				Description: "Timestamp when the share was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "updated_by",
				Description: "User who last modified the share.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "objects",
				Description: "A list of shared data objects within the share.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "permissions",
				Description: "A list of share permissions.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getSharingSharePermissions,
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

func listSharingShares(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_sharing_share.listSharingShares", "connection_error", err)
		return nil, err
	}

	shares, err := client.Shares.ListAll(ctx)
	if err != nil {
		logger.Error("databricks_sharing_share.listSharingShares", "api_error", err)
		return nil, err
	}

	for _, item := range shares {
		d.StreamListItem(ctx, item)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getSharingShare(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	name := d.EqualsQualString("name")

	// Return nil, if no input provided
	if name == "" {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_sharing_share.getSharingShare", "connection_error", err)
		return nil, err
	}

	request := sharing.GetShareRequest{
		Name:              name,
		IncludeSharedData: true,
	}
	share, err := client.Shares.Get(ctx, request)
	if err != nil {
		logger.Error("databricks_sharing_share.getSharingShare", "api_error", err)
		return nil, err
	}
	return share, nil
}

func getSharingSharePermissions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	name := h.Item.(sharing.ShareInfo).Name

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_sharing_share.getSharingSharePermissions", "connection_error", err)
		return nil, err
	}

	permissions, err := client.Shares.SharePermissionsByName(ctx, name)
	if err != nil {
		logger.Error("databricks_sharing_share.getSharingSharePermissions", "api_error", err)
		return nil, err
	}
	return permissions.PrivilegeAssignments, nil
}
