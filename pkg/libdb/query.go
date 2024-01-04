package libdb

import (
	"fmt"
	"strings"
)

// PrepareUpdateQuery return set query and values map to be used on update
func PrepareUpdateQuery(m map[string]interface{}, queryValues map[string]interface{}) (string, map[string]interface{}) {
	out := make([]string, 0, len(m))

	for key, v := range m {
		tmp := fmt.Sprintf("%s = :%s", key, key)
		out = append(out, tmp)
		queryValues[key] = v
	}

	querySet := strings.Join(out, ", ")
	return querySet, queryValues
}
