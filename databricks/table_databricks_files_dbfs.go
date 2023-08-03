package databricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/files"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDatabricksFilesDbfs(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "databricks_files_dbfs",
		Description: "List the contents of a directory, or details of the file.",
		List: &plugin.ListConfig{
			Hydrate:           listFilesDbfs,
			ShouldIgnoreError: isNotFoundError([]string{"RESOURCE_DOES_NOT_EXIST", "INVALID_PARAMETER_VALUE"}),
			KeyColumns:        plugin.AnyColumn([]string{"path", "path_prefix"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "path",
				Description: "The path of the file or directory.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "path_prefix",
				Description: "The path prefix of the file or directory.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("path_prefix"),
			},
			{
				Name:        "file_size",
				Description: "The length of the file in bytes or zero if the path is a directory.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "is_dir",
				Description: "True if the path is a directory.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "modification_time",
				Description: "Last modification time of given file/dir in milliseconds since Epoch.",
				Transform:   transform.FromGo().Transform(transform.UnixMsToTimestamp),
				Type:        proto.ColumnType_TIMESTAMP,
			},

			// JSON fields
			{
				Name:        "content",
				Description: "The content of the file.",
				Hydrate:     getFilesDbfsContent,
				Transform:   transform.FromValue(),
				Type:        proto.ColumnType_JSON,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Path"),
			},
		},
	}
}

//// LIST FUNCTION

func listFilesDbfs(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	path := d.EqualsQualString("path")
	pathPrefix := d.EqualsQualString("path_prefix")

	request := files.ListDbfsRequest{}

	if path != "" {
		request.Path = path
	} else if pathPrefix != "" {
		request.Path = pathPrefix
	} else {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_files_dbfs.listFilesDbfs", "connection_error", err)
		return nil, err
	}

	files, err := client.Dbfs.ListAll(ctx, request)
	if err != nil {
		logger.Error("databricks_files_dbfs.listFilesDbfs", "api_error", err)
		return nil, err
	}

	for _, item := range files {
		d.StreamListItem(ctx, item)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getFilesDbfsContent(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	file := h.Item.(files.FileInfo)

	if file.IsDir {
		return nil, nil
	}

	// Create client
	client, err := connectDatabricksWorkspace(ctx, d)
	if err != nil {
		logger.Error("databricks_files_dbfs.getFilesDbfsContent", "connection_error", err)
		return nil, err
	}

	request := files.ReadDbfsRequest{
		Path: file.Path,
	}

	content, err := client.Dbfs.Read(ctx, request)
	if err != nil {
		logger.Error("databricks_files_dbfs.getFilesDbfsContent", "api_error", err)
		return nil, err
	}

	return *content, nil
}
