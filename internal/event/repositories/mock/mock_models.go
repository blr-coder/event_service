package mock_repositories

import (
	"event_service/internal/event/repositories/repository_models"
	"github.com/jaswdr/faker"
	"testing"
	"time"
)

func NewRepoEvent(t *testing.T) *repository_models.EventRepositoryDTO {
	t.Helper()
	f := faker.New()
	return &repository_models.EventRepositoryDTO{
		ID: int64(f.RandomDigit()),
		TypeTitle: f.RandomStringElement([]string{
			"one", "two", "three",
		}),
		CampaignID:  int64(f.RandomDigit()),
		InsertionID: int64(f.RandomDigit()),
		UserID:      int64(f.RandomDigit()),
		CostAmount:  uint64(f.RandomDigit()),
		CostCurrency: f.RandomStringElement([]string{
			"EUR", "USD", "PLN",
		}),
		CreatedAt: f.Time().Time(time.Now()).UTC(),
	}
}

func NewRepoEventList(t *testing.T) (events []*repository_models.EventRepositoryDTO, count uint64) {
	t.Helper()

	f := faker.New()

	count = uint64(f.RandomDigit())
	for i := 0; i < int(count); i++ {
		events = append(events, NewRepoEvent(t))
	}

	return events, count
}
