package entity

import (
	"github.com/tsingshaner/gin-starter/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	logger.Debug("HashPassword", "hashedPassword", u.Password)
	return nil
}
