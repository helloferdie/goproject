package libdb

import (
	"reflect"
	"strings"
)

type Config struct {
	Table          string
	Fields         string
	CondSoftDelete string
	QuerySelect    string
	QueryInsert    string
	QueryUpdate    string
	QueryDelete    string
}

// NewConfig
func NewConfig(model interface{}, table string, softDelete bool) *Config {
	cfg := &Config{
		Table: table,
	}

	if softDelete {
		cfg.CondSoftDelete = "AND deleted_at IS NULL "
		cfg.QueryDelete = "UPDATE " + cfg.Table + " SET deleted_at = UTC_TIMESTAMP WHERE 1=1 %s"
	} else {
		cfg.QueryDelete = "DELETE FROM " + cfg.Table + " WHERE 1=1 %s"
	}
	cfg.QueryUpdate = "UPDATE " + cfg.Table + " SET %s WHERE 1=1 %s"

	queryCols, insertCols := getTags(model)

	cfg.Fields = strings.Join(queryCols, ", ")
	cfg.QuerySelect = "SELECT " + cfg.Fields + " FROM " + cfg.Table + " WHERE 1=1 %s " + cfg.CondSoftDelete

	if len(insertCols) > 0 {
		cfg.QueryInsert = "INSERT INTO " + cfg.Table + " (" + strings.Join(insertCols, ", ") +
			") VALUES (:" + strings.Join(insertCols, ", :") + ")"
	}

	return cfg
}

// getTags return slices of db tags and db tags without insert:"-".
func getTags(model interface{}) ([]string, []string) {
	tags := []string{}
	skipInsertTags := []string{}

	rType := reflect.TypeOf(model)

	// Iterate through the fields of the struct.
	for i := 0; i < rType.NumField(); i++ {
		field := rType.Field(i)

		// Handle nested structs (anonymous fields).
		if field.Anonymous && field.Type.Kind() == reflect.Struct {
			innerTags, innerSkipInsertTags := getTags(reflect.New(field.Type).Elem().Interface())
			tags = append(tags, innerTags...)
			skipInsertTags = append(skipInsertTags, innerSkipInsertTags...)
		} else {
			dbTag := field.Tag.Get("db")
			if dbTag != "" {
				tags = append(tags, dbTag)
				// Check for insert:"-"
				if field.Tag.Get("insert") != "-" {
					skipInsertTags = append(skipInsertTags, dbTag)
				}
			}
		}
	}
	return tags, skipInsertTags
}
