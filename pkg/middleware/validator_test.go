package middleware

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestValidator(t *testing.T) {
	type Body struct {
		Name string `json:"name" binding:"required"`
	}

	type Params struct {
		ID string `uri:"id" binding:"required"`
	}

	type Query struct {
		Key string `form:"key" binding:"required"`
	}

	type Header struct {
		ContentType string `header:"Content-Type" binding:"required"`
	}

	opts := &ValidatorOptions{
		BodyKey:   "body",
		ParamsKey: "params",
		QueryKey:  "query",
		HeaderKey: "header",
		Body:      &Body{},
		Params:    &Params{},
		Query:     &Query{},
		Header:    &Header{},
	}

	handler := Validator(opts)

	assert.NotNil(t, handler)

	c := &gin.Context{
		Params: gin.Params{{Key: "id", Value: "123"}},
		Request: &http.Request{
			Body: io.NopCloser(strings.NewReader("{\"name\":\"test\"}")),
			URL: &url.URL{
				Path:     "/path",
				RawQuery: "key=value",
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
	}

	handler(c)

	c.ShouldBindJSON(opts.Body)
	c.ShouldBindUri(opts.Params)
	c.ShouldBindQuery(opts.Query)
	c.ShouldBindHeader(opts.Header)

	assert.Equal(t, c.MustGet(opts.BodyKey).(*Body), opts.Body)
	assert.Equal(t, c.MustGet(opts.ParamsKey).(*Params), opts.Params)
	assert.Equal(t, c.MustGet(opts.QueryKey).(*Query), opts.Query)
	assert.Equal(t, c.MustGet(opts.HeaderKey).(*Header), opts.Header)
}
