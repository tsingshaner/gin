package app

import (
	"net/http"

	"github.com/tsingshaner/go-pkg/log/console"
)

// Start init and listening server
func (a *app) Start() {
	a.Ready()

	go func(a *app) {
		console.Info("server will listening on \x1b[36;4mhttp://%s:%d%s\x1b[0m",
			a.Options.Server.Host,
			a.Options.Server.Port,
			a.Options.Server.Base,
		)
		if a.Options.Swagger.Enable {
			console.Info("swagger docs \x1b[36;4mhttp://%s:%d%s%s/index.html\x1b[0m\n",
				a.Options.Server.Host,
				a.Options.Server.Port,
				a.Options.Server.Base,
				a.Options.Swagger.DocsBase,
			)
		}

		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			console.Fatal("server listening failed \n\t%v", err)
		}
	}(a)

	a.Shutdown()
}
