package router

import "github.com/gin-gonic/gin"

type Router func(*gin.RouterGroup)

func Register(r *gin.RouterGroup, routes *[]Router) {
	for _, register := range *routes {
		register(r)
	}
}
