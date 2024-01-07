package mysql

import (
	"spun/internal/model"
	"strings"
)

// applySortFields
func applySortFields(querySelect string, queryValues map[string]interface{}, pagination *model.Pagination, queryDefaultSort string) (string, map[string]interface{}) {
	if pagination != nil {
		// Set query sort
		sortFields := len(pagination.SortFields)
		if sortFields > 0 {
			tmp := make([]string, sortFields)
			for k, sf := range pagination.SortFields {
				if strings.ToLower(sf.Direction) == "desc" {
					tmp[k] = sf.Field + " DESC"
				} else {
					tmp[k] = sf.Field + " ASC"
				}
			}
			querySelect += "ORDER BY " + strings.Join(tmp, ", ")
		} else {
			// Default sort
			querySelect += queryDefaultSort
		}

		// Set query limit and offset
		querySelect += " LIMIT :limit OFFSET :cursor"
		queryValues["cursor"] = pagination.Cursor
		queryValues["limit"] = pagination.Limit
	}
	return querySelect, queryValues
}
