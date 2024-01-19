package config

import (
	"fmt"

	"github.com/lab-online/pkg/color"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	pflag.String("config_filename", "app", "config file name")
	pflag.String("config_type", "yaml", "config file type")
	pflag.String("config_dir", "./config", "config file path")

	filename := pflag.Lookup("config_filename").Value.String()
	config_type := pflag.Lookup("config_type").Value.String()
	dir := pflag.Lookup("config_dir").Value.String()

	setupConfig(filename, config_type, dir)
}

type DatabaseConfig struct {
	Postgres string // postgres URI
	Redis    string // redis URI
}

var Database DatabaseConfig

func unmarshalConfig() {
	viper.AutomaticEnv()
	viper.WatchConfig()

	setupServerConfig()

	requireConfig("database.postgres")
	if err := viper.UnmarshalKey("database", &Database); err != nil {
		panic(err)
	}

	setupLoggerConfig()
	setupJWTConfig()
}

func setupConfig(filename string, config_type string, dir string) {
	viper.SetConfigName(filename)
	viper.SetConfigType(config_type)
	viper.AddConfigPath(dir)

	if err := viper.ReadInConfig(); err != nil {
		errPrefix := color.Red("error:")

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
			panic("config field `" +color.Red(k) + "` is required")
		}
	}
}
