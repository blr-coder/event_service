package usecases

import (
	"context"
	"event_service/internal/event/repositories"
	"event_service/internal/event/repositories/mock"
	"event_service/internal/event/usecases/usecase_models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

// EventTestsSuite - набор тестов для событий :)
type EventTestsSuite struct {
	suite.Suite

	mockController       *gomock.Controller
	mockIEventRepository *mock_repositories.MockIEventRepository
	eventUseCase         *EventUseCase
}

func (ts *EventTestsSuite) SetupTest() {
	ts.mockController = gomock.NewController(ts.T())
	ts.mockIEventRepository = mock_repositories.NewMockIEventRepository(ts.mockController)
	ts.eventUseCase = NewEventUseCase(&repositories.Repository{
		Event: ts.mockIEventRepository,
	})
}

func (ts *EventTestsSuite) clear() {
	ts.mockController.Finish()
}

func TestEvents(t *testing.T) {
	suite.Run(t, new(EventTestsSuite))
}

func (ts *EventTestsSuite) TestCreate() {
	defer ts.clear()
	ctx := context.Background()

	type args struct {
		ctx    context.Context
		create *usecase_models.CreateEventInput
	}
	tests := []struct {
		name    string
		args    args
		want    *usecase_models.Event
		wantErr error
	}{
		{
			name: "first_OK",
			args: args{
				ctx: ctx,
				create: &usecase_models.CreateEventInput{
					TypeTitle:   "",
					CampaignID:  0,
					InsertionID: 0,
					UserID:      0,
					Cost:        nil,
				},
			},
			want: &usecase_models.Event{
				ID:           0,
				TypeTitle:    "",
				CampaignID:   0,
				InsertionID:  0,
				UserID:       0,
				CostAmount:   0,
				CostCurrency: "",
				CreatedAt:    time.Time{},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		ts.Run(tt.name, func() {
			actual, err := ts.eventUseCase.Create(tt.args.ctx, tt.args.create)
			ts.Require().Equal(tt.wantErr, err)
			ts.Require().Equal(tt.want, actual)
		})
	}
}
