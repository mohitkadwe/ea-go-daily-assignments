package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBooks(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	b := []book{
		{ID: 1, Title: "In Search of Lost Time", Price: 200},
		{ID: 2, Title: "The Great Gatsby", Price: 300.1},
		{ID: 3, Title: "War and Peace", Price: 233.2},
	}

	body := w.Body.Bytes()
	var books []book
	err := json.Unmarshal(body, &books)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, b, books)

}

func TestCreateBook(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/books", nil)
	req.PostForm = map[string][]string{
		"title": {"The Little Prince"},
		"price": {"100"},
	}

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "Book created successfully", w.Body.String())
}

func TestUpdateBook(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("PUT", "/books", nil)
	req.URL.RawQuery = "id=1"
	req.PostForm = map[string][]string{
		"title": {"The Little Prince"},
		"price": {"100"},
	}

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Book updated successfully", w.Body.String())
}

func TestDeleteBook(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("DELETE", "/books", nil)
	req.URL.RawQuery = "id=1"

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Book deleted successfully", w.Body.String())
}
