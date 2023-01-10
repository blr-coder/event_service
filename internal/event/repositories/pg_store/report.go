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

	return nil, nil
}
