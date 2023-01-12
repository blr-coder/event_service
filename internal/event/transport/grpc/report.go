package grpc

import (
	"context"
	"event_service/api/grpc/report_proto"
	"event_service/internal/event/usecases"
	"event_service/internal/event/usecases/usecase_models"
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
	grpcReportFilter *report_proto.ReportListFilter,
) (*report_proto.Reports, error) {

	reports, err := s.useCase.Report.List(ctx, decodeGrpcReportFilter(grpcReportFilter))
	if err != nil {
		// TODO: Add decodeError func
		return nil, err
	}

	return ucReportsToGrpcReports(reports), nil
}

func decodeGrpcReportFilter(grpcReportFilter *report_proto.ReportListFilter) *usecase_models.ReportFilter {
	return &usecase_models.ReportFilter{
		From: grpcReportFilter.From.AsTime(),
		To:   grpcReportFilter.To.AsTime(),
	}
}

func ucReportsToGrpcReports(ucReports *usecase_models.Reports) *report_proto.Reports {

	return &report_proto.Reports{}
}
