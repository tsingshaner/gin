package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/lab-online/pkg/jwt"

	"github.com/lab-online/internal/user/domain"
	"github.com/lab-online/internal/user/handler"
	"github.com/lab-online/internal/user/infra/repository"
	"github.com/lab-online/internal/user/interface/rest"
)

type UserRoutes struct {
	rest.UserHandler
}

func NewUserRoutes(db *gorm.DB, jwt jwt.JWTAction) *UserRoutes {
	r := repository.NewRepository(db)
	d := domain.NewDomain(r, jwt)
	h := handler.NewHandler(d)

	return &UserRoutes{h}
}

func (u *UserRoutes) Register(r *gin.RouterGroup) {
	user := r.Group("/user/")

	user.DELETE(":id", u.DeleteUser)
	user.GET("", u.GetUserList)
	user.GET(":id", u.GetUserProfile)
	user.PATCH(":id", u.UpdateUser)
	user.POST("", handler.RegisterValidator, u.UserHandler.Register)
	user.POST("login", handler.LoginValidator, u.Login)
	user.PUT(":id", u.UpdateUser)
}
