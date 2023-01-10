package usecase_models

import "time"

type Report struct {
}

type Reports []*Report

type ReportFilter struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}
