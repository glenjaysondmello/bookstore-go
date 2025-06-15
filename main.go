package main

import (
	"github.com/gin-gonic/gin"
	"github.com/glenjaysondmello/bookstore/db"
	"github.com/glenjaysondmello/bookstore/handlers"
)

func main() {
	db.InitDB()

	r := gin.Default()

	r.GET("/books", handlers.GetBooks)
	r.GET("/books/:id", handlers.GetBook)
	r.POST("/books", handlers.CreateBook)
	r.PUT("/books", handlers.UpdateBook)
	r.DELETE("/books", handlers.DeleteBook)

	r.Run(":8080")
}
