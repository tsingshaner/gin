package jwt

func (*JsonWebToken) ParseToken(password string) (*Claims, error) {
	return &Claims{}, nil
}
