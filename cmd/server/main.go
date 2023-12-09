package main

import (
	"database/sql"
	"log"

	"github.com/Okabe-Junya/api-template-go/internal/handlers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// initialize sql
	dsn := "user:password@tcp(db:3306)/mydatabase?parseTime=true"
	db, err := sql.Open("mysql", dsn)
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
