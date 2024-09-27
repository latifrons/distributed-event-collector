package core

import (
	"github.com/golobby/container/v3"
	"github.com/latifrons/distributed-event-collector/cache"
	"github.com/latifrons/distributed-event-collector/debug"
	"github.com/latifrons/distributed-event-collector/tools"
	"github.com/spf13/viper"
)

type SingletonBatch []interface{}

func BuildDependencies(additional []interface{}) (err error) {
	singletons := []interface{}{
		func() *cache.RedisCache {
			redisCache := &cache.RedisCache{
				Address:    tools.ViperMustGetString("redis.address"),
				Password:   viper.GetString("redis.password"),
				Db:         tools.ViperMustGetInt("redis.db"),
				RootFolder: viper.GetString("redis.root_folder"),
			}
			redisCache.Init()
			return redisCache
		},
		func() *debug.Flags {
			debugFlags := &debug.Flags{
				ReturnDetailError: viper.GetBool("debug.return_detail_error"),
				DbLog:             viper.GetBool("debug.db_log"),
				RequestLog:        viper.GetBool("debug.request_log"),
				ResponseLog:       viper.GetBool("debug.response_log"),
				RpcLog:            viper.GetBool("debug.rpc_log"),
				LogLevel:          viper.GetString("debug.log_level"),
			}
			return debugFlags
		},
	}

	singletons = append(singletons, additional...)

	for _, singleton := range singletons {
		err = container.SingletonLazy(singleton)
		if err != nil {
			return
		}
	}
	return
}
