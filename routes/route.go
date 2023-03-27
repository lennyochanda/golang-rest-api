package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lennyochanda/golang-rest-api/handlers"
)

func SetupRoutes() {
	router := gin.Default()
	router.GET("/books", handlers.Fetch_books)
	router.POST("/books", handlers.Insert_book)

	router.Run("localhost:8080")
}