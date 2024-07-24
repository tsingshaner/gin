package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Options copy from github.com/gin-contrib/cors exclude func Field
type CorsOptions struct {
	Enabled         bool `json:"enabled" yaml:"enabled" toml:"enabled"`
	AllowAllOrigins bool `json:"allowAllOrigins" yaml:"allowAllOrigins" toml:"allowAllOrigins"`

	// AllowOrigins is a list of origins a cross-domain request can be executed from.
	// If the special "*" value is present in the list, all origins will be allowed.
	// Default value is []
	AllowOrigins []string `json:"allowOrigins" yaml:"allowOrigins" toml:"allowOrigins"`

	// AllowMethods is a list of methods the client is allowed to use with
	// cross-domain requests. Default value is simple methods (GET, POST, PUT, PATCH, DELETE, HEAD, and OPTIONS)
	AllowMethods []string `json:"allowMethods" yaml:"allowMethods" toml:"allowMethods"`

	// AllowPrivateNetwork indicates whether the response should include allow private network header
	AllowPrivateNetwork bool `json:"allowPrivateNetwork" yaml:"allowPrivateNetwork" toml:"allowPrivateNetwork"`

	// AllowHeaders is list of non simple headers the client is allowed to use with
	// cross-domain requests.
	AllowHeaders []string `json:"allowHeaders" yaml:"allowHeaders" toml:"allowHeaders"`

	// AllowCredentials indicates whether the request can include user credentials like
	// cookies, HTTP authentication or client side SSL certificates.
	AllowCredentials bool `json:"allowCredentials" yaml:"allowCredentials" toml:"allowCredentials"`

	// ExposeHeaders indicates which headers are safe to expose to the API of a CORS
	// API specification
	ExposeHeaders []string `json:"exposeHeaders" yaml:"exposeHeaders" toml:"exposeHeaders"`

	// MaxAge indicates how long (with second-precision) the results of a preflight request
	// can be cached
	MaxAge time.Duration `json:"maxAge" yaml:"maxAge" toml:"maxAge"`

	// Allows to add origins like http://some-domain/*, https://api.* or http://some.*.subdomain.com
	AllowWildcard bool `json:"allowWildcard" yaml:"allowWildcard" toml:"allowWildcard"`

	// Allows usage of popular browser extensions schemas
	AllowBrowserExtensions bool `json:"allowBrowserExtensions" yaml:"allowBrowserExtensions" toml:"allowBrowserExtensions"`

	// Allows to add custom schema like tauri://
	CustomSchemas []string `json:"customSchemas" yaml:"customSchemas" toml:"customSchemas"`

	// Allows usage of WebSocket protocol
	AllowWebSockets bool `json:"allowWebSockets" yaml:"allowWebSockets" toml:"allowWebSockets"`

	// Allows usage of file:// schema (dangerous!) use it only when you 100% sure it's needed
	AllowFiles bool `json:"allowFiles" yaml:"allowFiles" toml:"allowFiles"`

	// Allows to pass custom OPTIONS response status code for old browsers / clients
	OptionsResponseStatusCode int `json:"optionsResponseStatusCode" yaml:"optionsResponseStatusCode" toml:"optionsResponseStatusCode"`
}

func Cors(opts *CorsOptions) gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:           opts.AllowAllOrigins,
		AllowOrigins:              opts.AllowOrigins,
		AllowMethods:              opts.AllowMethods,
		AllowPrivateNetwork:       opts.AllowPrivateNetwork,
		AllowHeaders:              opts.AllowHeaders,
		AllowCredentials:          opts.AllowCredentials,
		ExposeHeaders:             opts.ExposeHeaders,
		MaxAge:                    opts.MaxAge,
		AllowWildcard:             opts.AllowWildcard,
		AllowBrowserExtensions:    opts.AllowBrowserExtensions,
		CustomSchemas:             opts.CustomSchemas,
		AllowWebSockets:           opts.AllowWebSockets,
		AllowFiles:                opts.AllowFiles,
		OptionsResponseStatusCode: opts.OptionsResponseStatusCode,
	})
}
