package grpc

import (
	"github.com/latifrons/latigo/grpcserver"
)

type DECRouteProvider struct {
}

func (r *DECRouteProvider) ProvideAllServices() []grpcserver.GrpcService {

	return []grpcserver.GrpcService{}
}
