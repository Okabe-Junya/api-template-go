// Package handlers provides HTTP request handlers for the API
package handlers

import (
	"strconv"

	"github.com/Okabe-Junya/api-template-go/internal/db"
	"github.com/gin-gonic/gin"
)

// UserHandler handles HTTP requests for user operations
type UserHandler struct {
	Queries *db.Queries
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(queries *db.Queries) *UserHandler {
	return &UserHandler{
		Queries: queries,
	}
}

// CreateUser handles POST requests to create a new user
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user db.CreateUserParams
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"error": "invalid request",
		})
		return
	}
	if err := h.Queries.CreateUser(c, user); err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}

// GetUser handles GET requests to retrieve a user by ID
func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid id"})
		return
	}
	user, err := h.Queries.GetUser(c, int32(id))
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(200, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	})
}

// DeleteUser handles DELETE requests to remove a user by ID
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid id"})
		return
	}
	if err := h.Queries.DeleteUser(c, int32(id)); err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}
