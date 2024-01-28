package config

import (
	"github.com/lab-online/pkg/jwt"
	"github.com/spf13/viper"
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
