package constant

import (
	"net/http"

	"github.com/lab-online/pkg/gen"
	"github.com/lab-online/pkg/resp"
)

const domainCode = 100

const (
	SUCCESS_REGISTER = resp.CodeSuccess + domainCode + iota
	SUCCESS_LOGIN
	GET_PROFILE_SUCCESS
)

const (
	USER_ALREADY_EXISTS = resp.CodeBadRequest + domainCode + iota
	USER_NOT_EXISTS
	USER_PASSWORD_ERROR
)

const (
	DB_ERROR = resp.CodeServerError + domainCode + iota
	PASSWORD_HASH_ERROR
	TOKEN_GEN_ERROR
)

var ErrorMessage = map[int]string{
	USER_ALREADY_EXISTS: "user already exists",
	USER_NOT_EXISTS:     "user not exists",
	USER_PASSWORD_ERROR: "user password error",

	DB_ERROR:            "server database error",
	PASSWORD_HASH_ERROR: "password hash error",
	TOKEN_GEN_ERROR:     "token gen error",
}

var StatusMap = map[int]int{
	USER_ALREADY_EXISTS: http.StatusBadRequest,
	USER_NOT_EXISTS:     http.StatusBadRequest,
	USER_PASSWORD_ERROR: http.StatusBadRequest,

	DB_ERROR:            http.StatusInternalServerError,
	PASSWORD_HASH_ERROR: http.StatusInternalServerError,
	TOKEN_GEN_ERROR:     http.StatusInternalServerError,
}

func GetCodeMap() gen.CodeMap {
	return gen.CodeMap{
		"Register": {
			SUCCESS_REGISTER: "注册成功",

			USER_ALREADY_EXISTS: "用户已存在",

			DB_ERROR:            "服务端数据库链接错误",
			PASSWORD_HASH_ERROR: "服务端密码加密错误",
			TOKEN_GEN_ERROR:     "服务端token生成错误",
		},
		"Login": {
			SUCCESS_LOGIN: "登录成功",

			USER_NOT_EXISTS:     "用户不存在",
			USER_PASSWORD_ERROR: "用户密码错误",

			DB_ERROR:            "服务端数据库链接错误",
			PASSWORD_HASH_ERROR: "服务端密码加密错误",
			TOKEN_GEN_ERROR:     "服务端token生成错误",
		},
		"GetProfile": {
			GET_PROFILE_SUCCESS: "获取用户信息成功",
		},
	}
}
