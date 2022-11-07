package config

import "github.com/spf13/viper"

func ReadConfig() {
	InitConfig()
	viper.Get("server")
}
