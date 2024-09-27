package dec

import (
	"github.com/golobby/container/v3"
	"github.com/latifrons/distributed-event-collector/component/dec/grpc"
	"github.com/latifrons/distributed-event-collector/service"
	"github.com/spf13/viper"
)

var Singletons = []interface{}{

	func() *grpc.DECRouteProvider {
		var c grpc.DECRouteProvider
		err := container.Fill(&c)
		if err != nil {
			panic(err)
		}
		return &c
	},
	func() *DECSetup {
		var c DECSetup
		err := container.Fill(&c)
		if err != nil {
			panic(err)
		}
		return &c
	},
	func() *grpc.DecService {
		var c grpc.DecService
		err := container.Fill(&c)
		if err != nil {
			panic(err)
		}
		return &c
	},
	func() *service.SpanCollector {
		var c service.SpanCollector
		err := container.Fill(&c)
		if err != nil {
			panic(err)
		}
		c.CacheSize = viper.GetInt("general.cache_size")
		c.InitDefault()
		return &c
	},
}
