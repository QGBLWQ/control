package controller

import (
	"net/http"

	"github.com/Heath000/fzuSE2024/config"
	"github.com/gin-gonic/gin"
)

// IndexController is the default controller
type IndexController struct{}

// GetIndex home page
func (ctrl *IndexController) GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":   "Gin Skeleton",
		"content": "This is a skeleton based on gin framework",
	})
}

// GetVersion version json
func (ctrl *IndexController) GetVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": config.Server.Version,
	})
}
