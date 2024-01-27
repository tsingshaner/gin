package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lab-online/internal/user"
	"github.com/lab-online/pkg/router"
)

func (app *AppContext) RoutesRegister(r *gin.RouterGroup) {
	v1 := r.Group("/v1")

	userRoutes := user.NewUserRoutes(app.DB, app.jwt)

	router.Register(v1, &[]router.Router{
		userRoutes.Register,
	})
}
