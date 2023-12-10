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
	itemHandler := handlers.NewItemHandler(query)
	userItemHandler := handlers.NewUserItemHandler(query)

	// register routes
	router.GET("/", handlers.SampleHandler)

	router.POST("/users", userHandler.CreateUser)
	router.GET("/users/:id", userHandler.GetUser)
	router.DELETE("/users/:id", userHandler.DeleteUser)

	router.POST("/items", itemHandler.CreateItem)
	router.GET("/items/:id", itemHandler.GetItem)
	router.DELETE("/items/:id", itemHandler.DeleteItem)

	router.POST("/user_items", userItemHandler.CreateUserItem)
	router.GET("/user_items/:user_id/:item_id", userItemHandler.GetUserItem)
	router.DELETE("/user_items/:user_id/:item_id", userItemHandler.DeleteUserItem)

	// run server
	router.Run(":8080")
}
