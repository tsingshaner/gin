package config

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/tsingshaner/go-pkg/color"
)

func init() {
	var configFile string
	pflag.StringVarP(&configFile, "config", "c", "./config/app.yaml", "config file path")
	pflag.Parse()

	setupConfig(configFile)
}

func unmarshalConfig() {
	viper.AutomaticEnv()
	viper.WatchConfig()

	setupServerConfig()
	setupCORSConfig()
	setupDatabaseConfig()
	setupLoggerConfig()
	setupJWTConfig()
}

func setupConfig(configFile string) {
	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(color.UnsafeBold(color.UnsafeRed("error:")), err)
		panic(err)
	}
	fmt.Println(color.UnsafeBold(color.UnsafeGreen(" info:")), "using config file", viper.ConfigFileUsed())

	unmarshalConfig()
}

func requireConfig(key ...string) {
	for _, k := range key {
		if !viper.IsSet(k) {
			panic("config field `" + color.UnsafeRed(k) + "` is required")
		}
	}
}
