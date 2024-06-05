package constant

import (
	"net/http"

	"github.com/tsingshaner/gin-starter/pkg/gen"
	"github.com/tsingshaner/gin-starter/pkg/resp"
)

const domainCode = 100

const (
	SuccessRegister = resp.CodeSuccess + domainCode + iota
	SuccessLogin
	GetProfileSuccess
)

const (
	UserAlreadyExists = resp.CodeBadRequest + domainCode + iota
	UserNotExists
	UserPasswordError
)

const (
	DBError = resp.CodeServerError + domainCode + iota
	PasswordHashError
	TokenGenError
)

var ErrorMessage = map[int]string{
	UserAlreadyExists: "user already exists",
	UserNotExists:     "user not exists",
	UserPasswordError: "user password error",

	DBError:           "server database error",
	PasswordHashError: "password hash error",
	TokenGenError:     "token gen error",
}

var StatusMap = map[int]int{
	UserAlreadyExists: http.StatusBadRequest,
	UserNotExists:     http.StatusBadRequest,
	UserPasswordError: http.StatusBadRequest,

	DBError:           http.StatusInternalServerError,
	PasswordHashError: http.StatusInternalServerError,
	TokenGenError:     http.StatusInternalServerError,
}

func GetCodeMap() gen.CodeMap {
	return gen.CodeMap{
		"Register": {
			SuccessRegister: "注册成功",

			UserAlreadyExists: "用户已存在",

			DBError:           "服务端数据库链接错误",
			PasswordHashError: "服务端密码加密错误",
			TokenGenError:     "服务端token生成错误",
		},
		"Login": {
			SuccessLogin: "登录成功",

			UserNotExists:     "用户不存在",
			UserPasswordError: "用户密码错误",

			DBError:           "服务端数据库链接错误",
			PasswordHashError: "服务端密码加密错误",
			TokenGenError:     "服务端token生成错误",
		},
		"GetProfile": {
			GetProfileSuccess: "获取用户信息成功",
		},
	}
}
