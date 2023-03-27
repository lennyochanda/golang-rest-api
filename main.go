package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lennyochanda/golang-rest-api/db"
	"github.com/lennyochanda/golang-rest-api/routes"
)

func main() {
	fmt.Println("Connecting ...")

	database := db.Database_connection()
	defer database.Close()

	routes.SetupRoutes()
}