package app

import (
	"github.com/gin-gonic/gin"
	"github.com/tsingshaner/go-pkg/log"
	"gorm.io/gorm"
)

type Provider interface {
	Database() *gorm.DB
	Engine() *gin.Engine
	Providers() *providers
	Logger() log.Slog
}

type Server interface {
	Provider
	// Migrate gorm auto migrate
	Migrate()

	// Ready init server
	Ready()
	// InitProviders init dependencies mods, like user, comment mod
	InitProviders()
	// BindRoutes binding routes
	BindRoutes()
	// Start start serve and registry shutdown hook
	Start()
	// Shutdown listening kill signal
	Shutdown()
}
