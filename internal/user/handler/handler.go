package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/lab-online/pkg/errors"

	"github.com/lab-online/internal/user/constant"
	"github.com/lab-online/internal/user/domain"
	"github.com/lab-online/internal/user/interface/rest"
)

type Handler struct {
	domain domain.UserDomain
}

func NewHandler(domain domain.UserDomain) rest.UserHandler {
	return &Handler{
		domain: domain,
	}
}

func handleError(c *gin.Context, err error) {
	errors.HandleError(c, err, constant.StatusMap)
}
