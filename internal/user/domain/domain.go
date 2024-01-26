package domain

import (
	"github.com/lab-online/internal/user/constant"
	"github.com/lab-online/internal/user/entity"
	errors "github.com/lab-online/pkg/error"
	"github.com/lab-online/pkg/jwt"
)

type UserDomain interface {
	AddUser(entity.UserEntity) error
	Login(*entity.User) (string, error)
}

type Domain struct {
	repository UserRepository
	jwt        jwt.JWTAction
}

func NewDomain(repository UserRepository, jwt jwt.JWTAction) UserDomain {
	return &Domain{repository, jwt}
}

func newCustomError(code int) *errors.Error {
	return errors.New(code, constant.ErrorMessage)
}
