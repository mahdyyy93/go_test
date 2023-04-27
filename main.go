package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: 1, Title: "In search for time", Author: "Max", Quantity: 4},
	{ID: 2, Title: "In search for love", Author: "Manny", Quantity: 2},
	{ID: 4, Title: "In search for passion", Author: "Fog", Quantity: 1},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func getBookById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	book, err := findBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func findBookById(id int) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found.")
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookById)
	router.POST("/books", createBook)
	router.Run("localhost:8080")
}
