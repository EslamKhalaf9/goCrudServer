package controllers

import (
	"bookstorev2/models"
	customtypes "bookstorev2/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func HandleGetBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

func HandleGetBookById(c *gin.Context) {

	var bookSearch customtypes.GetBook
	var book models.Book
	var numOfRows int64

	if err := c.ShouldBindUri(&bookSearch); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	result := models.DB.First(&book, "ID=?", bookSearch.ID)
	result.Count(&numOfRows)

	if numOfRows != 0 {
		c.JSON(http.StatusOK, gin.H{"data": book})
		return
	}

	c.JSON(404, gin.H{"msg": "no books found"})
}

func HandleCreateBook(c *gin.Context) {
	body := customtypes.CreateBook{}

	// validating and populating body variable
	if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't bind request body"})
	} else {
		// body is fine
		// create a book instance
		book := models.Book{Title: body.Title, Author: body.Author}

		result := models.DB.Create(&book)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		}

		fmt.Println(result.RowsAffected)
		c.JSON(http.StatusOK, body)
	}
}

func HandleDeleteBook(c *gin.Context) {
	var bookSearch customtypes.GetBook
	var book models.Book

	if err := c.ShouldBindUri(&bookSearch); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	result := models.DB.Delete(&book, bookSearch.ID)
	if result.Error != nil {
		c.JSON(404, gin.H{"msg": result.Error})
		return
	}

	fmt.Println(result.RowsAffected)

	if result.RowsAffected >= 1 {
		c.JSON(200, gin.H{"msg": "deleted successfully!"})
		return
	}

	c.JSON(404, gin.H{"msg": "record not found"})

}

func HandleUpdateBook(c *gin.Context) {

}
