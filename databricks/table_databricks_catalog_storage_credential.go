package databricks

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksCatalogStorageCredential(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_catalog_storage_credential",
		Description: "Gets an array of storage credentials.",
		List: &plugin.ListConfig{
			Hydrate: listCatalogStorageCredentials,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getCatalogStorageCredential,
		},
		Columns: databricksAccountColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "Unique identifier of the credential.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "The credential name.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "comment",
				Description: "Comment associated with the credential.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "Time at which this credential was created.",
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "created_by",
				Description: "The user who created this credential.",
				Type:        proto.ColumnType_STRING,
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
				Name:        "read_only",
				Description: "Whether the storage credential is only usable for read operations.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "updated_at",
				Description: "Time at which this credential was last updated.",
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "updated_by",
				Description: "The user who last updated this credential.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "used_for_managed_storage",
				Description: "Whether this credential is the current metastore's root storage credential.",
				Type:        proto.ColumnType_BOOL,
			},

			// JSON fields
			{
				Name:        "aws_iam_role",
				Description: "The AWS IAM role configuration.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "azure_managed_identity",
				Description: "The Azure managed identity configuration.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "azure_service_principal",
				Description: "The Azure service principal configuration.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "databricks_gcp_service_account",
				Description: "The Databricks GCP service account configuration.",
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

func listCatalogStorageCredentials(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_catalog_storage_credential.listCatalogStorageCredentials", "connection_error", err)
		return nil, err
	}

	creds, err := client.StorageCredentials.ListAll(ctx)
	if err != nil {
		logger.Error("databricks_catalog_storage_credential.listCatalogStorageCredentials", "api_error", err)
		return nil, err
	}

	for _, item := range creds {
		d.StreamListItem(ctx, item)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getCatalogStorageCredential(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	name := d.EqualsQualString("name")

	// Return nil, if no input provided
	if name == "" {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_catalog_storage_credential.getCatalogStorageCredential", "connection_error", err)
		return nil, err
	}

	cred, err := client.StorageCredentials.GetByName(ctx, name)
	if err != nil {
		logger.Error("databricks_catalog_storage_credential.getCatalogStorageCredential", "api_error", err)
		return nil, err
	}

	return cred, nil
}
