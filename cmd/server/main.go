package main

import (
	"database/sql"
	"log"

	"github.com/Okabe-Junya/api-template-go/internal/db"
	"github.com/Okabe-Junya/api-template-go/internal/handlers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// initialize db connection
	conn, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/db")
	if err != nil {
		log.Fatal("Error opening database:", err)
	} else {
		log.Println("Success opening database")
		// try to connect to the server
		err = conn.Ping()
		if err != nil {
			log.Fatal("Error connecting to the server:", err)
		}
	}
	defer conn.Close()

	// initialize db queries
	query := db.New(conn)

	// initialize router
	router := gin.Default()

	// initialize handlers
	userHandler := handlers.NewUserHandler(query)

	// register routes
	router.GET("/", handlers.SampleHandler)

	router.POST("/users", userHandler.CreateUser)
	router.GET("/users/:id", userHandler.GetUser)
	router.DELETE("/users/:id", userHandler.DeleteUser)

	// run server
	router.Run(":8080")
}
