package main

import (
	"github.com/tsingshaner/gin-starter/config"
	app "github.com/tsingshaner/gin-starter/internal"
	"github.com/tsingshaner/gin-starter/pkg/database"
	"github.com/tsingshaner/gin-starter/pkg/logger"
	"gorm.io/gorm"
)

func main() {
	db := database.ConnectDB(config.Database.Postgres, &gorm.Config{})

	server := &app.Context{DB: db}

	if err := server.Migrate(); err != nil {
		logger.Error("migrate database failed")
		panic(err)
	}

	logger.Info("migrate database success")
}
