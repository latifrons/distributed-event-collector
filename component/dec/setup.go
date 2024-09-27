package dec

import (
	"github.com/atomeight/distributed-event-collector/component/dec/grpc"
	"github.com/atomeight/distributed-event-collector/debug"
	"github.com/latifrons/commongo/safe_viper"
	"github.com/latifrons/latigo"
	"github.com/latifrons/latigo/cron"
	"github.com/latifrons/latigo/grpcserver"
)

type DECSetup struct {
	DebugFlags   *debug.Flags           `container:"type"`
	GrpcProvider *grpc.DECRouteProvider `container:"type"`
}

func (s *DECSetup) ProvideBootSequence() []latigo.BootSequence {

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
