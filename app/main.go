package main

import (
	"github.com/gin-gonic/gin"
	"go-rest/rest/library"
)

func main() {
	router := gin.Default()
	router.GET("/books", library.GetAllBooks)
	router.GET("/books/available", library.GetAllAvailableBooks)
	router.POST("/books", library.PostBook)
	router.Run("localhost:8080")
}
