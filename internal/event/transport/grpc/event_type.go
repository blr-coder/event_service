package grpc

import (
	"context"
	"event_service/api/grpc/event_type_proto"
	"event_service/internal/event/usecases"
	"event_service/internal/event/usecases/usecase_models"
	"event_service/pkg/utils"
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

func (s *EventTypeServiceServer) List(ctx context.Context, filter *event_type_proto.EventTypeFilter) (*event_type_proto.EventTypes, error) {

	types, err := s.useCase.EventType.List(ctx, grpcListFilterToModelFilter(filter))
	if err != nil {
		// TODO: Add decodeError func
		return nil, err
	}

	return eventTypesModelToGRPC(types), err
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
	in *event_type_proto.DeleteEventTypeRequest,
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
		Title:     eventType.Title,
		CreatedAt: timestamppb.New(eventType.CreatedAt),
		UpdatedAt: timestamppb.New(eventType.UpdatedAt),
	}
}

func grpcListFilterToModelFilter(grpcFilter *event_type_proto.EventTypeFilter) *usecase_models.EventTypeFilter {
	return &usecase_models.EventTypeFilter{
		Titles: grpcFilter.Titles,
		Search: utils.Pointer(grpcFilter.GetSearch().GetValue()),
	}
}

func eventTypesModelToGRPC(types usecase_models.EventTypes) *event_type_proto.EventTypes {
	grpcTypes := &event_type_proto.EventTypes{
		EventTypes: make([]*event_type_proto.EventType, 0, len(types)),
	}

	for _, t := range types {
		grpcTypes.EventTypes = append(grpcTypes.EventTypes, eventTypeModelToGRPC(t))
	}

	return grpcTypes
}
