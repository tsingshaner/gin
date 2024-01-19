package config

import (
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Prefix string // api prefix
	Host   string // server host
	Port   int    // server port
	Mode   string // gin mode
}

var Server ServerConfig

func setupServerConfig() {
	Server.Prefix = "/api/v1"
	Server.Host = "127.0.0.1"
	Server.Port = 8080
	requireConfig("server.mode")
	if err := viper.UnmarshalKey("server", &Server); err != nil {
		panic(err)
	}
}
