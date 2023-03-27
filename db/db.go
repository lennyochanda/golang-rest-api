package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load_env_variables(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func Database_connection() *sql.DB {
	user := Load_env_variables("db_user")
	pass := Load_env_variables("db_pass")

	db, err := sql.Open("mysql", user+":"+pass+"@tcp(127.0.0.1:3306)/bookstore")
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}