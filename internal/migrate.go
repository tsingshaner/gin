package app

import "github.com/lab-online/internal/user/infra/model"

func (app *AppContext) Migrate() error {
	return app.DB.AutoMigrate(
		&model.User{},
	)
}
