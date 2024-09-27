package templatexx

import (
	"github.com/atomeight/distributed-event-collector/component/templatexx/grpc"
	"github.com/atomeight/distributed-event-collector/debug"
	"github.com/latifrons/commongo/safe_viper"
	"github.com/latifrons/latigo"
	"github.com/latifrons/latigo/cron"
	"github.com/latifrons/latigo/grpcserver"
)

type TemplatexxSetup struct {
	DebugFlags   *debug.Flags                  `container:"type"`
	GrpcProvider *grpc.TemplatexxRouteProvider `container:"type"`
}

func (s *TemplatexxSetup) ProvideBootSequence() []latigo.BootSequence {

	grpcServer := &grpcserver.GrpcServer{
		ServiceProvider: s.GrpcProvider,
		Port:            safe_viper.ViperMustGetString("grpc.port"),
		DebugFlags: grpcserver.DebugFlags{
			GRpcDebug:   s.DebugFlags.RpcLog,
			RequestLog:  s.DebugFlags.RequestLog,
			ResponseLog: s.DebugFlags.ResponseLog,
		},
	}

	bs := []latigo.BootSequence{
		{
			Type: latigo.BootTypeComponent,
			Job:  grpcServer,
		},
	}

	crons := []cron.CronJob{}

	for _, cr := range crons {
		bs = append(bs, latigo.BootSequence{
			Type: latigo.BootTypeCron,
			Job:  cr,
		})
	}
	return bs
}
