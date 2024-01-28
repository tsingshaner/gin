package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lab-online/pkg/auth"
	"github.com/lab-online/pkg/jwt"
	"gorm.io/gorm"
)

type Server interface {
	Migrate() error
	RoutesRegister(r *gin.RouterGroup)
}

type Context struct {
	DB   *gorm.DB
	jwt  jwt.JWTAction
	auth auth.Middleware
}

func New(db *gorm.DB, jwt jwt.JWTAction) Server {
	auth := auth.New(jwt)

	return &Context{db, jwt, auth}
}
