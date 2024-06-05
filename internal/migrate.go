package app

import "github.com/tsingshaner/gin-starter/internal/user/infra/model"

func (app *Context) Migrate() error {
	return app.DB.AutoMigrate(
		&model.User{},
	)
}
