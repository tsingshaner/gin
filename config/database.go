package config

import "github.com/spf13/viper"

type DatabaseConfig struct {
	Postgres string // postgres URI
	Redis    string // redis URI
}

var Database DatabaseConfig

func setupDatabaseConfig() {
	requireConfig("database.postgres")
	if err := viper.UnmarshalKey("database", &Database); err != nil {
		panic(err)
	}
}
