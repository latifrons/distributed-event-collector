package grpc

import (
	"github.com/latifrons/distributed-event-collector/pbgo/dec"
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
