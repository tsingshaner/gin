package main

import (
	"github.com/tsingshaner/gin-starter/config"
	app "github.com/tsingshaner/gin-starter/internal"
	"github.com/tsingshaner/gin-starter/pkg/gen"
)

func main() {
	maps := app.GetCodeMaps()
	gen.GenerateCodeDocs(config.Server.APIMarkdown, maps)
}
