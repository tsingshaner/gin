package middleware

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lab-online/pkg/logger"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LogParams struct {
	Method   string
	Path     string
	Status   int
	Latency  time.Duration
	Start    time.Time
	ErrorMsg string
}

type LoggerConfig struct {
	Output           io.Writer               // 文件日志输出
	Logger           func(*LogParams)        // 文件日志记录
	Console          bool                    // 启用控制台日志, 生产环境建议关闭
	ConsoleFormatter func(*LogParams) string // 控制台日志格式化
}

func formatStatus(status int) string {
	switch {
	case status >= 200 && status < 400:
		return logger.Style(fmt.Sprintf("%5d:", status), logger.ColorGreen, logger.FontBold)
	case status >= 400 && status < 500:
		return logger.Style(fmt.Sprintf("%5d:", status), logger.ColorYellow, logger.FontBold)
	case status >= 500:
		return logger.Style(fmt.Sprintf("%5d:", status), logger.ColorRed, logger.FontBold)
	default:
		return logger.Style(fmt.Sprintf("%5d:", status), logger.ColorCyan, logger.FontBold)
	}
}

func defaultLogger(params *LogParams) {
	var level slog.Level
	switch {
	case params.Status >= 200 && params.Status < 400:
		level = slog.LevelInfo
	case params.Status >= 400 && params.Status < 500:
		level = slog.LevelWarn
	case params.Status >= 500:
		level = slog.LevelError
	default:
		level = slog.LevelDebug
	}

	slog.LogAttrs(
		context.Background(),
		level,
		"logger middleware",
		slog.String("method", params.Method),
		slog.String("path", params.Path),
		slog.Int("status", params.Status),
		slog.String("latency", params.Latency.String()),
		slog.String("error", params.ErrorMsg),
	)
}

func defaultConsoleFormatter(params *LogParams) string {
	return fmt.Sprintf("%s %s %s %-50s +%s",
		formatStatus(params.Status),
		logger.Style(params.Start.Format("2006/01/02 15:04:05"), logger.FontDim),
		logger.FormatMethod(params.Method),
		logger.Style(params.Path, logger.ColorGreen),
		params.Latency,
	)
}

func Logger(conf LoggerConfig) gin.HandlerFunc {
	if conf.Logger == nil {
		conf.Logger = defaultLogger
	}
	if conf.Console && conf.ConsoleFormatter == nil {
		conf.ConsoleFormatter = defaultConsoleFormatter
	}
	if conf.Output == nil {
		logger.Error("LoggerConfig.Output is nil")
		panic("")
	}
	slog.SetDefault(slog.New(slog.NewJSONHandler(&lumberjack.Logger{
		Filename:   "logs/server.log",
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     7,
		LocalTime:  true,
		Compress:   true,
	}, nil)))

	return func(c *gin.Context) {
		params := LogParams{
			Method: c.Request.Method,
			Path:   c.Request.URL.Path,
			Start:  time.Now(),
		}
		if c.Request.URL.RawPath != "" {
			params.Path = params.Path + "?" + c.Request.URL.RawPath
		}

		c.Next()

		params.Latency = time.Since(params.Start)
		params.Status = c.Writer.Status()
		params.ErrorMsg = c.Errors.ByType(gin.ErrorTypePrivate).String()

		go func() {
			if params.Latency > time.Minute {
				params.Latency = params.Latency.Truncate(time.Second)
			}
			conf.Logger(&params)
			if conf.Console {
				fmt.Println(conf.ConsoleFormatter(&params))
			}
		}()
	}
}
