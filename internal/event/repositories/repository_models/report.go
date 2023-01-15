package repository_models

import "time"

type Report struct {
	TypeTitle string    `db:"type_title"`
	Count     int64     `db:"count"`
	Sum       int64     `db:"sum"`
	Date      time.Time `db:"date"`
}

type ReportRepositoryFilter struct {
	From    time.Time `db:"from"`
	To      time.Time `db:"to"`
	GroupBy GroupBy   `db:"group_by"`
	UserID  *int64    `db:"user_id"`
}

type GroupBy string

const (
	GroupByHour  GroupBy = "hour"
	GroupByDay   GroupBy = "day"
	GroupByWeek  GroupBy = "week"
	GroupByMonth GroupBy = "month"
)
