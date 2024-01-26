package domain

import (
	"github.com/lab-online/internal/user/entity"
)

func (d *Domain) Login(user *entity.User) (string, error) {
	token, err := d.jwt.GenToken(user.UserID, user.Role)

	return token, err
}
