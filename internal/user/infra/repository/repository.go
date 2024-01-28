package repository

import (
	"github.com/lab-online/internal/user/domain"
	"github.com/lab-online/internal/user/entity"
	"github.com/lab-online/internal/user/infra/repository/user"
	"gorm.io/gorm"
)

type Repository struct {
	User
}
type User interface {
	CheckUserExists(string) (bool, error)
	InsertUser(entity.UserEntity) error
	FindByUserID(string) (entity.UserEntity, error)
}

func NewRepository(db *gorm.DB) domain.UserRepository {
	u := user.New(db)

	return &Repository{u}
}
