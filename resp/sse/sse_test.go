package sse

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type mockedResponseWriter struct {
	header http.Header
	body   []byte
}

var _ http.ResponseWriter = &mockedResponseWriter{}

func (m *mockedResponseWriter) Header() http.Header {
	return m.header
}

func (m *mockedResponseWriter) Write([]byte) (int, error) {
	m.body = append(m.body, []byte("data: test\n\n")...)
	return 0, nil
}

func (m *mockedResponseWriter) WriteHeader(int) {}

func TestSSEWithUnSupportedRespWriter(t *testing.T) {
	mockedResponseWriter := &mockedResponseWriter{
		header: make(http.Header),
	}

	err := Upgrade(mockedResponseWriter, time.Now().Add(5))

	assert.ErrorIs(t, err, RESTErrSSEUnsupported)
}

type mockedResponseWriterWithSetWriteDeadline struct {
	mockedResponseWriter
	Deadline time.Time
}

func (m *mockedResponseWriterWithSetWriteDeadline) SetWriteDeadline(deadline time.Time) error {
	m.Deadline = deadline
	return nil
}

func TestSSEWithSupportedRespWriter(t *testing.T) {
	mockedResponseWriter := &mockedResponseWriterWithSetWriteDeadline{
		mockedResponseWriter: mockedResponseWriter{
			header: make(http.Header),
		},
	}

	deadline := time.Now().Add(5)
	err := Upgrade(mockedResponseWriter, deadline)

	assert.NoError(t, err)

	assert.Equal(t, mockedResponseWriter.Deadline, deadline)
	assert.Equal(t, mockedResponseWriter.Header().Get("Content-Type"), "text/event-stream, charset=utf-8")
	assert.Equal(t, mockedResponseWriter.Header().Get("Cache-Control"), "no-cache")
	assert.Equal(t, mockedResponseWriter.Header().Get("Connection"), "keep-alive")
}
