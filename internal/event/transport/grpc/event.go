package grpc

import (
	"context"
	"event_service/api/grpc/event_proto"
)

type EventServiceServer struct {
	event_proto.UnimplementedEventServiceServer
	//eventService *service.EventService
}

func NewEventServiceServer() *EventServiceServer {
	return &EventServiceServer{}
}

func (s *EventServiceServer) Create(
	ctx context.Context,
	request *event_proto.CreateEventRequest,
) (*event_proto.Event, error) {

	//TODO implement me
	panic("implement me Create")
}

func (s *EventServiceServer) List(
	ctx context.Context,
	request *event_proto.ListEventsRequest,
) (*event_proto.Events, error) {

	//TODO implement me
	panic("implement me List")
}
