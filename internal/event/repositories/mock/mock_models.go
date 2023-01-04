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
