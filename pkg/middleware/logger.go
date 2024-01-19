package middleware

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lab-online/pkg/color"
	"github.com/lab-online/pkg/logger"
)

type LogParams struct {
	Method   string
	Path     string
	Status   int
	Latency  time.Duration
	Start    time.Time
	ErrorMsg string
	Level    slog.Level
}

type LoggerConfig struct {
	Level            *int                     // 日志级别 0-199: -4, 200-399: 0, 400-499: 4, 500-: 8
	FileLogger       func(*LogParams)        // 文件日志记录
	Console          bool                    // 启用控制台日志, 生产环境建议关闭
	ConsoleFormatter func(*LogParams) string // 控制台日志格式化
}

func Logger(conf LoggerConfig) gin.HandlerFunc {
	if conf.FileLogger == nil {
		conf.FileLogger = defaultFileLogger
	}
	if conf.Console && conf.ConsoleFormatter == nil {
		conf.ConsoleFormatter = defaultConsoleFormatter
	}

	return func(c *gin.Context) {
		params := LogParams{
			Method: c.Request.Method,
			Path:   c.Request.URL.Path,
			Start:  time.Now(),
		}
		if c.Request.URL.RawQuery != "" {
			params.Path = params.Path + "?" + c.Request.URL.RawQuery
		}

		c.Next()

		params.Status = c.Writer.Status()
		params.Level = getLevelByStatus(params.Status)
		if level := int(params.Level); *conf.Level > level {
			return
		}
		params.Latency = time.Since(params.Start)
		params.ErrorMsg = c.Errors.ByType(gin.ErrorTypePrivate).String()

		go func(params *LogParams) {
			if params.Latency > time.Minute {
				params.Latency = params.Latency.Truncate(time.Second)
			}
			conf.FileLogger(params)
			if conf.Console {
				fmt.Println(conf.ConsoleFormatter(params))
			}
		}(&params)
	}
}

func getLevelByStatus(status int) slog.Level {
	switch {
	case status >= 200 && status < 400:
		return slog.LevelInfo
	case status >= 400 && status < 500:
		return slog.LevelWarn
	case status >= 500:
		return slog.LevelError
	default:
		return slog.LevelDebug
	}
}

func formatStatus(status int) string {
	formattedStatus := color.Bold(fmt.Sprintf("%5d:", status))

	switch {
	case status >= 200 && status < 400:
		return color.Green(formattedStatus)
	case status >= 400 && status < 500:
		return color.Yellow(formattedStatus)
	case status >= 500:
		return color.Red(formattedStatus)
	default:
		return color.Cyan(formattedStatus)
	}
}

func defaultFileLogger(params *LogParams) {
	slog.LogAttrs(
		context.Background(),
		params.Level,
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
		color.Dim(params.Start.Format("2006/01/02 15:04:05")),
		logger.FormatMethod(params.Method),
		color.Green(params.Path),
		params.Latency,
	)
}
