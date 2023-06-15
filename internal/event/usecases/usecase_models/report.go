package usecase_models

import (
	"event_service/internal/event/usecases/usecase_errors"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Report struct {
	Type  string    `json:"type"`
	Count uint64    `json:"count"`
	Sum   int64     `json:"sum"`
	Date  time.Time `json:"date"`
}

type Reports []*Report

type ReportFilter struct {
	From    time.Time     `json:"from"`
	To      time.Time     `json:"to"`
	GroupBy ReportGroupBy `json:"group_by"`
	UserID  *int64        `json:"user_id"`
}

func (f *ReportFilter) Validate() error {
	err := validation.ValidateStruct(
		f,
		//TODO: How I can validate time on this case? (And maybe we need validate period between from and to?)
		validation.Field(&f.From, validation.Required),
		validation.Field(&f.To, validation.Required),
		validation.Field(&f.GroupBy, validation.In(
			ReportGroupByHour,
			ReportGroupByDay,
			ReportGroupByWeek,
			ReportGroupByMonth,
		)),
	)

	if err != nil {
		// TODO: Is it good solution?
		return usecase_errors.NewValidationErr(err)
	}

	return nil
}

type ReportGroupBy string

const (
	ReportGroupByHour  ReportGroupBy = "hour"
	ReportGroupByDay   ReportGroupBy = "day"
	ReportGroupByWeek  ReportGroupBy = "week"
	ReportGroupByMonth ReportGroupBy = "month"
)
