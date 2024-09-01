package app

import userModel "github.com/tsingshaner/gin/mod/user/model"

func (a *app) Migrate() {
	a.Database().AutoMigrate(&userModel.User{})
}
