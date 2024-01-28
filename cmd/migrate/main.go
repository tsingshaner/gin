package main

import (
	"github.com/lab-online/config"
	app "github.com/lab-online/internal"
	"github.com/lab-online/pkg/database"
	"github.com/lab-online/pkg/logger"
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
