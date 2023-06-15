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
	event, err := s.useCase.Event.Create(ctx, grpcCreateEventToUCInput(grpcRequest))
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
	ucEventsList, err := s.useCase.Event.List(ctx, grpcEventFilterToUCFilter(grpcFilter))
	if err != nil {
		return nil, err
	}

	return ucEventsListToGrpc(ucEventsList), nil
}

func ucEventsListToGrpc(eventsWithCount *usecase_models.Events) *event_proto.Events {
	grpcEvents := &event_proto.Events{
		Count:  eventsWithCount.Count,
		Events: make([]*event_proto.Event, 0, len(eventsWithCount.Events)),
	}

	for _, e := range eventsWithCount.Events {
		grpcEvents.Events = append(grpcEvents.Events, useCaseEventToGRPC(e))
	}

	return grpcEvents
}

func grpcEventFilterToUCFilter(grpcFilter *event_proto.ListEventFilter) *usecase_models.EventFilter {
	ucFilter := &usecase_models.EventFilter{}

	if grpcFilter.TypeTitle != nil {
		ucFilter.TypeTitle = &grpcFilter.TypeTitle.Value
	}
	if grpcFilter.CampaignId != nil {
		ucFilter.CampaignID = &grpcFilter.CampaignId.Value
	}
	if grpcFilter.InsertionId != nil {
		ucFilter.InsertionID = &grpcFilter.InsertionId.Value
	}
	if grpcFilter.UserId != nil {
		ucFilter.UserID = &grpcFilter.UserId.Value
	}
	if grpcFilter.CostCurrency != nil {
		ucFilter.CostCurrency = &grpcFilter.CostCurrency.Value
	}
	if grpcFilter.SortBy != nil {
		ucFilter.SortBy = grpcEventSortByToUc(grpcFilter.SortBy)
	}
	if grpcFilter.SortOrder != 0 {
		order := usecase_models.SortOrderDESC
		ucFilter.SortOrder = &order
	}

	if grpcFilter.PageSize != nil {
		ucFilter.PageSize = &grpcFilter.PageSize.Value
	}
	if grpcFilter.PageNumber != nil {
		ucFilter.PageNumber = &grpcFilter.PageNumber.Value
	}

	return ucFilter
}

func grpcEventSortByToUc(sortBy []event_proto.EventSortBy) (sortByList []usecase_models.EventSortBy) {
	for _, item := range sortBy {
		switch item {
		case event_proto.EventSortBy_CREATED_AT:
			sortByList = append(sortByList, usecase_models.EventSortByCreatedAt)
		case event_proto.EventSortBy_TYPE_TITLE:
			sortByList = append(sortByList, usecase_models.EventSortByTypeTitle)
		case event_proto.EventSortBy_ID:
			sortByList = append(sortByList, usecase_models.EventSortByID)
		case event_proto.EventSortBy_COST_CURRENCY:
			sortByList = append(sortByList, usecase_models.EventSortByCostCurrency)
		}
	}

	return sortByList
}

func grpcCreateEventToUCInput(grpcCreateEvent *event_proto.CreateEventRequest) *usecase_models.CreateEventInput {
	return &usecase_models.CreateEventInput{
		TypeTitle:   grpcCreateEvent.GetTypeTitle(),
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
		TypeTitle:   ucEvent.TypeTitle,
		Cost: &event_proto.Cost{
			Amount:   int64(ucEvent.CostAmount),
			Currency: ucEvent.CostCurrency,
		},
	}
}
