package dec

import (
	"github.com/latifrons/commongo/safe_viper"
	"github.com/latifrons/distributed-event-collector/component/dec/grpc"
	"github.com/latifrons/distributed-event-collector/debug"
	"github.com/latifrons/distributed-event-collector/service"
	"github.com/latifrons/latigo"
	"github.com/latifrons/latigo/cron"
	"github.com/latifrons/latigo/grpcserver"
	"time"
)

type DECSetup struct {
	DebugFlags    *debug.Flags           `container:"type"`
	GrpcProvider  *grpc.DECRouteProvider `container:"type"`
	SpanCollector *service.SpanCollector `container:"type"`
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

	crons := []cron.CronJob{
		{
			Name:             "Dump",
			Type:             cron.CronJobTypeInterval,
			Cron:             "",
			WaitForSchedule:  false,
			DisableSingleton: false,
			Interval:         time.Second * 10,
			Function:         s.SpanCollector.Dump,
			Params:           nil,
		},
	}

	for _, cr := range crons {
		bs = append(bs, latigo.BootSequence{
			Type: latigo.BootTypeCron,
			Job:  cr,
		})
	}
	return bs
}
