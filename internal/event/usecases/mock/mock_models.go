package mock_usecases

import (
	"event_service/internal/event/usecases/usecase_models"
	"github.com/jaswdr/faker"
	"testing"
	"time"
)

func NewEvent(t *testing.T) *usecase_models.Event {
	t.Helper()
	f := faker.New()
	return &usecase_models.Event{
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
