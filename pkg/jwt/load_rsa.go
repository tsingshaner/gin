package jwt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"

	"github.com/lab-online/pkg/color"
)

func (*JsonWebToken) LoadRSA(publicKeyPath string, privateKeyPath string) (*RSAKey, error) {
	rsaKey := &RSAKey{}

	if pub, err := getKey[rsa.PublicKey](publicKeyPath, x509.ParsePKIXPublicKey); err != nil {
		return nil, err
	} else {
		rsaKey.PublicKey = pub
		jwtLogger(color.ColorGreen, "RSA public key load success")
	}

	if pri, err := getKey[rsa.PrivateKey](privateKeyPath, x509.ParsePKCS8PrivateKey); err != nil {
		return nil, err
	} else {
		rsaKey.PrivateKey = pri
		jwtLogger(color.ColorGreen, "RSA private key load success")
	}

	return rsaKey, nil
}

func getKey[T any](path string, parser func([]byte) (any, error)) (*T, error) {
	if file, err := os.ReadFile(path); err != nil {
		jwtError("pem file load failed: ", path)
		return nil, err
	} else {
		block, _ := pem.Decode(file)
		if block == nil {
			err := errors.New("jwt failed to parse PEM block containing the key")
			jwtError(err.Error())
			return nil, err
		}
		if pub, err := parser(block.Bytes); err != nil {
			jwtError("parse rsa key failed")
			return nil, err
		} else {
			return pub.(*T), nil
		}
	}
}

func jwtError(msg ...any) {
	jwtLogger(color.ColorRed, msg)
}

func jwtLogger(style int, msg ...any) {
	color.Log("jwt", style, msg)
}
