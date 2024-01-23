package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID uint `gorm:"primarykey"`
	DateModel
}

type DateModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
