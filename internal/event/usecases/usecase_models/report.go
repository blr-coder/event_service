package usecase_models

import "time"

type Report struct {
}

type Reports []*Report

type ReportFilter struct {
	From    time.Time     `json:"from"`
	To      time.Time     `json:"to"`
	GroupBy ReportGroupBy `json:"group_by"`
}

type ReportGroupBy string

const (
	EventSortByHour  EventSortBy = "hour"
	EventSortByDay   EventSortBy = "day"
	EventSortByWeek  EventSortBy = "week"
	EventSortByMonth EventSortBy = "month"
)
