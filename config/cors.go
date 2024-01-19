package config

import "github.com/spf13/viper"

type CORSConfig struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials bool
	MaxAge           int
}

var CORS CORSConfig

func setupCORSConfig() {
	CORS.AllowOrigins = []string{"*"}
	CORS.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	CORS.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	CORS.AllowCredentials = true
	CORS.MaxAge = 60
	if err := viper.UnmarshalKey("cors", &CORS); err != nil {
		panic(err)
	}
}
