package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lab-online/internal/user/entity"
	errors "github.com/lab-online/internal/user/error"
	"github.com/lab-online/pkg/middleware"
)

type RegisterReqBody struct {
	UserID   string `json:"userID" form:"userID" binding:"required"`
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterReqQuery struct {
	Username string `json:"username" form:"username" binding:"required"`
}

var RegisterReq = &middleware.ValidatorOptions{
	Body:  &RegisterReqBody{},
	Query: &RegisterReqQuery{},
}

func (h *Handler) Register(c *gin.Context) {
	user := c.MustGet(middleware.KeyBody).(*RegisterReqBody)
	username := c.MustGet(middleware.KeyQuery).(*RegisterReqQuery)

	userEntity := userSchemaToEntity(user)
	if err := h.domain.AddUser(userEntity); err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":   "success",
		"data":  user,
		"query": username,
	})
}

func userSchemaToEntity(user *RegisterReqBody) *entity.User {
	return &entity.User{
		UserID:   user.UserID,
		Username: user.Username,
		Password: user.Password,
	}
}
