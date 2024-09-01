package app

import (
	"time"

	"github.com/tsingshaner/gin/middleware"
	"github.com/tsingshaner/gin/swagger"
	"github.com/tsingshaner/go-pkg/jwt"
	"github.com/tsingshaner/go-pkg/log"
)

type (
	Options struct {
		Cors     *middleware.CorsOptions `mapstructure:"cors"`
		JWT      *jwt.Options            `mapstructure:"jwt"`
		Logger   *LoggerOptions          `mapstructure:"logger"`
		Postgres *Postgres               `mapstructure:"postgres"`
		Server   *ServerOptions          `mapstructure:"server"`
		Swagger  *swagger.Options        `mapstructure:"swagger"`
	}

	FileWriterOptions struct {
		Enable    bool   `mapstructure:"enable"`
		Directory string `mapstructure:"director"`
		Filename  string `mapstructure:"filename"`
	}

	StdoutWriterOption struct {
		Enable bool `mapstructure:"enable"`
	}

	Postgres struct {
		Host     string `mapstructure:"host"`
		Port     int64  `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Database string `mapstructure:"database"`
		SSLMode  string `mapstructure:"sslMode"`
		TimeZone string `mapstructure:"timezone"`
	}

	GormLoggerOptions struct {
		LogLevel                  log.Level     `mapstructure:"level"`
		SlowThreshold             time.Duration `mapstructure:"slowThreshold"`
		IgnoreRecordNotFoundError bool          `mapstructure:"ignoreRecordNotFoundError"`
		ParameterizedQueries      bool          `mapstructure:"parameterizedQueries"`
	}

	LoggerOptions struct {
		Level        log.Level         `mapstructure:"level"`
		GormLogger   GormLoggerOptions `mapstructure:"gorm"`
		FileWriter   FileWriterOptions `mapstructure:"fileWriter"`
		StdoutWriter FileWriterOptions `mapstructure:"stdoutWriter"`
	}

	ServerOptions struct {
		Host string `mapstructure:"host"`
		Port int64  `mapstructure:"port"`
		Base string `mapstructure:"base"`
		Mode string `mapstructure:"mode"`

		RequestIdHeader string        `mapstructure:"requestIdHeader"`
		ReadTimeout     time.Duration `mapstructure:"readTimeout"`
		WriteTimeout    time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderBytes  int           `mapstructure:"maxHeaderBytes"`
	}
)
