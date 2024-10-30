package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
type QueryController struct{}

func (ctrl *QueryController) QueryRegion(c *gin.Context) {//查地区表
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "test",
	})
}
func (ctrl *QueryController) QueryCategory(c *gin.Context) {//查多级类目
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "test",
	})
}
func (ctrl *QueryController) QueryBasicTable(c *gin.Context) {//查真正的时间序列数据表
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "test",
	})
}