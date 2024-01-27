package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lab-online/pkg/jwt"
	"gorm.io/gorm"
)

type Server interface {
	Migrate() error
	RoutesRegister(r *gin.RouterGroup)
}

type AppContext struct {
	DB  *gorm.DB
	jwt jwt.JWTAction
}

func NewApp(db *gorm.DB, jwt jwt.JWTAction) Server {
	return &AppContext{db, jwt}
}
