package database

import (
	"time"

	"github.com/lab-online/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(dns string, config *gorm.Config) *gorm.DB {
	var db *gorm.DB
	var err error

	for i := 0; i < 5; i++ {
		db, err = gorm.Open(postgres.Open(dns), config)
		if err == nil {
			logger.Info("connected database", dns)
			return db
		}

		logger.Error("failed to connect database, will retry after", i*2, "s")
		time.Sleep(time.Second * time.Duration(i*2))
	}

	logger.Error("failed to connect database", dns)
	panic(err)
}
