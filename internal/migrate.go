package app

import (
	userModel "github.com/lab-online/internal/user/infra/model"
)

func (app *AppContext) Migrate() error {
	return app.DB.AutoMigrate(&userModel.User{})
}
