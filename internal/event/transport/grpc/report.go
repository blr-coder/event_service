package grpc

import (
	"context"
	"event_service/api/grpc/report_proto"
	"event_service/internal/event/usecases"
	"event_service/internal/event/usecases/usecase_models"
	"google.golang.org/protobuf/types/known/timestamppb"
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

	ucFilter := &usecase_models.ReportFilter{
		From:    grpcReportFilter.From.AsTime(),
		To:      grpcReportFilter.To.AsTime(),
		GroupBy: GroupByFromGrpc[grpcReportFilter.ReportByPeriod],
	}

	if grpcReportFilter.UserId != nil {
		ucFilter.UserID = &grpcReportFilter.UserId.Value
	}

	return ucFilter
}

const basicCurrency = "EUR"

func ucReportsToGrpcReports(ucReports usecase_models.Reports) *report_proto.Reports {

	var reports []*report_proto.Report

	for _, report := range ucReports {
		reports = append(reports, &report_proto.Report{
			Date:      timestamppb.New(report.Date),
			EventType: report.Type,
			Cost: &report_proto.Cost{
				// TODO: Add getting amount from currency_api
				Amount:   report.Sum,
				Currency: basicCurrency,
			},
			Count: report.Count,
		})
	}

	return &report_proto.Reports{
		Reports: reports,
	}
}

var GroupByFromGrpc = map[report_proto.REPORT_BY]usecase_models.ReportGroupBy{
	report_proto.REPORT_BY_HOUR:  usecase_models.ReportGroupByHour,
	report_proto.REPORT_BY_DAY:   usecase_models.ReportGroupByDay,
	report_proto.REPORT_BY_WEEK:  usecase_models.ReportGroupByWeek,
	report_proto.REPORT_BY_MONTH: usecase_models.ReportGroupByMonth,
}
