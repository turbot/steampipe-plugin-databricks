package databricks

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksAccountEncryptionKey(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_account_encryption_key",
		Description: "Gets all customer-managed key configuration objects for an account.",
		List: &plugin.ListConfig{
			Hydrate: listAccountEncryptionKeys,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("customer_managed_key_id"),
			Hydrate:    getAccountEncryptionKey,
		},
		Columns: []*plugin.Column{
			{
				Name:        "customer_managed_key_id",
				Description: "ID of the encryption key configuration object.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "account_id",
				Description: "The Databricks account ID that holds the customer-managed key.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "creation_time",
				Description: "Time in epoch milliseconds when the customer key was created.",
				Type:        proto.ColumnType_BOOL,
			},

			// JSON fields
			{
				Name:        "aws_key_info",
				Description: "AWS KMS key information.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "gcp_key_info",
				Description: "GCP KMS key information.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "use_cases",
				Description: "The cases that the key can be used for.",
				Type:        proto.ColumnType_JSON,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("CustomerManagedKeyId"),
			},
		},
	}
}

//// LIST FUNCTION

func listAccountEncryptionKeys(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectDatabricksAccount(ctx, d)
	if err != nil {
		logger.Error("databricks_account_encryption_key.listAccountEncryptionKeys", "connection_error", err)
		return nil, err
	}

	keys, err := client.EncryptionKeys.List(ctx)
	if err != nil {
		logger.Error("databricks_account_encryption_key.listAccountEncryptionKeys", "api_error", err)
		return nil, err
	}

	for _, item := range keys {
		d.StreamListItem(ctx, &item)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getAccountEncryptionKey(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("id")

	// Return nil, if no input provided
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksAccount(ctx, d)
	if err != nil {
		logger.Error("databricks_account_encryption_key.getAccountEncryptionKey", "connection_error", err)
		return nil, err
	}

	key, err := client.EncryptionKeys.GetByCustomerManagedKeyId(ctx, id)
	if err != nil {
		logger.Error("databricks_account_encryption_key.getAccountEncryptionKey", "api_error", err)
		return nil, err
	}

	return key, nil
}
