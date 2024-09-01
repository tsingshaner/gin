package app

import (
	"context"
	"fmt"

	"github.com/tsingshaner/gin/api/openapi"
	"github.com/tsingshaner/gin/mod/user"
	"github.com/tsingshaner/gin/shared"
	"github.com/tsingshaner/gin/swagger"
)

func (a *app) BindRoutes() {
	r := a.engine.Group(a.Options.Server.Base)

	if a.Options.Swagger.Enable {
		a.Options.Swagger.Host = fmt.Sprintf("%s:%d%s",
			a.Options.Server.Host,
			a.Options.Server.Port,
			a.Options.Server.Base)
		swagger.MergeDocsOptions(openapi.SwaggerInfo, a.Options.Swagger)

		r.GET(
			fmt.Sprintf("%s/*any", a.Options.Swagger.DocsBase),
			swagger.New(a.Options.Swagger.Server),
		)
	}

	for _, route := range a.BuildRestHandlers() {
		route.RegisterRoutes(r)
	}
}

func (a *app) BuildRestHandlers() []shared.Handler {
	if a.providers == nil {
		a.InitProviders()
	}

	p := a.Providers()

	return []shared.Handler{
		user.NewHandler(context.TODO(), &user.HandlerProvider{
			Auth:       p.Auth,
			UserQuery:  p.UserQuery,
			UserModify: p.UserModify,
			Verify:     p.Verify,
		}),
	}
}
