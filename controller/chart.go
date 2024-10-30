package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
type ChartController struct{}

// GetIndex home page
func (ctrl *ChartController) ChartPie(c *gin.Context) {//饼图生成写在这里
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "test",
	})
}
func (ctrl *ChartController) ChartLine(c *gin.Context) {//折线图生成
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "test",
	})
}
func (ctrl *ChartController) ChartBar(c *gin.Context) {//柱状图图生成
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "test",
	})
}
func (ctrl *ChartController) ChartLineBarMixed(c *gin.Context) {//折柱混合图生成
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "test",
	})
}