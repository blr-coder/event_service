package usecases

import (
	"context"
	"event_service/internal/event/repositories"
	"event_service/internal/event/repositories/mock"
	"event_service/internal/event/repositories/repository_models"
	"event_service/internal/event/usecases/usecase_errors"
	"event_service/internal/event/usecases/usecase_models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
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

func (ts *EventTestsSuite) TestCreateOK() {
	defer ts.clear()
	ctx := context.Background()
	newRepoEvent := mock_repositories.NewRepoEvent(ts.T())

	ts.mockIEventRepository.
		EXPECT().
		Create(ctx, &repository_models.CreateEventRepositoryDTO{
			TypeTitle:    newRepoEvent.TypeTitle,
			CampaignID:   newRepoEvent.CampaignID,
			InsertionID:  newRepoEvent.InsertionID,
			UserID:       newRepoEvent.UserID,
			CostAmount:   newRepoEvent.CostAmount,
			CostCurrency: newRepoEvent.CostCurrency,
		}).
		Return(newRepoEvent, nil).
		Times(1)

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
			name: "OK",
			args: args{
				ctx: ctx,
				create: &usecase_models.CreateEventInput{
					TypeTitle:   newRepoEvent.TypeTitle,
					CampaignID:  newRepoEvent.CampaignID,
					InsertionID: newRepoEvent.InsertionID,
					UserID:      newRepoEvent.UserID,
					Cost: &usecase_models.Cost{
						Amount:   newRepoEvent.CostAmount,
						Currency: newRepoEvent.CostCurrency,
					},
				},
			},
			want: &usecase_models.Event{
				ID:           newRepoEvent.ID,
				TypeTitle:    newRepoEvent.TypeTitle,
				CampaignID:   newRepoEvent.CampaignID,
				InsertionID:  newRepoEvent.InsertionID,
				UserID:       newRepoEvent.UserID,
				CostAmount:   newRepoEvent.CostAmount,
				CostCurrency: newRepoEvent.CostCurrency,
				CreatedAt:    newRepoEvent.CreatedAt,
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

func (ts *EventTestsSuite) TestCreateValidateError() {
	defer ts.clear()
	ctx := context.Background()
	newRepoEvent := mock_repositories.NewRepoEvent(ts.T())

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
			name: "EMPTY TYPE TITLE",
			args: args{
				ctx: ctx,
				create: &usecase_models.CreateEventInput{
					CampaignID:  newRepoEvent.CampaignID,
					InsertionID: newRepoEvent.InsertionID,
					UserID:      newRepoEvent.UserID,
					Cost: &usecase_models.Cost{
						Amount:   newRepoEvent.CostAmount,
						Currency: newRepoEvent.CostCurrency,
					},
				},
			},
			want: nil,
			wantErr: &usecase_errors.ValidationErr{
				ErrMessage: "type_title: cannot be blank.",
			},
		},
		{
			name: "EMPTY COST",
			args: args{
				ctx: ctx,
				create: &usecase_models.CreateEventInput{
					TypeTitle:   newRepoEvent.TypeTitle,
					CampaignID:  newRepoEvent.CampaignID,
					InsertionID: newRepoEvent.InsertionID,
					UserID:      newRepoEvent.UserID,
					Cost:        nil,
				},
			},
			want: nil,
			wantErr: &usecase_errors.ValidationErr{
				ErrMessage: "cost: cannot be blank.",
			},
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
