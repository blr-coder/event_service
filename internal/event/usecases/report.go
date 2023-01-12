package usecases

import (
	"context"
	"event_service/internal/event/repositories"
	"event_service/internal/event/repositories/repository_models"
	"event_service/internal/event/usecases/usecase_models"
)

type ReportUseCase struct {
	repository *repositories.Repository
}

func NewReportUseCase(repository *repositories.Repository) *ReportUseCase {
	return &ReportUseCase{repository: repository}
}

func (c *ReportUseCase) List(ctx context.Context, filter *usecase_models.ReportFilter) (*usecase_models.Reports, error) {

	reports, err := c.repository.Report.List(ctx, useCaseReportFilterToRepo(filter))
	if err != nil {
		return nil, err
	}

	return repoReportsToUseCase(reports), err
}

func useCaseReportFilterToRepo(filter *usecase_models.ReportFilter) *repository_models.ReportRepositoryFilter {
	return &repository_models.ReportRepositoryFilter{}
}

func repoReportsToUseCase(repoReports []*repository_models.Report) *usecase_models.Reports {
	reports := make(usecase_models.Reports, 0, len(repoReports))

	for _, r := range repoReports {
		reports = append(reports, repoReportToUseCase(r))
	}

	return &reports
}

func repoReportToUseCase(repoReport *repository_models.Report) *usecase_models.Report {
	return &usecase_models.Report{}
}
