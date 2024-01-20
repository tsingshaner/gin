package domain

import (
	"github.com/lab-online/internal/user/entity"
	errors "github.com/lab-online/internal/user/error"
)

func (d *Domain) AddUser(user *entity.User) error {
	if exist, err := d.repository.CheckUserExists(user.UserID); err != nil {
		return err
	} else if exist {
		return errors.New(errors.USER_ALREADY_EXISTS)
	}

	if hashedPassword, err := d.HashPassword(user.Password); err != nil {
		return errors.New(errors.USER_PASSWORD_EMPTY)
	} else {
		user.Password = hashedPassword
	}

	return d.repository.InsertUser(user)
}
