package databricks

import (
	"strconv"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func isNotFoundError(notFoundErrors []string) plugin.ErrorPredicate {
	return func(err error) bool {
		errMsg := err.(*apierr.APIError)
		for _, msg := range notFoundErrors {
			if strings.Contains(errMsg.ErrorCode, msg) {
				return true
			} else if strings.Contains(strconv.Itoa(errMsg.StatusCode), msg) {
				return true
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
							filters += filterQualItem.PropertyPath + GcpFilterOperatorMap[qual.Operator] + value.GetStringValue()
						}
						// case "boolean":
						// 	boolValue := value.GetBoolValue()
						// 	switch qual.Operator {
						// 	case "<>":
						// 		filters += filterQualItem.PropertyPath + GcpFilterOperatorMap[qual.Operator] + !boolValue
						// 		filters = append(filters, fmt.Sprintf("(%s = %t)", filterQualItem.PropertyPath, !boolValue))
						// 	case "=":
						// 		filters = append(filters, fmt.Sprintf("(%s = %t)", filterQualItem.PropertyPath, boolValue))
						// 	}
						// }
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

var GcpFilterOperatorMap = map[string]string{
	"=":  " eq ",
	"<>": " ne ",
	"!=": " ne ",
}
