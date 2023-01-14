package pg_store

import (
	"context"
	"event_service/internal/event/repositories/repository_models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ReportPsqlStore struct {
	db *sqlx.DB
}

func NewReportPsqlStore(db *sqlx.DB) *ReportPsqlStore {
	return &ReportPsqlStore{db: db}
}

func (s *ReportPsqlStore) List(ctx context.Context, repositoryFilter *repository_models.ReportRepositoryFilter) (reports []*repository_models.Report, err error) {

	var groupBy, where string

	groupBy = "day"
	where = "TRUE"

	q := fmt.Sprintf(`SELECT type_title, count(type_title), sum(cost_amount), date_trunc('%s', created_at) AS date
	FROM events
	WHERE %s
	GROUP BY type_title, date_trunc('%s', created_at)
	ORDER BY date_trunc('%s', created_at)`, groupBy, where, groupBy, groupBy)

	err = s.db.SelectContext(ctx, &reports, q)

	return reports, nil
}

// TODO: Add decode func

func (s *ReportPsqlStore) repositoryFilterToQuery(repositoryFilter *repository_models.ReportRepositoryFilter) (string, []any) {

	return "", nil
}
