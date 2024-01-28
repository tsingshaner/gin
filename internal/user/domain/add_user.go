package domain

import (
	"github.com/lab-online/internal/user/constant"
	"github.com/lab-online/internal/user/entity"
)

func (d *Domain) AddUser(user entity.UserEntity) error {
	if exist, err := d.repository.CheckUserExists(user.GetUserID()); err != nil {
		return newCustomError(constant.DBError)
	} else if exist {
		return newCustomError(constant.UserAlreadyExists)
	}

	if err := user.HashPassword(); err != nil {
		return newCustomError(constant.PasswordHashError)
	}

	return d.repository.InsertUser(user)
}
