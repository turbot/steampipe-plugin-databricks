package databricks

import (
	"context"
	"math"
	"strings"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func isNotFoundError(notFoundErrors []string) plugin.ErrorPredicate {
	return func(err error) bool {
		errMsg := err.Error()
		for _, msg := range notFoundErrors {
			if strings.Contains(errMsg, msg) {
				return true
			}
		}
		return false
	}
}

func convertTimestamp(_ context.Context, d *transform.TransformData) (interface{}, error) {

	epochTime := getEpochTime(d.Value)

	if epochTime != 0 {
		timeInSec := math.Floor(float64(epochTime) / 1000)
		unixTimestamp := time.Unix(int64(timeInSec), 0)
		timestampRFC3339Format := unixTimestamp.Format(time.RFC3339)
		return timestampRFC3339Format, nil
	}
	return nil, nil
}

func getEpochTime(item interface{}) int64 {
	switch item := item.(type) {
	case int64:
		return item
	case int:
		return int64(item)
	}
	return 0
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
