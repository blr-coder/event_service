package grpc

import (
	"context"
	"event_service/api/grpc/event_proto"
	"event_service/internal/event/usecases"
	"event_service/internal/event/usecases/usecase_models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EventServiceServer struct {
	event_proto.UnimplementedEventServiceServer
	useCase *usecases.UseCase
}

func NewEventServiceServer(useCase *usecases.UseCase) *EventServiceServer {
	return &EventServiceServer{
		useCase: useCase,
	}
}

func (s *EventServiceServer) Create(
	ctx context.Context,
	grpcRequest *event_proto.CreateEventRequest,
) (*event_proto.Event, error) {

	event, err := s.useCase.Event.Create(ctx, grpcCreateEventToUseCaseInput(grpcRequest))
	if err != nil {
		// TODO: Add decodeError func
		return nil, err
	}

	return useCaseEventToGRPC(event), nil
}

func (s *EventServiceServer) List(
	ctx context.Context,
	grpcFilter *event_proto.ListEventFilter,
) (*event_proto.Events, error) {

	//TODO implement me
	panic("implement me List")
}

func grpcCreateEventToUseCaseInput(grpcCreateEvent *event_proto.CreateEventRequest) *usecase_models.CreateEventInput {
	return &usecase_models.CreateEventInput{
		TypeID:      grpcCreateEvent.GetTypeId(),
		CampaignID:  grpcCreateEvent.GetCampaignId(),
		InsertionID: grpcCreateEvent.GetInsertionId(),
		UserID:      grpcCreateEvent.GetUserId(),
		Cost: &usecase_models.Cost{
			Amount:   uint64(grpcCreateEvent.Cost.GetAmount()),
			Currency: grpcCreateEvent.Cost.GetCurrency(),
		},
	}
}

func useCaseEventToGRPC(ucEvent *usecase_models.Event) *event_proto.Event {
	return &event_proto.Event{
		Id:          ucEvent.ID,
		CampaignId:  ucEvent.CampaignID,
		InsertionId: ucEvent.InsertionID,
		UserId:      ucEvent.UserID,
		CreatedAt:   timestamppb.New(ucEvent.CreatedAt),
		TypeId:      ucEvent.TypeID,
		Cost: &event_proto.Cost{
			Amount:   int64(ucEvent.CostAmount),
			Currency: ucEvent.CostCurrency,
		},
	}
}
