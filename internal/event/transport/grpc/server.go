package grpc

import (
	"event_service/api/grpc/event_proto"
	"event_service/api/grpc/event_type_proto"
	"event_service/api/grpc/report_proto"
	"google.golang.org/grpc"
)

func NewGRPCServer(eventServer *EventServiceServer, eventTypeServer *EventTypeServiceServer, reportServer *ReportServiceServer) *grpc.Server {

	grpcServer := grpc.NewServer()

	// register grpcServerServices
	event_proto.RegisterEventServiceServer(grpcServer, eventServer)
	event_type_proto.RegisterEventTypeServiceServer(grpcServer, eventTypeServer)
	report_proto.RegisterReportServiceServer(grpcServer, reportServer)

	return grpcServer
}
