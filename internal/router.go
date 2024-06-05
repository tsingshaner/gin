package app

import (
	"github.com/gin-gonic/gin"
	"github.com/tsingshaner/gin-starter/internal/user"
	"github.com/tsingshaner/gin-starter/pkg/router"
)

func (app *Context) RoutesRegister(r *gin.RouterGroup) {
	v1 := r.Group("/v1")

	userRoutes := user.NewUserRoutes(app.DB, app.jwt, app.auth)

	router.Register(v1, &[]router.Router{
		userRoutes.Register,
	})
}
