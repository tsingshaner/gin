package repository

import (
	"github.com/lab-online/internal/user/domain"
	"gorm.io/gorm"
)

type Repository struct{
	db *gorm.DB
}

func NewRepository(db *gorm.DB) domain.UserRepository {
	return &Repository{db}
}
