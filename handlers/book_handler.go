package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/glenjaysondmello/bookstore/db"
	"github.com/glenjaysondmello/bookstore/models"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	err := db.DB.Select(&books, "SELECT * FROM books")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching books"})
		return
	}

	c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	err := db.DB.Get(&book, "SELECT * FROM books WHERE id = $1", id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not Found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	err := db.DB.QueryRowx("INSERT INTO books (title, author, year) VALUES ($1, $2, $3) RETURNING id", book.Title, book.Author, book.Year).Scan(&book.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Insert Failed"})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	idC, _ := strconv.Atoi(id)
	book.ID = idC

	_, err := db.DB.Exec("UPDATE books SET title = $1, author = $2, year = $3 WHERE id = $4", book.Title, book.Author, book.Year, book.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Update the book"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	idC, _ := strconv.Atoi(id)

	_, err := db.DB.Exec("DELETE FROM books WHERE id = $1", idC)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete Failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book Deleted"})
}
