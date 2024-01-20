package http

import "github.com/gin-gonic/gin"

type UserHandler interface {
	Register(*gin.Context)

	DeleteUser(*gin.Context)
	GetUserProfile(*gin.Context)
	GetUserList(*gin.Context)
	Login(*gin.Context)
	SearchUser(*gin.Context)
	UpdateUser(*gin.Context)
}
