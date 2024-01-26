package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/lab-online/pkg/jwt"
	"github.com/lab-online/pkg/middleware"

	"github.com/lab-online/internal/user/domain"
	"github.com/lab-online/internal/user/handler"
	"github.com/lab-online/internal/user/infra/repository"
	"github.com/lab-online/internal/user/interface/http"
)

type UserRoutes struct {
	http.UserHandler
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
	user.POST("", middleware.Validator(handler.RegisterReq), u.UserHandler.Register)
	user.POST("login", u.Login)
	user.PUT(":id", u.UpdateUser)
}
