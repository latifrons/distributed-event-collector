package dec

import (
	"github.com/atomeight/distributed-event-collector/component/dec/grpc"
	"github.com/golobby/container/v3"
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
}
