package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lennyochanda/golang-rest-api/handlers"
)

func SetupRoutes() {
	router := gin.Default()
	router.GET("/books", handlers.FetchBooks)
	router.POST("/books", handlers.InsertBook)

	router.GET("/books/:id", handlers.FetchBookById)
	router.DELETE("books/:id", handlers.DeleteById)

	router.Run("localhost:8080")
}