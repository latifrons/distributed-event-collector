package grpc

import (
	"github.com/latifrons/latigo/grpcserver"
)

type TemplatexxRouteProvider struct {
}

func (r *TemplatexxRouteProvider) ProvideAllServices() []grpcserver.GrpcService {

	return []grpcserver.GrpcService{}
}
