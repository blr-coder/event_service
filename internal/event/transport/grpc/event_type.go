package grpc

import (
	"context"
	"event_service/api/grpc/event_type_proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type EventTypeServiceServer struct {
	event_type_proto.UnimplementedEventTypeServiceServer
}

func NewEventTypeServiceServer() *EventTypeServiceServer {
	return &EventTypeServiceServer{}
}

func (e EventTypeServiceServer) Create(
	ctx context.Context,
	in *event_type_proto.CreateEventTypeRequest,
) (*event_type_proto.EventType, error) {

	//TODO: implement me
	panic("implement me")
}

func (e EventTypeServiceServer) Get(
	ctx context.Context,
	in *event_type_proto.GetEventTypeRequest,
) (*event_type_proto.EventType, error) {

	//TODO: implement me
	panic("implement me")
}

func (e EventTypeServiceServer) List(ctx context.Context, in *emptypb.Empty) (*event_type_proto.EventTypes, error) {

	//TODO: implement me
	panic("implement me")
}

func (e EventTypeServiceServer) Update(
	ctx context.Context,
	in *event_type_proto.UpdateEventTypeRequest,
) (*event_type_proto.EventType, error) {

	//TODO: implement me
	panic("implement me")
}

func (e EventTypeServiceServer) Delete(
	ctx context.Context,
	in *event_type_proto.GetEventTypeRequest,
) (*emptypb.Empty, error) {

	//TODO: implement me
	panic("implement me")
}
