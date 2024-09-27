package templatexx

import (
	"github.com/atomeight/distributed-event-collector/component/templatexx/grpc"
	"github.com/golobby/container/v3"
)

var Singletons = []interface{}{

	func() *grpc.TemplatexxRouteProvider {
		var c grpc.TemplatexxRouteProvider
		err := container.Fill(&c)
		if err != nil {
			panic(err)
		}
		return &c
	},
	func() *TemplatexxSetup {
		var c TemplatexxSetup
		err := container.Fill(&c)
		if err != nil {
			panic(err)
		}
		return &c
	},
}
