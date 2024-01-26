package repository

import (
	"github.com/lab-online/internal/user/domain"
	"github.com/lab-online/pkg/model"
	"gorm.io/gorm"
)

type Repository struct {
	db   *gorm.DB
	user *gorm.DB
}

func NewRepository(db *gorm.DB) domain.UserRepository {
	return &Repository{
		db:   db,
		user: db.Model(&model.User{}),
	}
}
