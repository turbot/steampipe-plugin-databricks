package databricks

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksAccountCredential(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_account_credential",
		Description: "Gets all Databricks credential configurations associated with an account.",
		List: &plugin.ListConfig{
			Hydrate: listAccountCredentials,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"credentials_id", "credentials_name"}),
			Hydrate:    getAccountCredential,
		},
		Columns: []*plugin.Column{
			{
				Name:        "credentials_id",
				Description: "Databricks credential configuration ID.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "credentials_name",
				Description: "The human-readable name of the credential configuration object.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "creation_time",
				Description: "Time in epoch milliseconds when the credential was created.",
				Transform:   transform.FromGo().Transform(convertTimestamp),
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "account_id",
				Description: "The Databricks account ID that hosts the credential.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "aws_credentials",
				Description: "AWS credential configuration.",
				Type:        proto.ColumnType_JSON,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("CredentialsName"),
			},
		},
	}
}

//// LIST FUNCTION

func listAccountCredentials(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectDatabricksAccount(ctx, d)
	if err != nil {
		logger.Error("databricks_account_credential.listAccountCredentials", "connection_error", err)
		return nil, err
	}

	credentials, err := client.Credentials.List(ctx)
	if err != nil {
		logger.Error("databricks_account_credential.listAccountCredentials", "api_error", err)
		return nil, err
	}

	for _, item := range credentials {
		d.StreamListItem(ctx, item)

		// Context can be cancelled due to manual cancellation or if the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getAccountCredential(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectDatabricksAccount(ctx, d)
	if err != nil {
		logger.Error("databricks_account_credential.getAccountCredential", "connection_error", err)
		return nil, err
	}

	// Get by id if id provided as input
	if d.EqualsQuals["credentials_id"] != nil {
		id := d.EqualsQualString("credentials_id")

		credential, err := client.Credentials.GetByCredentialsId(ctx, id)
		if err != nil {
			logger.Error("databricks_account_credential.getAccountCredential", "api_error", err)
			return nil, err
		}
		return credential, nil
	}

	// Get by name if name provided as input
	if d.EqualsQuals["name"] != nil {
		name := d.EqualsQualString("name")

		credential, err := client.Credentials.GetByCredentialsName(ctx, name)
		if err != nil {
			logger.Error("databricks_account_credential.getAccountCredential", "api_error", err)
			return nil, err
		}
		return *credential, nil
	}

	return nil, nil
}
