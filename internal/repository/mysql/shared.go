package mysql

import (
	"errors"
	"fmt"
	"goproj/internal/model"
	"strings"

	"github.com/go-sql-driver/mysql"
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

// parseError convert default error to recognizeable error
func parseError(err error) error {
	if err != nil {
		// Convert the error to the MySQL error type
		// This type assertion is safe as long as you're using the MySQL driver
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			// Check if the error code is 1054, which is for unknown column
			if mysqlErr.Number == 1054 {
				if strings.Contains(strings.ToLower(mysqlErr.Message), "order clause") {
					return errors.New("common.error.server.repository.sort")
				}
			}
		} else {
			// Handle other types of errors
			fmt.Println("An error occurred:", err)
		}
		return errors.New("common.error.server.repository.default")
	}
	return nil
}
