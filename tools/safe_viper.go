package tools

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func ViperMustGetString(key string) string {
	if !viper.IsSet(key) || viper.GetString(key) == "" {
		log.Fatal().Str("key", key).Msg("config missing")
	}
	return viper.GetString(key)
}
func ViperMustGetInt(key string) int {
	if !viper.IsSet(key) || viper.GetString(key) == "" {
		log.Fatal().Str("key", key).Msg("config missing")
	}
	return viper.GetInt(key)
}
func ViperMustGetBool(key string) bool {
	if !viper.IsSet(key) || viper.GetString(key) == "" {
		log.Fatal().Str("key", key).Msg("config missing")
	}
	return viper.GetBool(key)
}
