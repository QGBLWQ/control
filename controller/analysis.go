package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
type AnalysisController struct{}

func (ctrl *AnalysisController) AnalysisOverview(c *gin.Context) {//总览
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "test",
	})
}
func (ctrl *AnalysisController) AnalysisLinearRegress(c *gin.Context) {//线性回归预测
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "test",
	})
}
func (ctrl *AnalysisController) AnalysisARIMA(c *gin.Context) {//ARIMA预测
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "test",
	})
}
func (ctrl *AnalysisController) AnalysisGreyPredict(c *gin.Context) {//灰色预测
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "test",
	})
}
func (ctrl *AnalysisController) AnalysisBP(c *gin.Context) {//BP神经网络预测
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "test",
	})
}