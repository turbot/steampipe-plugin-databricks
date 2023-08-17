package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksSharingRecipient(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_sharing_recipient",
		Description: "Gets an array of all share recipients within the current metastore.",
		List: &plugin.ListConfig{
			Hydrate:           listSharingRecipients,
			ShouldIgnoreError: isNotFoundError([]string{"INVALID_PARAMETER_VALUE"}),
			KeyColumns:        plugin.OptionalColumns([]string{"data_recipient_global_metastore_id"}),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"name"}),
			Hydrate:    getSharingRecipient,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "Name of the recipient.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "data_recipient_global_metastore_id",
				Description: "The global Unity Catalog metastore id of the data recipient.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "activated",
				Description: "A boolean status field showing whether the Recipient's activation URL has been exercised or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "activation_url",
				Description: "Full activation url to retrieve the access token.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "authentication_type",
				Description: "The delta sharing authentication type.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "cloud",
				Description: "Cloud vendor of the recipient's Unity Catalog metastore.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "comment",
				Description: "Description about the recipient.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "Timestamp when the recipient was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "created_by",
				Description: "User who created the recipient.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "metastore_id",
				Description: "UUID of the recipient's Unity Catalog metastore.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "owner",
				Description: "Username of the recipient owner.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "recipient_profile_str",
				Description: "The recipient profile as a string.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "region",
				Description: "Cloud region of the recipient's Unity Catalog metastore.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "sharing_code",
				Description: "The one-time sharing code provided by the data recipient.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "updated_at",
				Description: "Timestamp when the recipient was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "updated_by",
				Description: "User who last modified the recipient.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "ip_access_list",
				Description: "IP Access list",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "properties_kvpairs",
				Description: "Recipient properties as map of string key-value pairs.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "permissions",
				Description: "An array of recipient's shares.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getSharingRecipientPermissions,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "tokens",
				Description: "This field is only present when the __authentication_type__ is **TOKEN**.",
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

func listSharingRecipients(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	request := sharing.ListRecipientsRequest{}
	if d.EqualsQualString("data_recipient_global_metastore_id") != "" {
		request.DataRecipientGlobalMetastoreId = d.EqualsQualString("data_recipient_global_metastore_id")
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_sharing_recipient.listSharingRecipients", "connection_error", err)
		return nil, err
	}

	recipients, err := client.Recipients.ListAll(ctx, request)
	if err != nil {
		logger.Error("databricks_sharing_recipient.listSharingRecipients", "api_error", err)
		return nil, err
	}

	for _, item := range recipients {
		d.StreamListItem(ctx, item)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getSharingRecipient(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	name := d.EqualsQualString("name")

	// Return nil, if no input provided
	if name == "" {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_sharing_recipient.getSharingRecipient", "connection_error", err)
		return nil, err
	}

	recipient, err := client.Recipients.GetByName(ctx, name)
	if err != nil {
		logger.Error("databricks_sharing_recipient.getSharingRecipient", "api_error", err)
		return nil, err
	}
	return recipient, nil
}

func getSharingRecipientPermissions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	name := h.Item.(sharing.RecipientInfo).Name

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_sharing_recipient.getSharingRecipientPermissions", "connection_error", err)
		return nil, err
	}

	permissions, err := client.Recipients.SharePermissionsByName(ctx, name)
	if err != nil {
		logger.Error("databricks_sharing_recipient.getSharingRecipientPermissions", "api_error", err)
		return nil, err
	}
	return permissions.PermissionsOut, nil
}
