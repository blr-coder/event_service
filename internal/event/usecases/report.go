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

	_, err := c.repository.Report.List(ctx, &repository_models.ReportRepositoryFilter{})
	if err != nil {
		return nil, err
	}

	return nil, err
}
