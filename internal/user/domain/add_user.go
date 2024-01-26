package domain

import (
	"github.com/lab-online/internal/user/constant"
	"github.com/lab-online/internal/user/entity"
)

func (d *Domain) AddUser(user entity.UserEntity) error {
	if exist, err := d.repository.CheckUserExists(user.GetUserID()); err != nil {
		return newCustomError(constant.DB_ERROR)
	} else if exist {
		return newCustomError(constant.USER_ALREADY_EXISTS)
	}

	if err := user.HashPassword(); err != nil {
		return newCustomError(constant.PASSWORD_HASH_ERROR)
	}

	return d.repository.InsertUser(user)
}
