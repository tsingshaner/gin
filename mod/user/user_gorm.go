//go:generate go run github.com/tsingshaner/gin/cmd/gen --config=../../config/config.example.yml
package user

import (
	"path/filepath"
	"runtime"

	"github.com/tsingshaner/gin/mod/user/dto"
	"github.com/tsingshaner/gin/mod/user/model"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE username = @username{{if role != nil}} AND role = @role{{end}}
	FilterWithNameAndRole(username string, role *dto.Role) ([]gen.T, error)
}

func GenQuery(db *gorm.DB) {
	_, file, _, _ := runtime.Caller(0)

	g := gen.NewGenerator(gen.Config{
		OutPath: filepath.Dir(file) + "/internal/repository/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(db)

	g.ApplyBasic(model.User{})

	g.ApplyInterface(func(Querier) {}, model.User{})

	g.Execute()
}
