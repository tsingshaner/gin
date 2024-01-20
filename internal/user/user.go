package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lab-online/internal/user/domain"
	"github.com/lab-online/internal/user/handler"
	"github.com/lab-online/internal/user/infra/repository"
	"github.com/lab-online/internal/user/interface/http"
	"github.com/lab-online/pkg/middleware"
	"gorm.io/gorm"
)

type UserRoutes struct {
	http.UserHandler
}

func NewUserRoutes(db *gorm.DB) *UserRoutes {
	r := repository.NewRepository(db)
	d := domain.NewDomain(r)
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
