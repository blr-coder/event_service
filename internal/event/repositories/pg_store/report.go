package pg_store

import (
	"context"
	"event_service/internal/event/repositories/repository_models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type ReportPsqlStore struct {
	db *sqlx.DB
}

func NewReportPsqlStore(db *sqlx.DB) *ReportPsqlStore {
	return &ReportPsqlStore{db: db}
}

func (s *ReportPsqlStore) List(ctx context.Context, repositoryFilter *repository_models.ReportRepositoryFilter) (reports []*repository_models.Report, err error) {

	err = s.db.SelectContext(ctx, &reports, s.repositoryFilterToQuery(repositoryFilter))
	if err != nil {
		return nil, err
	}

	return reports, nil
}

func (s *ReportPsqlStore) repositoryFilterToQuery(filter *repository_models.ReportRepositoryFilter) string {

	query := fmt.Sprintf("SELECT type_title, count(type_title), sum(cost_amount), date_trunc('%s', created_at) AS date FROM events WHERE created_at >= '%s' AND created_at < '%s'", filter.GroupBy, filter.From.Format(time.RFC3339), filter.To.Format(time.RFC3339))

	if filter.UserID != nil {
		query = fmt.Sprintf("%s AND user_id = %d", query, *filter.UserID)
	}

	//TODO: Add check other fields

	query = fmt.Sprintf("%s GROUP BY type_title, date_trunc('%s', events.created_at) ORDER BY date_trunc('%s', created_at)", query, filter.GroupBy, filter.GroupBy)

	query = s.db.Rebind(query)

	return query
}
