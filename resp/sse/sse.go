package sse

import (
	"errors"
	"net/http"
	"time"

	errs "github.com/tsingshaner/go-pkg/errors"
)

const (
	TypeData  = "data"
	TypePing  = "ping"
	TypeOpen  = "open"
	TypeClose = "close"
	TypeError = "error"
)

var (
	// RESTErrSSEUnsupported is a REST error for unsupported to update connection deadline.
	RESTErrSSEUnsupported = errs.NewREST(http.StatusHTTPVersionNotSupported, "SSE505", "Unsupported to update connection deadline")
	// RESTErrSSEUnkown is a REST error for unknown error.
	RESTErrSSEUnkown = errs.NewREST(http.StatusInternalServerError, "SSE500", "Unknown error")
)

// Upgrade updates the connection deadline, and sets the response headers for SSE.
func Upgrade(writer http.ResponseWriter, deadline time.Time) error {
	responseController := http.NewResponseController(writer)
	if err := responseController.SetWriteDeadline(deadline); err != nil {
		if errors.Is(err, http.ErrNotSupported) {
			return errors.Join(RESTErrSSEUnsupported, err)
		}

		return errors.Join(RESTErrSSEUnkown, err)
	}

	writer.Header().Set("Content-Type", "text/event-stream, charset=utf-8")
	writer.Header().Set("Cache-Control", "no-cache")
	writer.Header().Set("Connection", "keep-alive")

	return nil
}
