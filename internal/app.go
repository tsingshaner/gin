package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server interface {
	Migrate() error
	RoutesRegister(r *gin.RouterGroup)
}

type AppContext struct {
	db *gorm.DB
}

func NewApp(db *gorm.DB) *AppContext {
	return &AppContext{db}
}
