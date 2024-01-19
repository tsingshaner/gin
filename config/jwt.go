package config

import "github.com/spf13/viper"

type JsonWebTokenConfig struct {
	Issuer         string // jwt issuer
	Secret         string // jwt secret
	Expire         int    // jwt expire
	PublicKeyPath  string // jwt public key path
	PrivateKeyPath string // jwt private key path
}

var JWT JsonWebTokenConfig

func setupJWTConfig() {
	requireConfig("jwt.secret", "jwt.issuer")
	JWT.Expire = 2
	JWT.PublicKeyPath = "./public_key.pem"
	JWT.PrivateKeyPath = "./private_key.pem"
	if err := viper.UnmarshalKey("jwt", &JWT); err != nil {
		panic(err)
	}
}
