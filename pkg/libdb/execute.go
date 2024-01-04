package libdb

import (
	"github.com/helloferdie/golib/v2/liblogger"
	"github.com/jmoiron/sqlx"
)

// Exec return insertedID, total inserted rows & error
func Exec(d *sqlx.DB, query string, values interface{}) (int64, int64, error) {
	result, err := d.NamedExec(query, values)
	if err != nil {
		liblogger.Errorf("Error execute query %v", err)
		return 0, 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, 0, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, 0, err
	}
	return id, rows, nil
}

// Get return single row from query
func Get(db *sqlx.DB, list interface{}, query string, values map[string]interface{}) (bool, error) {
	exist := false
	rows, err := db.NamedQuery(query, values)
	if err != nil {
		liblogger.Errorf("Error get query %v", err)
		return exist, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.StructScan(list)
		if err != nil {
			liblogger.Errorf("Error scan row %v", err)
			return exist, err
		}
		exist = true
	}
	return exist, nil
}
