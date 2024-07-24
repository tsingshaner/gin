package config

import (
	"github.com/tsingshaner/gin/app"
	"github.com/tsingshaner/go-pkg/conf"
	"github.com/tsingshaner/go-pkg/log"
	"github.com/tsingshaner/go-pkg/log/console"
)

type store struct {
	app.Options  `mapstructure:",squash"`
	ConsoleLevel log.Level `mapstructure:"consoleLevel"`
}

var s = conf.New(&store{}, conf.ParseArgs())

func init() {
	if err := s.Load(); err != nil {
		console.Fatal("load config error: %v", err)
	}
	console.SetLevel(Store().ConsoleLevel)
}

func Store() *store {
	return s.Value
}
