package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// 
type LlmController struct{}

// GetUser gets the user info
func (ctrl *LlmController) GetReport(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "test",
	})
}