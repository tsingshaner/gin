package rest

import "github.com/gin-gonic/gin"

type UserHandler interface {
	Register(*gin.Context)

	DeleteUser(*gin.Context)
	GetProfile(*gin.Context)
	GetUserList(*gin.Context)
	Login(*gin.Context)
	SearchUser(*gin.Context)
	UpdateUser(*gin.Context)
}
