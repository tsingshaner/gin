package main

import (
	"github.com/tsingshaner/gin/app"
	"github.com/tsingshaner/gin/config"
)

func main() {
	app.New(&config.Store().Options).Start()
}
