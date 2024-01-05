package libdb

import (
	"reflect"
	"strings"
)

// Config stores the configuration and generated SQL queries for a model and table.
type Config struct {
	Table          string
	Fields         string
	CondSoftDelete string
	QuerySelect    string
	QuerySelectRaw string // without softdelete
	QueryInsert    string
	QueryUpdate    string
	QueryDelete    string
	QueryTotal     string
	QueryTotalRaw  string // without softdelete
}

// NewConfig creates a new Config for a given model and table.
func NewConfig(model interface{}, table string, softDelete bool) *Config {
	cfg := &Config{
		Table: table,
	}

	queryCols, insertCols := getTags(model)
	cfg.Fields = strings.Join(queryCols, ", ")

	// Construct insert query
	if len(insertCols) > 0 {
		cfg.QueryInsert = "INSERT INTO " + cfg.Table + " (" + strings.Join(insertCols, ", ") +
			") VALUES (:" + strings.Join(insertCols, ", :") + ")"
	}

	// Construct delete query
	if softDelete {
		cfg.CondSoftDelete = "AND deleted_at IS NULL "
		cfg.QueryDelete = "UPDATE " + cfg.Table + " SET deleted_at = UTC_TIMESTAMP WHERE 1=1 %s"
	} else {
		cfg.QueryDelete = "DELETE FROM " + cfg.Table + " WHERE 1=1 %s"
	}

	// Construct update query
	cfg.QueryUpdate = "UPDATE " + cfg.Table + " SET %s WHERE 1=1 %s"

	// Construct select query
	cfg.QuerySelectRaw = "SELECT " + cfg.Fields + " FROM " + cfg.Table + " WHERE 1=1 %s "
	cfg.QuerySelect = cfg.QuerySelectRaw + cfg.CondSoftDelete

	// Construct total query
	cfg.QueryTotalRaw = "SELECT COUNT(*) as total FROM " + cfg.Table + " WHERE 1=1 %s "
	cfg.QueryTotal = cfg.QueryTotalRaw + cfg.CondSoftDelete

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
