package app

import (
	user "github.com/lab-online/internal/user/constant"
	"github.com/lab-online/pkg/gen"
)

func GetCodeMaps() []gen.CodeMap {
	return []gen.CodeMap{
		user.GetCodeMap(),
	}
}
