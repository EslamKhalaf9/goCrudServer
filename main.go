package main

import (
	"bookstorev2/controllers"
	"bookstorev2/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDB()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/book", controllers.HandleGetBooks)
	r.GET("/book/:id", controllers.HandleGetBookById)
	r.DELETE("/book/:id", controllers.HandleDeleteBook)
	r.POST("/book", controllers.HandleCreateBook)
	r.PUT("/book/:id", controllers.HandleUpdateBook)
	r.Run() // listen and serve on 0.0.0.0:8080
}
