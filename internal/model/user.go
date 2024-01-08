package model

import "database/sql"

type User struct {
	ID              int64        `db:"id" insert:"-"`
	UUID            string       `db:"uuid" `
	AccountNo       string       `db:"account_no"`
	FirstName       string       `db:"first_name"`
	LastName        string       `db:"last_name"`
	Email           string       `db:"email"`
	EmailVerifiedAt sql.NullTime `db:"email_verified_at"`
	Phone           string       `db:"phone"`
	PhoneVerifiedAt sql.NullTime `db:"phone_verified_at"`
	Password        string       `db:"password" `
	IsActive        bool         `db:"is_active"`
	LastLoginAt     sql.NullTime `db:"last_login_at"`
	DefaultLanguage string       `db:"default_language"`
	DefaultTimezone string       `db:"default_timezone"`
	Timestamp
}
