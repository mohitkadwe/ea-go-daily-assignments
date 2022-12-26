package main

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorldHandler(t *testing.T) {

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/hello", nil)

	helloWorldHandler(rec, req)

	body := rec.Result().Body
	data, _ := io.ReadAll(body)
	assert.Equal(t, "Hello World", string(data))
}
