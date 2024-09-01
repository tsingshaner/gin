package constant

import (
	_ "embed"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/tsingshaner/go-pkg/conf"
	"github.com/tsingshaner/go-pkg/errors/gen"
	"github.com/tsingshaner/go-pkg/log/console"
)

//go:embed success.go.template
var successPkgTemplateStr string
var SuccessPkgTemplate *template.Template

func init() {
	var err error
	if SuccessPkgTemplate, err = template.New("SuccessPkg").Parse(successPkgTemplateStr); err != nil {
		console.Fatal("%+v", err)
	}
}

func GenerateSuccessAndErrorConstants() {
	c := conf.Read[config]()
	mkdirForConstant(c.SuccessFile)
	mkdirForConstant(c.File)
	generateSuccessCode(c)
	gen.GeneratePkg(&c.ErrorConfig)
}

type (
	successCode struct {
		Key string `mapstructure:"key"`
		Msg string `mapstructure:"msg"`
	}

	config struct {
		gen.ErrorConfig `mapstructure:",squash"`
		SuccessPackage  string                 `mapstructure:"successPkg"`
		SuccessFile     string                 `mapstructure:"successFile"`
		Successes       map[string]successCode `mapstructure:"success"`
	}
)

type (
	successInfo struct {
		Code string
		Key  string
		Msg  string
	}

	successPkgData struct {
		Package   string
		Successes []successInfo
	}
)

func parseSuccessPkgData(c *config) *successPkgData {
	data := &successPkgData{
		Package: c.SuccessPackage,
	}

	for key, value := range c.Successes {
		data.Successes = append(data.Successes, successInfo{
			Code: strings.ToUpper(c.ModCode + "OK" + key),
			Key:  value.Key,
			Msg:  value.Msg,
		})
	}

	return data
}

func generateSuccessCode(c *config) {
	file, err := os.OpenFile(c.SuccessFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		console.Fatal("open file (%s) err %+v", c.SuccessFile, err)
	}

	if err := SuccessPkgTemplate.Execute(file, parseSuccessPkgData(c)); err != nil {
		console.Fatal("%+v", err)
	}

	console.Info("generate pkg %s success", c.SuccessFile)
}

func mkdirForConstant(path string) {
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		console.Fatal("mkdir %s err %+v", filepath.Dir(path), err)
	}
}
