package domain

import (
	"github.com/lab-online/internal/user/entity"
	"github.com/lab-online/pkg/jwt"
)

func (*Domain) Login(user *entity.User) (string, error) {
	token := jwt.GenToken(user.UserID, user.Role)

	return token, nil
}
