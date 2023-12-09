package main

import (
	"database/sql"
	"log"

	"github.com/Okabe-Junya/api-template-go/internal/handlers"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// initialize sql
	db, err := sql.Open("sqlite3", "file:mydatabase.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	defer db.Close()

	// initialize router
	router := gin.Default()
	router.GET("/", handlers.ExampleHandler)

	// run server
	router.Run(":8080")
}
