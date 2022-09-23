package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"pustaka-api/book"
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Roki Adhytama",
		"bio":  "as a full ngestack dev",
	})
}
func ShowBooks(c *gin.Context) {
	title := c.Query("title")
	name := c.Query("name")

	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"title": title,
	})
}

func AddBook(c *gin.Context){
	var bookInput book.BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {

		for _, e := range err.(validator.ValidationErrors){
			errorMsg := fmt.Sprintf("Error on field %s. conditions: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errorMsg)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
	})
}