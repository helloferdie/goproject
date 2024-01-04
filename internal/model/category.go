package model

type Category struct {
	ID          int64  `db:"id" insert:"-"`
	Name        string `db:"name"`
	Description string `db:"description"`
	IsActive    string `db:"is_active"`
	FirebaseID  string `db:"firebase_id"`
	Timestamp
}
