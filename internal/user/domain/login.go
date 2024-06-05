package domain

import (
	"github.com/tsingshaner/gin-starter/internal/user/constant"
	"github.com/tsingshaner/gin-starter/internal/user/entity"
)

func (d *Domain) Login(user entity.UserEntity) (string, error) {
	if registeredUser, err := d.repository.FindByUserID(user.GetUserID()); err != nil {
		return "", newCustomError(constant.UserNotExists)
	} else if !registeredUser.ComparePassword(user.GetPassword()) {
		return "", newCustomError(constant.UserPasswordError)
	} else {
		token, err := d.jwt.GenToken(registeredUser.GetUserID(), registeredUser.GetRole())

		if err != nil {
			return "", newCustomError(constant.TokenGenError)
		}
		return token, nil
	}
}
