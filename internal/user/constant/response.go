package constant

import (
	"net/http"

	"github.com/lab-online/pkg/gen"
	"github.com/lab-online/pkg/resp"
)

const domainCode = 100

const (
	SUCCESS_REGISTER = resp.CodeSuccess + domainCode + iota
)

const (
	USER_ALREADY_EXISTS = resp.CodeBadRequest + domainCode + iota
	USER_PASSWORD_EMPTY
)

const (
	DB_ERROR = resp.CodeServerError + domainCode + iota
)

var ErrorMessage = map[int]string{
	USER_ALREADY_EXISTS: "user already exists",
	USER_PASSWORD_EMPTY: "user password empty",
}

var StatusMap = map[int]int{
	USER_ALREADY_EXISTS: http.StatusBadRequest,
	USER_PASSWORD_EMPTY: http.StatusBadRequest,
}

func GetCodeMap() gen.CodeMap {
	return gen.CodeMap{
		"Register": {
			SUCCESS_REGISTER: "注册成功",

			USER_ALREADY_EXISTS: "用户已存在",
			USER_PASSWORD_EMPTY: "用户密码为空",

			DB_ERROR: "服务端数据库链接错误",
		},
	}
}
