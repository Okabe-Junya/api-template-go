package handlers

import (
	"strconv"

	"github.com/Okabe-Junya/api-template-go/internal/db"
	"github.com/gin-gonic/gin"
)

// UserItemHandler handles HTTP requests for user-item relationship operations
type UserItemHandler struct {
	Queries *db.Queries
}

// NewUserItemHandler creates a new UserItemHandler instance
func NewUserItemHandler(queries *db.Queries) *UserItemHandler {
	return &UserItemHandler{
		Queries: queries,
	}
}

// CreateUserItem handles POST requests to create a new user-item relationship
func (h *UserItemHandler) CreateUserItem(c *gin.Context) {
	var userItem db.CreateUserItemParams
	if err := c.BindJSON(&userItem); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}
	if err := h.Queries.CreateUserItem(c, userItem); err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

// GetUserItem handles GET requests to retrieve a user-item relationship by user ID and item ID
func (h *UserItemHandler) GetUserItem(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user id"})
		return
	}
	itemID, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid item id"})
		return
	}
	userItem, err := h.Queries.GetUserItem(c, db.GetUserItemParams{UserID: int32(userID), ItemID: int32(itemID)})
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}
	c.JSON(200, userItem)
}

// DeleteUserItem handles DELETE requests to remove a user-item relationship by user ID and item ID
func (h *UserItemHandler) DeleteUserItem(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user id"})
		return
	}
	itemID, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid item id"})
		return
	}
	if err := h.Queries.DeleteUserItem(c, db.DeleteUserItemParams{UserID: int32(userID), ItemID: int32(itemID)}); err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}
