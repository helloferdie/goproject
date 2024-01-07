package model

import (
	"database/sql"
	"reflect"
)

type Total struct {
	Total int64 `db:"total"`
}

type Timestamp struct {
	CreatedAt sql.NullTime `db:"created_at" insert:"-"`
	UpdatedAt sql.NullTime `db:"updated_at" insert:"-"`
	DeletedAt sql.NullTime `db:"deleted_at" insert:"-"`
}

type SortField struct {
	Field     string
	Direction string
}

type Pagination struct {
	Cursor     int64
	Limit      int64
	SortFields []SortField
}

// GetChanges return changes from two struct
func GetChanges(original, modified interface{}) map[string]interface{} {
	changes := make(map[string]interface{})
	originalVal := reflect.ValueOf(original).Elem()
	modifiedVal := reflect.ValueOf(modified).Elem()

	for i := 0; i < originalVal.NumField(); i++ {
		// Check if the field has 'insert' tag set to '-'
		if tag, ok := originalVal.Type().Field(i).Tag.Lookup("insert"); ok && tag == "-" {
			continue
		}

		originalField := originalVal.Field(i)
		modifiedField := modifiedVal.Field(i)

		// Compare the values of the original and modified structs
		if !reflect.DeepEqual(originalField.Interface(), modifiedField.Interface()) {
			// If they are not equal, add them to the changes map
			fieldname := originalVal.Type().Field(i).Tag.Get("db")
			changes[fieldname] = modifiedField.Interface()
		}
	}
	return changes
}
