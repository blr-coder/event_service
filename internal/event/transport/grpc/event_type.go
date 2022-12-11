package grpc

import (
	"context"
	"event_service/api/grpc/event_type_proto"
	"event_service/internal/event/repositories/repository_models"
	"event_service/internal/event/usecases"
	"event_service/internal/event/usecases/usecase_models"
	"github.com/golang/protobuf/ptypes/empty"
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
	request *event_type_proto.UpdateEventTypeRequest,
) (*event_type_proto.EventType, error) {

	eventType, err := s.useCase.EventType.Update(ctx, grpcUpdateEventTypeToModelUpdate(request))
	if err != nil {
		// TODO: Add decodeError func
		return nil, err
	}

	return eventTypeModelToGRPC(eventType), err
}

func (s *EventTypeServiceServer) Delete(
	ctx context.Context,
	request *event_type_proto.DeleteEventTypeRequest,
) (*emptypb.Empty, error) {

	return &empty.Empty{}, s.useCase.EventType.Delete(ctx, grpcDeleteEventTypeToModelUpdate(request))
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
	useCaseFilter := &usecase_models.EventTypeFilter{
		Titles: grpcFilter.Titles,
	}

	if grpcFilter.Search != nil {
		useCaseFilter.Search = &grpcFilter.Search.Value
	}

	if grpcFilter.OrderBy != nil {
		useCaseFilter.OrderBy = grpcOrderByToRepo(grpcFilter.OrderBy)
	}

	if grpcFilter.OrderDirection != event_type_proto.OrderDirection_EMPTY {

		var direction repository_models.OrderDirection

		switch grpcFilter.OrderDirection {
		case event_type_proto.OrderDirection_ASC:
			direction = repository_models.OrderDirectionTypeASC
			useCaseFilter.OrderDirection = &direction
		case event_type_proto.OrderDirection_DESC:
			direction = repository_models.OrderDirectionTypeDESC
			useCaseFilter.OrderDirection = &direction
		}
	}

	if grpcFilter.PageSize != nil {
		useCaseFilter.PageSize = &grpcFilter.PageSize.Value
	}
	if grpcFilter.PageNumber != nil {
		useCaseFilter.PageNumber = &grpcFilter.PageNumber.Value
	}

	return useCaseFilter
}

func grpcOrderByToRepo(orderBy []event_type_proto.OrderBy) (orderByList repository_models.OrderByList) {
	for _, item := range orderBy {
		switch item {
		case event_type_proto.OrderBy_TITLE:
			orderByList = append(orderByList, repository_models.OrderByTypeTitle)
		case event_type_proto.OrderBy_CREATED_AT:
			orderByList = append(orderByList, repository_models.OrderByTypeCreatedAt)
		}
	}
	return orderByList
}

func eventTypesModelToGRPC(types usecase_models.EventTypes) *event_type_proto.EventTypes {
	grpcTypes := &event_type_proto.EventTypes{
		Count:      100000,
		EventTypes: make([]*event_type_proto.EventType, 0, len(types)),
	}

	for _, t := range types {
		grpcTypes.EventTypes = append(grpcTypes.EventTypes, eventTypeModelToGRPC(t))
	}

	return grpcTypes
}

func grpcUpdateEventTypeToModelUpdate(grpcUpdate *event_type_proto.UpdateEventTypeRequest) *usecase_models.UpdateEventTypeInput {
	return &usecase_models.UpdateEventTypeInput{
		Title:    grpcUpdate.Title,
		NewTitle: grpcUpdate.NewTitle,
	}
}

func grpcDeleteEventTypeToModelUpdate(grpcDelete *event_type_proto.DeleteEventTypeRequest) *usecase_models.DeleteEventTypeInput {
	return &usecase_models.DeleteEventTypeInput{
		Title: grpcDelete.Title,
	}
}
