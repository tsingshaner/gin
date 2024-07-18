package config

import (
	"github.com/spf13/viper"
	"github.com/tsingshaner/gin/app"
	"github.com/tsingshaner/gin/pkg/middleware"
	"github.com/tsingshaner/go-pkg/log"
	"github.com/tsingshaner/go-pkg/log/console"
)

type store struct {
	ConsoleLevel log.Level `json:"consoleLevel" yaml:"consoleLevel" toml:"consoleLevel"`

	Cors   *middleware.CorsOptions `json:"cors" yaml:"cors" toml:"cors"`
	Server *app.ServerOptions      `json:"server" yaml:"server" toml:"server"`
	Logger *app.LoggerOptions      `json:"logger" yaml:"logger" toml:"logger"`
}

var s *store = nil

func Store() *store {
	return s
}

func NewStore() *store {
	s := &store{}

	if err := viper.Unmarshal(s); err != nil {
		console.Fatal("unmarshal config error: %s", err)
	}

	return s
}
