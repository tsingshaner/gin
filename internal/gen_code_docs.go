package app

import (
	user "github.com/tsingshaner/gin-starter/internal/user/constant"
	"github.com/tsingshaner/gin-starter/pkg/gen"
)

func GetCodeMaps() []gen.CodeMap {
	return []gen.CodeMap{
		user.GetCodeMap(),
	}
}
