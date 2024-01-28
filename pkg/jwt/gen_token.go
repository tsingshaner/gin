package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (j *JsonWebToken) GenToken(userID string, role uint8) (string, error) {
	now := time.Now()

	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.config.Issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(j.config.Expire) * time.Hour)),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(j.rsa.PrivateKey)
}
