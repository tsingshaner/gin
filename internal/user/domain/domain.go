package domain

import (
	"github.com/lab-online/internal/user/entity"
)

type UserDomain interface {
	EqualPassword(password string, hashedPassword string) bool
	HashPassword(password string) (string, error)
	AddUser(*entity.User) error
	Login(*entity.User) (string, error)
}

type Domain struct {
	repository UserRepository
}

func NewDomain(repository UserRepository) UserDomain {
	return &Domain{
		repository: repository,
	}
}
