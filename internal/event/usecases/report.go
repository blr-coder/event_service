package usecases

import (
	"context"
	"event_service/internal/event/repositories"
	"event_service/internal/event/repositories/repository_models"
	"event_service/internal/event/usecases/usecase_models"
	"time"
)

type ReportUseCase struct {
	repository *repositories.Repository
}

func NewReportUseCase(repository *repositories.Repository) *ReportUseCase {
	return &ReportUseCase{repository: repository}
}

func (c *ReportUseCase) List(ctx context.Context, filter *usecase_models.ReportFilter) (usecase_models.Reports, error) {
	if err := filter.Validate(); err != nil {
		// TODO: add good errors handling
		return nil, err
	}

	reports, err := c.repository.Report.List(ctx, useCaseReportFilterToRepo(filter))
	if err != nil {
		return nil, err
	}

	return repoReportsToUseCase(reports), err
}

func useCaseReportFilterToRepo(filter *usecase_models.ReportFilter) *repository_models.ReportRepositoryFilter {
	repoFilter := &repository_models.ReportRepositoryFilter{
		From:    filter.From,
		To:      filter.To,
		GroupBy: ucGroupByToRepo[filter.GroupBy],
	}

	if filter.UserID != nil {
		repoFilter.UserID = filter.UserID
	}

	return repoFilter
}

var ucGroupByToRepo = map[usecase_models.ReportGroupBy]repository_models.GroupBy{
	usecase_models.ReportGroupByHour:  repository_models.GroupByHour,
	usecase_models.ReportGroupByDay:   repository_models.GroupByDay,
	usecase_models.ReportGroupByWeek:  repository_models.GroupByWeek,
	usecase_models.ReportGroupByMonth: repository_models.GroupByMonth,
}

func repoReportsToUseCase(repoReports []*repository_models.Report) usecase_models.Reports {
	reports := make(usecase_models.Reports, 0, len(repoReports))

	for _, r := range repoReports {
		reports = append(reports, repoReportToUseCase(r))
	}

	return reports
}

func repoReportToUseCase(repoReport *repository_models.Report) *usecase_models.Report {
	return &usecase_models.Report{
		Type:  repoReport.TypeTitle,
		Count: uint64(repoReport.Count),
		Sum:   repoReport.Sum,
		Date:  time.Time{},
	}
}
