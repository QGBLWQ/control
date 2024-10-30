package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
type DataProcessingController struct{}

func (ctrl *DataProcessingController) DataStandalize(c *gin.Context) {//数据标准化
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "test",
	})
}
func (ctrl *DataProcessingController) DataOutliersHandle(c *gin.Context) {//异常值处理生成
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "test",
	})
}
func (ctrl *DataProcessingController) DataMissingValuesHandle(c *gin.Context) {//异常值处理生成
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "test",
	})
}
func (ctrl *DataProcessingController) DataFeatureCorrelation(c *gin.Context) {//异常值处理生成
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "test",
	})
}
func (ctrl *DataProcessingController) DataFeatureVariance(c *gin.Context) {//异常值处理生成
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "test",
	})
}
func (ctrl *DataProcessingController) DataFeatureChiSquare(c *gin.Context) {//异常值处理生成
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "test",
	})
}
//还有缺失值处理	特征筛选-相关系数法、方差选择法、卡方检验法	 