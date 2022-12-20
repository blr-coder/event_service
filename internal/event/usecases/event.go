package usecases

import (
	"context"
	"event_service/internal/event/repositories"
	"event_service/internal/event/repositories/repository_models"
	"event_service/internal/event/usecases/usecase_models"
)

type EventUseCase struct {
	repository *repositories.Repository
}

func NewEventUseCase(repo *repositories.Repository) *EventUseCase {
	return &EventUseCase{repository: repo}
}

func (c *EventUseCase) Create(ctx context.Context, createEvent *usecase_models.CreateEventInput) (*usecase_models.Event, error) {
	// TODO: Validate

	event, err := c.repository.Event.Create(ctx, useCaseCreateEventDtoToRepo(createEvent))
	if err != nil {
		return nil, err
	}

	return repoEventDtoToUseCase(event), nil
}

func useCaseCreateEventDtoToRepo(createEventUC *usecase_models.CreateEventInput) *repository_models.CreateEventRepositoryDTO {
	return &repository_models.CreateEventRepositoryDTO{
		TypeTitle:    createEventUC.TypeTitle,
		CampaignID:   createEventUC.CampaignID,
		InsertionID:  createEventUC.InsertionID,
		UserID:       createEventUC.UserID,
		CostAmount:   createEventUC.Cost.Amount,
		CostCurrency: createEventUC.Cost.Currency,
	}
}

func repoEventDtoToUseCase(repoEvent *repository_models.EventRepositoryDTO) *usecase_models.Event {
	return &usecase_models.Event{
		ID:           repoEvent.ID,
		TypeTitle:    repoEvent.TypeTitle,
		CampaignID:   repoEvent.CampaignID,
		InsertionID:  repoEvent.InsertionID,
		UserID:       repoEvent.UserID,
		CostAmount:   repoEvent.CostAmount,
		CostCurrency: repoEvent.CostCurrency,
		CreatedAt:    repoEvent.CreatedAt,
	}
}
