package main

import (
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"

	userModels "github.com/tsingshaner/gin/mod/user/model"
	"github.com/tsingshaner/go-pkg/log/console"
)

func main() {
	stmts, err := gormschema.New("postgres").Load(&userModels.User{})
	if err != nil {
		console.Fatal("failed to load gorm schema: %v\n", err)
	}
	io.WriteString(os.Stdout, stmts)
}
