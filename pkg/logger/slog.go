package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/tsingshaner/go-pkg/color"
	"gopkg.in/natefinch/lumberjack.v2"
)

type DefaultIOWrite struct {
	FileEnable bool
	FileWriter io.Writer

	ConsoleEnable bool
	ConsoleWriter io.Writer
}

func (io *DefaultIOWrite) Write(p []byte) (int, error) {
	if io.ConsoleEnable {
		io.ConsoleWriter.Write(p)
	}
	if io.FileEnable {
		io.FileWriter.Write(p)
	}

	return 0, nil
}

type LoggerConfig struct {
	FileEnable    bool
	FileWriter    io.Writer
	ConsoleEnable bool
	ConsoleWriter io.Writer

	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	LocalTime  bool
	Compress   bool
}

func DefaultWriter(conf LoggerConfig) *DefaultIOWrite {
	ioWriter := &DefaultIOWrite{
		FileEnable: conf.FileEnable,
		FileWriter: conf.FileWriter,
	}

	if ioWriter.ConsoleEnable && conf.ConsoleWriter == nil {
		ioWriter.ConsoleWriter = os.Stdout
	}
	if ioWriter.FileEnable && conf.FileWriter == nil {
		logger := &lumberjack.Logger{
			Filename:   "logs/server.log",
			MaxSize:    10,
			MaxBackups: 3,
			MaxAge:     7,
			LocalTime:  conf.LocalTime,
			Compress:   conf.Compress,
		}

		if conf.Filename != "" {
			logger.Filename = conf.Filename
		}
		if conf.MaxSize != 0 {
			logger.MaxSize = conf.MaxSize
		}
		if conf.MaxBackups != 0 {
			logger.MaxBackups = conf.MaxBackups
		}
		if conf.MaxAge != 0 {
			logger.MaxAge = conf.MaxAge
		}

		ioWriter.FileWriter = logger
		return ioWriter
	} else {
		ioWriter.FileWriter = conf.FileWriter
		return ioWriter
	}
}

func InitSlog(w io.Writer, level int) {
	slogLevel := slog.LevelWarn
	switch level {
	case -4:
		slogLevel = slog.LevelDebug
	case 0:
		slogLevel = slog.LevelInfo
	case 4:
		slogLevel = slog.LevelWarn
	case 8:
		slogLevel = slog.LevelError
	}

	prefix := color.UnsafeGreen(color.UnsafeBold(" slog:"))
	fmt.Println(prefix, "slog min level", slogLevel)

	slog.SetDefault(slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{
		AddSource: true,
		Level:     slogLevel,
	})))
}
