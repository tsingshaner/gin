package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/tsingshaner/gin-starter/pkg/logger"
)

type LoggerConfig struct {
	ConsoleLevel int    // 控制台输出等级 位掩码 1: debug, 2: info, 4: warn, 8: error
	SlogLevel    int    // slog 输出等级 -4: debug, 0: info, 4: warn, 8: error
	ConsoleHttp  bool   // 控制台输出http日志
	HttpLevel    int    // logger 中间件输出等级，优先级低于slog与console配置 -4: debug, 0: info, 4: warn, 8: error
	Path         string // 日志文件夹路径 默认 ./logs
	FileName     string // 日志文件名 默认 server.log
	MaxSize      int    // max size per log file
	MaxBackups   int    // max backups per log file
	MaxAge       int    // max age per log file
}

var Logger LoggerConfig

func setupLoggerConfig() {
	Logger.ConsoleLevel = 15
	Logger.SlogLevel = -4
	Logger.ConsoleHttp = false
	Logger.HttpLevel = -4
	Logger.Path = "./logs"
	Logger.FileName = "app.log"
	Logger.MaxSize = 500
	Logger.MaxBackups = 3
	Logger.MaxAge = 7

	if err := viper.UnmarshalKey("logger", &Logger); err != nil {
		panic(err)
	}

	setupLogger()

	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := viper.UnmarshalKey("logger", &Logger); err != nil {
			panic(err)
		}
		setupLogger()
	})
}

func setupLogger() {
	setupLog(Logger.ConsoleLevel)
	setupSlog(Logger.SlogLevel)
}

func setupLog(level int) {
	logger.SetLevel(level)
}

func setupSlog(level int) {
	logger.InitSlog(logger.DefaultWriter(
		logger.LoggerConfig{
			FileEnable:    true,
			ConsoleEnable: true,

			Filename:   fmt.Sprintf("%s/%s", Logger.Path, Logger.FileName),
			MaxSize:    Logger.MaxSize,
			MaxBackups: Logger.MaxBackups,
			MaxAge:     Logger.MaxAge,
		},
	), level)
}
