package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/tsingshaner/gin-starter/pkg/auth"
	"github.com/tsingshaner/gin-starter/pkg/jwt"

	"github.com/tsingshaner/gin-starter/internal/user/domain"
	"github.com/tsingshaner/gin-starter/internal/user/handler"
	"github.com/tsingshaner/gin-starter/internal/user/infra/repository"
	"github.com/tsingshaner/gin-starter/internal/user/interface/rest"
)

type Route struct {
	rest.UserHandler
	auth auth.Middleware
}

func NewUserRoutes(db *gorm.DB, jwt jwt.JWTAction, auth auth.Middleware) *Route {
	r := repository.NewRepository(db)
	d := domain.NewDomain(r, jwt)
	h := handler.NewHandler(d)

	return &Route{h, auth}
}

func (r *Route) Register(router *gin.RouterGroup) {
	user := router.Group("/user/")

	user.DELETE(":id", r.DeleteUser)
	user.GET("", r.GetUserList)
	user.GET(
		"profile", r.auth([]auth.Role{auth.RoleNone}),
		r.GetProfile,
	)
	user.PATCH(
		":id", r.auth([]auth.Role{auth.RoleNone}),
		r.UpdateUser,
	)
	user.POST(
		"", r.auth([]auth.Role{auth.RoleTeacher}),
		handler.RegisterValidator, r.UserHandler.Register,
	)
	user.POST("login", handler.LoginValidator, r.Login)
	user.PUT(":id", r.UpdateUser)
}
