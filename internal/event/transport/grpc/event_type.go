package grpc

import (
	"context"
	"event_service/api/grpc/event_type_proto"
	"event_service/internal/event/usecases"
	"event_service/internal/event/usecases/usecase_models"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EventTypeServiceServer struct {
	event_type_proto.UnimplementedEventTypeServiceServer
	useCase *usecases.UseCase
}

func NewEventTypeServiceServer(useCase *usecases.UseCase) *EventTypeServiceServer {
	return &EventTypeServiceServer{
		useCase: useCase,
	}
}

func (s *EventTypeServiceServer) Create(
	ctx context.Context,
	request *event_type_proto.CreateEventTypeRequest,
) (*event_type_proto.EventType, error) {

	eventType, err := s.useCase.EventType.Create(ctx, grpcCreateEventTypeToModelCreate(request))
	if err != nil {
		// TODO: Add decodeError func
		return nil, err
	}
	return eventTypeModelToGRPC(eventType), err
}

func (s *EventTypeServiceServer) Get(
	ctx context.Context,
	in *event_type_proto.GetEventTypeRequest,
) (*event_type_proto.EventType, error) {

	eventType, err := s.useCase.EventType.Get(
		ctx,
		&usecase_models.GetEventTypeInput{
			ID: in.GetId(),
		},
	)
	if err != nil {
		// TODO: Add decodeError func
		return nil, err
	}

	return eventTypeModelToGRPC(eventType), err
}

func (s *EventTypeServiceServer) List(ctx context.Context, in *emptypb.Empty) (*event_type_proto.EventTypes, error) {

	//TODO: implement me
	panic("implement me")
}

func (s *EventTypeServiceServer) Update(
	ctx context.Context,
	in *event_type_proto.UpdateEventTypeRequest,
) (*event_type_proto.EventType, error) {

	//TODO: implement me
	panic("implement me")
}

func (s *EventTypeServiceServer) Delete(
	ctx context.Context,
	in *event_type_proto.GetEventTypeRequest,
) (*emptypb.Empty, error) {

	//TODO: implement me
	panic("implement me")
}

func grpcCreateEventTypeToModelCreate(
	grpcEventType *event_type_proto.CreateEventTypeRequest,
) *usecase_models.CreateEventTypeInput {
	return &usecase_models.CreateEventTypeInput{
		Title: grpcEventType.GetTitle(),
	}
}

func eventTypeModelToGRPC(
	eventType *usecase_models.EventType,
) *event_type_proto.EventType {
	return &event_type_proto.EventType{
		Id:        eventType.ID,
		Title:     eventType.Title,
		CreatedAt: timestamppb.New(eventType.CreatedAt),
	}
}
