package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddJob(t *testing.T) {

	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/add-job", nil)
	req.PostForm = map[string][]string{
		"Id":   {"123"},
		"Name": {"build-server"},
	}

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Equal(t, "{\"message\":\"OK\",\"results_length\":4}", w.Body.String())
}
