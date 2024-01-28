package export

import (
	"github.com/lab-online/internal/user/entity"
	"github.com/lab-online/internal/user/interface/export"
)

type User struct {
	repository Repository
}

func New(repository Repository) export.UserExport {
	return &User{repository}
}

func (u *User) GetUser(id uint) (entity.UserEntity, error) {
	return u.repository.FindUserByID(id)
}

func (u *User) GetManyUser(id []uint) ([]entity.UserEntity, error) {
	return u.repository.FindManyUserByID(id)
}
