package requestid

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

type Options struct {
	// HeaderKey is the key of request id in header
	HeaderKey string `mapstructure:"headerKey"`
}

func New(opts *Options) gin.HandlerFunc {
	return requestid.New(
		requestid.WithCustomHeaderStrKey(requestid.HeaderStrKey(opts.HeaderKey)),
		requestid.WithGenerator(RequestIdGenerator),
	)
}

// Get get request id from gin context
var Get = requestid.Get
