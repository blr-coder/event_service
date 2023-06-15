package usecases

import (
	"context"
	"event_service/internal/event/repositories"
	mock_repositories "event_service/internal/event/repositories/mock"
	"event_service/internal/event/repositories/repository_models"
	"event_service/internal/event/usecases/usecase_errors"
	"event_service/internal/event/usecases/usecase_models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
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

func (ts *EventTestsSuite) TestListOK() {
	defer ts.clear()
	ctx := context.Background()

	eventFilter := &usecase_models.EventFilter{}

	repoEvents, repoCount := mock_repositories.NewRepoEventList(ts.T())

	ts.mockIEventRepository.
		EXPECT().
		List(ctx, useCaseEventFilterToRepo(eventFilter)).
		Return(repoEvents, repoCount, nil).
		Times(1)

	type args struct {
		ctx    context.Context
		filter *usecase_models.EventFilter
	}
	tests := []struct {
		name    string
		args    args
		want    *usecase_models.Events
		wantErr error
	}{
		{
			name: "OK",
			args: args{
				ctx:    ctx,
				filter: eventFilter,
			},
			want:    repoEventsToUseCase(repoEvents, repoCount),
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		ts.Run(tt.name, func() {
			actual, err := ts.eventUseCase.List(tt.args.ctx, tt.args.filter)
			ts.Require().Equal(tt.want, actual)
			ts.Require().Equal(tt.wantErr, err)
			// TODO: В чэм разница??? Какой юзать?
			// ts.Assert().Equal(tt.wantErr, err)
			//ts.Assert().Equal(tt.want, actual)
		})
	}
}

func (ts *EventTestsSuite) TestListRepoErr() {
	defer ts.clear()
	ctx := context.Background()

	eventFilter := &usecase_models.EventFilter{}

	ts.mockIEventRepository.
		EXPECT().
		List(ctx, useCaseEventFilterToRepo(eventFilter)).
		Return(nil, uint64(0), usecase_errors.UnexpectedStoreError).
		Times(1)

	type args struct {
		ctx    context.Context
		filter *usecase_models.EventFilter
	}
	tests := []struct {
		name    string
		args    args
		want    *usecase_models.Events
		wantErr error
	}{
		{
			name: "STORE ERROR",
			args: args{
				ctx:    ctx,
				filter: eventFilter,
			},
			want:    nil,
			wantErr: usecase_errors.UnexpectedStoreError,
		},
	}

	for _, tt := range tests {
		ts.Run(tt.name, func() {
			actual, err := ts.eventUseCase.List(tt.args.ctx, tt.args.filter)
			ts.Require().Equal(tt.want, actual)
			ts.Require().Equal(tt.wantErr, err)
			ts.Require().NotEmpty(err)
		})
	}
}
