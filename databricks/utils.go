package databricks

import (
	"context"
	"strconv"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func isNotFoundError(notFoundErrors []string) plugin.ErrorPredicate {
	return func(err error) bool {
		switch err := err.(type) {
		case *apierr.APIError:
			for _, msg := range notFoundErrors {
				if strings.Contains(err.ErrorCode, msg) {
					return true
				} else if strings.Contains(strconv.Itoa(err.StatusCode), msg) {
					return true
				}
			}
		}
		return false
	}
}

func shouldRetryError(retryErrors []string) plugin.ErrorPredicateWithContext {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
		switch err := err.(type) {
		case *apierr.APIError:
			for _, msg := range retryErrors {
				if strings.Contains(err.ErrorCode, msg) {
					plugin.Logger(ctx).Error("databricks_errors.shouldRetryError", "rate_limit_error", err)
					return true
				} else if strings.Contains(strconv.Itoa(err.StatusCode), msg) {
					plugin.Logger(ctx).Error("databricks_errors.shouldRetryError", "rate_limit_error", err)
					return true
				}
			}
		}
		return false
	}
}

func buildQueryFilterFromQuals(filterQuals []filterQualMap, equalQuals plugin.KeyColumnQualMap) string {

	filters := ""

	for _, filterQualItem := range filterQuals {
		filterQual := equalQuals[filterQualItem.ColumnName]
		if filterQual == nil {
			continue
		}

		// Check only if filter qual map matches with optional column name
		if filterQual.Name == filterQualItem.ColumnName {
			if filterQual.Quals == nil {
				continue
			}

			for _, qual := range filterQual.Quals {
				if qual.Value != nil {
					value := qual.Value
					switch filterQualItem.Type {
					case "string":
						switch qual.Operator {
						case "=", "<>":
							if filters != "" {
								filters += " and "
							}
							filters += filterQualItem.PropertyPath + DatabricksCompositeFilterOperatorMap[qual.Operator] + value.GetStringValue()
						}
					}
				}
			}
		}
	}

	return filters
}

type filterQualMap struct {
	ColumnName   string
	PropertyPath string
	Type         string
}

var DatabricksCompositeFilterOperatorMap = map[string]string{
	"=":  " eq ",
	"<>": " ne ",
	"!=": " ne ",
}
