package usecases

import (
	"context"
	"event_service/internal/event/repositories"
	mock_repositories "event_service/internal/event/repositories/mock"
	"event_service/internal/event/repositories/repository_models"
	"event_service/internal/event/usecases/usecase_models"
	"github.com/golang/mock/gomock"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEventTypeUseCase_Create(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	f := faker.New()

	ctx := context.Background()

	eventTypeRepository := mock_repositories.NewMockIEventTypeRepository(controller)
	eventTypeUseCase := NewEventTypeUseCase(&repositories.Repository{EventType: eventTypeRepository})

	repoEventType := &repository_models.EventTypeRepositoryDTO{
		Title:     f.RandomStringWithLength(5),
		CreatedAt: f.Time().Time(time.Now()).UTC(),
		UpdatedAt: f.Time().Time(time.Now()).UTC(),
		DeletedAt: nil,
	}

	ucEventType := &usecase_models.EventType{
		Title:     repoEventType.Title,
		CreatedAt: repoEventType.CreatedAt,
		UpdatedAt: repoEventType.UpdatedAt,
	}

	eventTypeRepository.
		EXPECT().
		Create(ctx, &repository_models.CreateEventTypeRepositoryDTO{
			Title: repoEventType.Title,
		}).
		Return(repoEventType, nil).
		Times(1)

	actual, err := eventTypeUseCase.Create(ctx, &usecase_models.CreateEventTypeInput{
		Title: ucEventType.Title,
	})

	assert.NoError(t, err)

	assert.NotNil(t, actual)
	assert.IsType(t, &usecase_models.EventType{}, actual)
	assert.Equal(t, ucEventType, actual)
}
