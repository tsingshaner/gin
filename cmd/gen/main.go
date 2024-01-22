package main

import (
	"github.com/lab-online/config"
	app "github.com/lab-online/internal"
	"github.com/lab-online/pkg/gen"
)

func main() {
	maps := app.GetCodeMaps()
	gen.GenerateCodeDocs(config.Server.APIMarkdown, maps)
}
