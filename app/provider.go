package app

import (
	"github.com/tsingshaner/gin/mod/user"
	ui "github.com/tsingshaner/gin/mod/user/interfaces"
)

type providers struct {
	Repo struct {
		User *user.UserRepo
	}

	ui.Auth
	ui.Verify

	ui.UserQuery
	ui.UserModify
}

func (a *app) Providers() *providers {
	return a.providers
}

func (a *app) InitProviders() {
	a.Ready()
	p := &providers{}
	a.providers = p

	p.Repo.User = user.NewRepository(a.Database())

	p.UserQuery = user.NewUserQueryProvider(p.Repo.User)
	p.UserModify = user.NewUserModifyProvider(p.Repo.User, p.UserQuery)
	p.Auth = user.NewAuthProvider(a.jwtMeta, p.UserQuery, p.UserModify)
	p.Verify = user.NewVerifyProvider(a.jwtMeta)
}
