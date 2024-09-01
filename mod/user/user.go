package user

import (
	"github.com/tsingshaner/gin/mod/user/internal/handler"
	"github.com/tsingshaner/gin/mod/user/internal/repository"
	"github.com/tsingshaner/gin/mod/user/internal/service"
)

var (
	NewHandler            = handler.New
	NewRepository         = repository.NewUser
	NewAuthProvider       = service.NewAuth
	NewUserQueryProvider  = service.NewUserQuery
	NewUserModifyProvider = service.NewUserModify
	NewVerifyProvider     = service.NewVerify
)

type (
	HandlerProvider = handler.Provider
	UserRepo        = repository.User
)
