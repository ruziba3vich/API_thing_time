package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/time_project/backend/internal/models"
	"github.com/time_project/backend/internal/services"
)

func GetResult(c *gin.Context) {
	var time models.RFC3339
	if err := c.ShouldBindJSON(&time); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	result, err := services.GetResult(time)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": result})
}
