package library

import (
	"github.com/gin-gonic/gin"
	"go-rest/domain/library"
	"net/http"
)

func GetAllBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, library.DefaultLibrary.List())
}

func GetAllAvailableBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, library.DefaultLibrary.ListAvailable())
}

func PostBook(c *gin.Context) {
	var newBook library.Book
	if err := c.BindJSON(&newBook); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		if library.DefaultLibrary.Exists(newBook.Name) || !newBook.Valid() {
			c.AbortWithStatus(http.StatusConflict)
		} else {
			library.DefaultLibrary.Add(newBook)
			c.Status(http.StatusAccepted)
		}
	}
}
