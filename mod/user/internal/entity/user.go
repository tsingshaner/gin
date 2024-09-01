package entity

import (
	"errors"

	"github.com/tsingshaner/gin/mod/user/constant/errs"
	"github.com/tsingshaner/gin/mod/user/dto"
	"github.com/tsingshaner/go-pkg/jwt"
	"github.com/tsingshaner/go-pkg/util"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	passwordHashed bool
	*dto.User
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.Join(errs.Basic.Hash, err)
	}

	return string(hashedPassword), nil
}

func (u *User) HashPassword() error {
	if hashedPassword, err := HashPassword(u.Password); err != nil {
		return err
	} else {
		u.Password = hashedPassword
	}

	u.passwordHashed = true
	return nil
}

func (u *User) VerifyPassword(password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	if err == nil {
		return true, nil
	}

	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return false, nil
	}

	return false, err
}

func (u *User) PasswordHasHashed() bool {
	return u.passwordHashed
}

func (u *User) BuildAuthPayload(base *jwt.RegisteredClaims) *dto.AuthPayload {
	return &dto.AuthPayload{
		RegisteredClaims: base,
		Role:             u.Role,
		UID:              u.ID,
	}
}

func New(fns ...util.WithFn[User]) *User {
	return util.BuildWithOpts(&User{}, fns...)
}

func WithUser(u *dto.User) util.WithFn[User] {
	return func(e *User) {
		e.User = u
	}
}

func WithPasswordHashed() util.WithFn[User] {
	return func(u *User) {
		u.passwordHashed = true
	}
}
