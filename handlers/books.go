package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lennyochanda/golang-rest-api/db"
	"github.com/lennyochanda/golang-rest-api/models"
)

func FetchBooks(c *gin.Context) {
	//how can this be further improved to take a limit and offset?
	var books [] models.Book

	result, err := db.Database_connection().Query("SELECT * FROM books")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	
	for result.Next() {
		var book models.Book
		err = result.Scan(&book.ISBN, &book.Title, &book.Author, &book.Price)
		if err != nil {
			panic(err.Error())
		}

		books = append(books, book)
	}

	c.IndentedJSON(http.StatusOK, books)
}

func InsertBook(c *gin.Context) {
	var book models.Book

	c.BindJSON(&book);

	result, err := db.Database_connection().Query("INSERT INTO books (isbn, title, author, price) VALUES (?, ?, ?, ?)", book.ISBN, book.Title, book.Author, strconv.FormatFloat(book.Price, 'E', 4, 32))
	
	if err != nil {
		panic(err.Error())
	}
	
	defer result.Close()

	c.IndentedJSON(http.StatusCreated, book)
}

func FetchBookById(c *gin.Context) {	
	stmt, err := db.Database_connection().Prepare("SELECT * FROM books WHERE isbn=?")
	
	if err != nil {
		panic(err.Error())
	}
	
	id := c.Param("id")
	var book models.Book

	error := stmt.QueryRow(id).Scan(&book.ISBN, &book.Title, &book.Author, &book.Price); 
	if error != nil {
		if error == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		}
	}

	c.IndentedJSON(http.StatusOK, book)
}