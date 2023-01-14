package repository_models

import "time"

type Report struct {
	TypeTitle string    `db:"type_title"`
	Count     int64     `db:"count"`
	Sum       int64     `db:"sum"`
	Date      time.Time `db:"date"`
}

type ReportRepositoryFilter struct {
}
