package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lab-online/config"
)

func (j *JsonWebToken) GenToken(userID string, role uint8) (string, error) {
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

	return jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(j.RSA.PrivateKey)
}
