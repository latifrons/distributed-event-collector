package grpc

import (
	"github.com/atomeight/distributed-event-collector/pbgo/dec"
	"github.com/latifrons/latigo/grpcserver"
)

type DECRouteProvider struct {
	DecService *DecService `container:"type"`
}

func (r *DECRouteProvider) ProvideAllServices() []grpcserver.GrpcService {

	return []grpcserver.GrpcService{
		{Desc: &dec.DecService_ServiceDesc, SS: r.DecService},
	}
}
