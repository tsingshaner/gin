package repository

import (
	"github.com/tsingshaner/gin-starter/internal/user/domain"
	"github.com/tsingshaner/gin-starter/internal/user/entity"
	"github.com/tsingshaner/gin-starter/internal/user/infra/repository/user"
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
