package config

import (
	"github.com/spf13/viper"
	"github.com/tsingshaner/gin-starter/pkg/jwt"
)

var JWT = &jwt.Config{}

func setupJWTConfig() {
	requireConfig("jwt.issuer")
	JWT.Expire = 2
	JWT.PublicKeyPath = "./public_key.pem"
	JWT.PrivateKeyPath = "./private_key.pem"
	if err := viper.UnmarshalKey("jwt", &JWT); err != nil {
		panic(err)
	}
}
