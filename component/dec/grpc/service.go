package grpc

import (
	"context"
	"github.com/atomeight/distributed-event-collector/pbgo/dec"
	"github.com/atomeight/distributed-event-collector/service"
)

type DecService struct {
	dec.DecServiceServer
	SpanCollector *service.SpanCollector `container:"type"`
}

func (d DecService) GetEventFlow(ctx context.Context, request *dec.GetEventFlowRequest) (*dec.GetEventFlowResponse, error) {
	return d.SpanCollector.GetEventFlow(request.Owner)
}

func (d DecService) GetEventStatistics(ctx context.Context, request *dec.GetEventStatisticsRequest) (*dec.GetEventStatisticsResponse, error) {
	return d.SpanCollector.GetEventStatistics(request.Samples)
}

func (d DecService) Report(ctx context.Context, request *dec.ReportRequest) (resp *dec.ReportResponse, err error) {
	err = d.SpanCollector.Report(request)
	if err != nil {
		return
	}
	resp = &dec.ReportResponse{}
	return
}

func (d DecService) mustEmbedUnimplementedDecServiceServer() {
	//TODO implement me
	panic("implement me")
}
