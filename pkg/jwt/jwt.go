package jwt

import (
	"crypto/rsa"

	"github.com/golang-jwt/jwt/v5"
)

type Config struct {
	PublicKeyPath  string
	PrivateKeyPath string

	Issuer string
}

type Claims struct {
	UserID string `json:"user_id"`
	Role   uint8  `json:"role"`
	Scoped string `json:"scoped"`
	jwt.RegisteredClaims
}

type JWTAction interface {
	LoadRSA(publicKeyPath string, privateKeyPath string) (*RSAKey, error)
	GenToken(userID string, role uint8) (string, error)
	ParseToken(tokenString string) (*Claims, error)
}

type RSAKey struct {
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
}

type JsonWebToken struct {
	RSA *RSAKey
}

func NewJWT(publicKeyPath string, privateKeyPath string) (JWTAction, error) {
	jsonWebToken := &JsonWebToken{}
	if rsa, err := jsonWebToken.LoadRSA(publicKeyPath, privateKeyPath); err != nil {
		return nil, err
	} else {
		jsonWebToken.RSA = rsa
	}

	return jsonWebToken, nil
}
