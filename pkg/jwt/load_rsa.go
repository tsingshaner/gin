package jwt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"

	"github.com/lab-online/config"
	"github.com/lab-online/pkg/color"
)

type RSAKey struct {
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
}

var RSA RSAKey

func init() {
	if pubKey, err := os.ReadFile(config.JWT.PublicKeyPath); err != nil {
		jwtError(err, "JWT public key load failed: ", config.JWT.PublicKeyPath)
	} else {
		parsePublicKey(pubKey)
	}

	if priKey, err := os.ReadFile(config.JWT.PrivateKeyPath); err != nil {
		jwtError(err, "JWT private key load failed: ", config.JWT.PrivateKeyPath)
	} else {
		parsePrivateKey(priKey)
	}

	jwtLogger(color.ColorGreen, "JWT RSA key load success")
}

func parsePublicKey(pubKey []byte) {
	block, _ := pem.Decode(pubKey)
	if block == nil {
		panic("jwt failed to parse PEM block containing the public key")
	}
	if pub, err := x509.ParsePKIXPublicKey(block.Bytes); err != nil {
		jwtError(err, "parse rsa public key failed")
	} else {
		RSA.PublicKey = pub.(*rsa.PublicKey)
	}
}

func parsePrivateKey(priKey []byte) {
	block, _ := pem.Decode(priKey)
	if block == nil {
		panic("jwt failed to parse PEM block containing the private key")
	}
	if pri, err := x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
		jwtError(err, "parse rsa private key failed")
	} else {
		RSA.PrivateKey = pri.(*rsa.PrivateKey)
	}
}

func jwtError(err error, msg ...any) {
	jwtLogger(color.ColorRed, msg)
	panic(err)
}

func jwtLogger(style int, msg ...any) {
	color.Log("jwt", style, msg)
}
