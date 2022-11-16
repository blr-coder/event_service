package grpc

import (
	"event_service/api/grpc/event_proto"
	"event_service/api/grpc/event_type_proto"
	"google.golang.org/grpc"
)

func NewGRPCServer(eventServer *EventServiceServer, eventTypeServer *EventTypeServiceServer) *grpc.Server {

	grpcServer := grpc.NewServer()

	// register grpcServerServices
	event_proto.RegisterEventServiceServer(grpcServer, eventServer)
	event_type_proto.RegisterEventTypeServiceServer(grpcServer, eventTypeServer)

	return grpcServer
}
