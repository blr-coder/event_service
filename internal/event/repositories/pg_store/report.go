package pg_store

import (
	"context"
	"event_service/internal/event/repositories/repository_models"
	"github.com/jmoiron/sqlx"
)

type ReportPsqlStore struct {
	db *sqlx.DB
}

func NewReportPsqlStore(db *sqlx.DB) *ReportPsqlStore {
	return &ReportPsqlStore{db: db}
}

func (s *ReportPsqlStore) List(ctx context.Context, repositoryFilter *repository_models.ReportRepositoryFilter) ([]*repository_models.Report, error) {

	// TODO: Need to fix query for HOUR/DAY/WEEK/MONTH grouping
	_ = `SELECT type_title, count(type_title) AS count_of_type, sum(cost_amount), DATE(created_at)
	FROM events
	GROUP BY type_title, DATE(created_at)`

	return nil, nil
}
