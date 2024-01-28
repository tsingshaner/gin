package domain

import (
	"github.com/lab-online/internal/user/constant"
	"github.com/lab-online/internal/user/entity"
	"github.com/lab-online/pkg/errors"
	"github.com/lab-online/pkg/jwt"
)

type UserDomain interface {
	AddUser(entity.UserEntity) error
	GetUser(entity.UserEntity) (entity.UserEntity, error)
	Login(entity.UserEntity) (string, error)
}

type Domain struct {
	repository UserRepository
	jwt        jwt.JWTAction
}

func NewDomain(repository UserRepository, jwt jwt.JWTAction) UserDomain {
	return &Domain{repository, jwt}
}

func newCustomError(code int) *errors.Error {
	return errors.New(code, constant.ErrorMessage[code])
}
