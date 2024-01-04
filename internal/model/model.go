package model

import "database/sql"

type Timestamp struct {
	CreatedAt sql.NullTime `db:"created_at" insert:"-"`
	UpdatedAt sql.NullTime `db:"updated_at" insert:"-"`
	DeletedAt sql.NullTime `db:"deleted_at" insert:"-"`
}
