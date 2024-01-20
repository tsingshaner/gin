package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lab-online/config"
)

func GenToken(userID string, role uint8) string {
	now := time.Now()

	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.JWT.Issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(config.JWT.Expire) * time.Hour)),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(RSA.PrivateKey)
	if err != nil {
		panic(err)
	}

	return token
}
