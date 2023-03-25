package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Book struct {
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author string `json:"author"`
    Price float64 `json:"price"`
}

func load_env_variables(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func database_connection() *sql.DB {
	user := load_env_variables("db_user")
	pass := load_env_variables("db_pass")

	db, err := sql.Open("mysql", user+":"+pass+"@tcp(127.0.0.1:3306)/bookstore")
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}

func fetch_books(db *sql.DB) {
	//can further be improved to take a limit and offset
	results, err := db.Query("SELECT * FROM books")
	if err != nil {
		panic(err.Error())
	}
	
	for results.Next() {
		var book Book
		err = results.Scan(&book.Isbn, &book.Title, &book.Author, &book.Price)
		if err != nil {
			panic(err.Error())
		}
	
		log.Printf("Title: %s \n Isbn: %s \n Author: %s \n Price: %f\n", book.Title, book.Isbn, book.Author, book.Price)
	}

}

// func insert_book(db *sql.DB, book Book) {
// 	insert, err := db.Query("INSERT INTO books VALUES (" + book.Isbn + "," + book.Title + "," + book.Author + "," + strconv.FormatFloat(book.Price, 'E', 4, 32) + ")")
// 	if err != nil {
// 		panic(err.Error())
// 	}
	
// 	defer insert.Close()
// }

func main() {
	fmt.Println("Connecting ...")

	db := database_connection()
	defer db.Close()

	fetch_books(db)
}