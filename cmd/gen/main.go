package main

import (
	"github.com/tsingshaner/gin/app"
	"github.com/tsingshaner/gin/config"
	"github.com/tsingshaner/gin/mod/user"
)

func main() {
	config.Init()
	conf := config.Store().Options
	server := app.New(&conf)
	db := app.ConnectPostgres(server.Logger(), config.Store().Postgres, conf.Logger.GormLogger)

	user.GenQuery(db)
}
