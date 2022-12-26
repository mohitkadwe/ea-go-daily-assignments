package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID    int64   `json:"id,string,omitempty"`
	Title string  `json:"title"`
	Price float64 `json:"price,string,omitempty"`
}

var books = []book{
	{ID: 1, Title: "In Search of Lost Time", Price: 200},
	{ID: 2, Title: "The Great Gatsby", Price: 300.1},
	{ID: 3, Title: "War and Peace", Price: 233.2},
}

func convertToFloat(price string) float64 {
	pr, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return 0
	}
	return pr
}

func getBooks(c *gin.Context) {
	result, err := json.Marshal(books)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Request", "error": err.Error()})
	}
	c.String(http.StatusOK, string(result))
}

func createBook(c *gin.Context) {
	var newBook book

	id := len(books) + 1
	newBook.ID = int64(id)
	title := c.PostForm("title")
	newBook.Title = title
	price := c.PostForm("price")
	newBook.Price = convertToFloat(price)

	books = append(books, newBook)

	c.String(http.StatusCreated, "Book created successfully")
}

func updateBook(c *gin.Context) {

	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	i, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Request", "error": err.Error()})
	}

	var updatedBook book
	title := c.PostForm("title")
	updatedBook.Title = title
	price := c.PostForm("price")
	updatedBook.Price = convertToFloat(price)
	updatedBook.ID = i

	for ele, book := range books {
		if book.ID == int64(i) {
			books[ele] = updatedBook
		}
	}

	c.String(http.StatusOK, "Book updated successfully")
}

func RemoveIndex(s []book, index int64) []book {
	return append(s[:index], s[index+1:]...)
}

func deleteBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	i, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Request", "error": err.Error()})
	}

	books = RemoveIndex(books, i)

	c.String(http.StatusOK, "Book deleted successfully")
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.PUT("/books", updateBook)
	router.DELETE("/books", deleteBook)
	return router
}

func main() {
	router := setupRouter()
	router.Run("localhost:8080")
}
