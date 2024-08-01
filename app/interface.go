package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Provider interface {
	Engine() *gin.Engine
	Database() *gorm.DB
}

type Server interface {
	Provider

	Ready() error
	Start()
	Shutdown() error
}
