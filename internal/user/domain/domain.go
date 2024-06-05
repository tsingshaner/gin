package domain

import (
	"github.com/tsingshaner/gin-starter/internal/user/constant"
	"github.com/tsingshaner/gin-starter/internal/user/entity"
	"github.com/tsingshaner/gin-starter/pkg/errors"
	"github.com/tsingshaner/gin-starter/pkg/jwt"
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
