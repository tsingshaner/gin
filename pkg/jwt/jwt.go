package jwt

import (
	"crypto/rsa"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tsingshaner/gin-starter/pkg/color"
)

type Config struct {
	PublicKeyPath  string
	PrivateKeyPath string

	Issuer string
	Expire int8
}

type Claims struct {
	UserID string `json:"uid"`
	Role   uint8  `json:"role"`
	jwt.RegisteredClaims
}

type JWTAction interface {
	LoadRSA() error
	GenToken(userID string, role uint8) (string, error)
	ParseToken(token string) (*Claims, error)
}

type RSAKey struct {
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
}

type JsonWebToken struct {
	config *Config
	rsa    *RSAKey
}

func New(config *Config) (JWTAction, error) {
	jwtInfo("JWT Issuer", config.Issuer)
	jwtInfo("JWT Expire", config.Expire, "hours")

	jsonWebToken := &JsonWebToken{config, &RSAKey{}}
	if err := jsonWebToken.LoadRSA(); err != nil {
		return nil, err
	}

	return jsonWebToken, nil
}

func jwtInfo(msg ...any) {
	jwtLogger(color.ColorGreen, msg...)
}

func jwtError(msg ...any) {
	jwtLogger(color.ColorRed, msg...)
}

func jwtLogger(style int, msg ...any) {
	color.Log("jwt", style, msg)
}
