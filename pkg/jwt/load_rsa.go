package jwt

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func (j *JsonWebToken) LoadRSA() error {
	if key, err := parseKey(j.config.PublicKeyPath, jwt.ParseRSAPublicKeyFromPEM); err != nil {
		return err
	} else {
		j.rsa.PublicKey = key
		jwtInfo("RSA public key load success")
	}

	if key, err := parseKey(j.config.PrivateKeyPath, jwt.ParseRSAPrivateKeyFromPEM); err != nil {
		return err
	} else {
		j.rsa.PrivateKey = key
		jwtInfo("RSA private key load success")
	}

	return nil
}

func parseKey[T any](path string, parser func([]byte) (*T, error)) (*T, error) {
	if file, err := os.ReadFile(path); err != nil {
		jwtError("pem file load failed: ", path)
		return nil, err
	} else {
		if key, err := parser(file); err != nil {
			jwtError("pem file parse failed: ", path)
			return nil, err
		} else {
			return key, nil
		}
	}
}
