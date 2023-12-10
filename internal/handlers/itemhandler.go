package handlers

import (
	"strconv"

	"github.com/Okabe-Junya/api-template-go/internal/db"
	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	Queries *db.Queries
}

func NewItemHandler(queries *db.Queries) *ItemHandler {
	return &ItemHandler{
		Queries: queries,
	}
}

func (h *ItemHandler) CreateItem(c *gin.Context) {
	var item db.CreateItemParams
	if err := c.BindJSON(&item); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}
	if err := h.Queries.CreateItem(c, item); err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

func (h *ItemHandler) GetItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	item, err := h.Queries.GetItem(c, int32(id))
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}
	c.JSON(200, item)
}

func (h *ItemHandler) DeleteItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	if err := h.Queries.DeleteItem(c, int32(id)); err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}
