package app

import (
	"fmt"
	"log/slog"

	"github.com/tsingshaner/go-pkg/color"
	"github.com/tsingshaner/go-pkg/log"
	"github.com/tsingshaner/go-pkg/log/console"
	"github.com/tsingshaner/go-pkg/log/helper"
	expRetry "github.com/tsingshaner/go-pkg/util/expretry"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres(logger log.Slog, opts *Postgres, loggerOpts GormLoggerOptions) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		opts.Host,
		opts.User,
		opts.Password,
		opts.Database,
		opts.Port,
		opts.SSLMode,
		opts.TimeZone,
	)

	db, err := expRetry.New(func() (*gorm.DB, error) {
		return gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: helper.NewGormLogger(logger, helper.GORMLoggerOptions{
				SlowThreshold:             loggerOpts.SlowThreshold,
				LogLevel:                  slog.Level(loggerOpts.LogLevel),
				IgnoreRecordNotFoundError: loggerOpts.IgnoreRecordNotFoundError,
				ParameterizedQueries:      loggerOpts.ParameterizedQueries,
			}),
		})
	}, func(ers *expRetry.ExpRetrySvc[*gorm.DB]) {
		ers.ErrorHandler = func(e error, times int) {}
	}).Run()
	if err != nil {
		console.Fatal("failed to connect database %s", err)
	}

	console.Info("database connected %s",
		color.Underline(color.UnsafeCyan(fmt.Sprintf("postgres://%s:***@%s:%d/%s\n",
			opts.User,
			opts.Host,
			opts.Port,
			opts.Database,
		))),
	)

	return db
}

func (a *app) connectDatabase() {
	a.db = ConnectPostgres(a.logger, a.Options.Postgres, a.Options.Logger.GormLogger)
}
