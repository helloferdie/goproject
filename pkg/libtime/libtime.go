package libtime

import (
	"database/sql"
	"time"
)

func NullFormat(loc *time.Location, v sql.NullTime) interface{} {
	if v.Valid {
		return v.Time.In(loc).Format("2006-01-02T15:04:05-0700")
	}
	return nil
}
