package grpc

import (
	"context"
	"event_service/api/grpc/report_proto"
	"event_service/internal/event/usecases"
	"event_service/internal/event/usecases/usecase_models"
	"time"
)

type ReportServiceServer struct {
	report_proto.UnimplementedReportServiceServer
	useCase *usecases.UseCase
}

func NewReportServiceServer(useCase *usecases.UseCase) *ReportServiceServer {
	return &ReportServiceServer{useCase: useCase}
}

func (s *ReportServiceServer) ListByFilter(
	ctx context.Context,
	grpcReportFilter *report_proto.ReportListFilter) (*report_proto.Reports, error) {

	_, err := s.useCase.Report.List(ctx, &usecase_models.ReportFilter{
		From: time.Time{},
		To:   time.Time{},
	})
	if err != nil {
		// TODO: Add decodeError func
		return nil, err
	}

	return nil, nil
}
