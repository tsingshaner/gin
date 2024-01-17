package config

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Prefix string // api prefix
	Host   string // server host
	Port   int    // server port
}

var Server ServerConfig

type DatabaseConfig struct {
	Postgres string // postgres URI
	Redis    string // redis URI
}

var Database DatabaseConfig

type LoggerConfig struct {
	Level      string // http, debug, info, warn, error
	Path       string // log file path
	FileName   string // log file name
	MaxSize    int    // max size per log file
	MaxBackups int    // max backups per log file
}

var Logger LoggerConfig

type JsonWebTokenConfig struct {
	Secret string // jwt secret
	Expire int 	  // jwt expire
}

var JsonWebToken JsonWebTokenConfig

func unmarshalConfig() {
	viper.AutomaticEnv()

	optionalConfig(
		&optional{"server.prefix", "/api/v1"},
		&optional{"server.host", "127.0.0.1"},
		&optional{"server.port", 8080},
	)
	if err := viper.UnmarshalKey("server", &Server); err != nil {
		panic(err)
	}

	requireConfig("database.postgres")
	if err := viper.UnmarshalKey("database", &Database); err != nil {
		panic(err)
	}

	optionalConfig(
		&optional{"logger.level", "debug"},
		&optional{"logger.path", "./logs"},
		&optional{"logger.filename", "app.log"},
		&optional{"logger.max_size", 500},
		&optional{"logger.max_backups", 3},
	)
	if err := viper.UnmarshalKey("logger", &Logger); err != nil {
		panic(err)
	}

	requireConfig("jwt.secret")
	optionalConfig(&optional{"jwt.expire", 3600})
	if err := viper.UnmarshalKey("jwt", &JsonWebToken); err != nil {
		panic(err)
	}
}

func init() {
	pflag.String("config_filename", "app", "config file name")
	pflag.String("config_type", "yaml", "config file type")
	pflag.String("config_dir", "./config", "config file path")

	filename := pflag.Lookup("config_filename").Value.String()
	config_type := pflag.Lookup("config_type").Value.String()
	dir := pflag.Lookup("config_dir").Value.String()

	setupConfig(filename, config_type, dir)
}

func setupConfig(filename string, config_type string, dir string) {
	viper.SetConfigName(filename)
	viper.SetConfigType(config_type)
	viper.AddConfigPath(dir)

	if err := viper.ReadInConfig(); err != nil {
		errPrefix := "\x1b[1;31merror:\x1b[0m"

		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Printf(
				"%s config file not found: %s%s.%s\n",
				errPrefix, dir, filename, config_type,
			)
		} else {
			fmt.Println(errPrefix + " read config file error")
		}
		panic(err)
	}

	unmarshalConfig()
}

func requireConfig(key ...string) {
	for _, k := range key {
		if !viper.IsSet(k) {
			panic("config field `\x1b[1;31m" + k + "\x1b[0m` is required")
		}
	}
}

type optional struct {
	key   string
	Value any
}

func optionalConfig(optionalConfigs ...*optional) {
	for _, item := range optionalConfigs {
		viper.SetDefault(item.key, item.Value)
	}
}
