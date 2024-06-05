package export

import (
	"github.com/tsingshaner/gin-starter/internal/user/entity"
	"github.com/tsingshaner/gin-starter/internal/user/interface/export"
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
