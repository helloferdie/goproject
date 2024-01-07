package model

type Country struct {
	ID       int64  `db:"id" insert:"-"`
	ISO2     string `db:"iso2"`
	ISO3     string `db:"iso3"`
	Callcode string `db:"callcode"`
	Name     string `db:"name"`
	Timestamp
}
