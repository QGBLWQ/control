package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type AnalysisController struct{}

// 线性回归预测接口
// 输入的数据格式：
//
//	{
//	   "x_data": [1.0, 2.0, 3.0, 4.0, 5.0],
//	   "y_data": [2.0, 4.0, 5.0, 4.0, 5.0],
//	   "predict_data": [6.0, 7.0]
//	}
func (ctrl *AnalysisController) AnalysisLinearRegress(c *gin.Context) {
	// 接收输入数据
	var inputData map[string]interface{}
	if err := c.ShouldBindJSON(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	// 将输入数据转换为 JSON 字符串
	jsonData, err := json.Marshal(inputData)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}

	// 调用 Python 脚本进行线性回归预测
	currentDir, err := os.Getwd()
	if err != nil {
		log.Println("Error getting current directory:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}
	scriptPath := filepath.Join(currentDir, "controller", "analysis", "linear_regression.py")

	cmd := exec.Command("python", scriptPath, string(jsonData))
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error executing Python script: %s, output: %s", err, string(output))
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}

	// 解析 Python 脚本返回的数据
	var predictions []float64
	if err := json.Unmarshal(output, &predictions); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error parsing Python response"})
		return
	}

	// 返回预测结果
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    predictions,
	})
}

// ARIMA预测接口
// 输入的数据格式：
//
//	{
//	   "data": [123.0, 150.5, 180.7, 210.2, 170.1, 160.3],
//	   "p": 1,
//	   "d": 1,
//	   "q": 1,
//	   "forecast_steps": 5
//	}
func (ctrl *AnalysisController) AnalysisARIMA(c *gin.Context) {
	// 接收输入数据
	var inputData map[string]interface{}
	if err := c.ShouldBindJSON(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	// 将输入数据转换为 JSON 字符串
	jsonData, err := json.Marshal(inputData)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}

	// 调用 Python 脚本进行 ARIMA 预测
	currentDir, err := os.Getwd()
	if err != nil {
		log.Println("Error getting current directory:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}
	scriptPath := filepath.Join(currentDir, "controller", "analysis", "ARIMA.py")

	cmd := exec.Command("python", scriptPath, string(jsonData))
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error executing Python script: %s, output: %s", err, string(output))
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}

	log.Printf("Python script output: %s", string(output))

	// 解析 Python 脚本返回的数据
	var predictions []float64
	if err := json.Unmarshal(output, &predictions); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error parsing Python response"})
		return
	}

	// 返回预测结果
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    predictions,
	})
}

// AnalysisGreyPredict 灰色预测接口
// 输入的数据格式：
//
//	{
//	   "data": [120.5, 132.1, 149.7, 165.3, 178.8, 190.2],
//	   "forecast_steps": 5
//	}
func (ctrl *AnalysisController) AnalysisGreyPredict(c *gin.Context) {
	// 接收输入数据
	var inputData map[string]interface{}
	if err := c.ShouldBindJSON(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	// 将输入数据转换为 JSON 字符串
	jsonData, err := json.Marshal(inputData)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}

	// 调用 Python 脚本进行灰色预测
	currentDir, err := os.Getwd()
	if err != nil {
		log.Println("Error getting current directory:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}
	scriptPath := filepath.Join(currentDir, "controller", "analysis", "grey_predict.py")

	cmd := exec.Command("python", scriptPath, string(jsonData))
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error executing Python script: %s, output: %s", err, string(output))
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}

	// 解析 Python 脚本返回的数据
	var predictions []float64
	if err := json.Unmarshal(output, &predictions); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error parsing Python response"})
		return
	}

	// 返回预测结果
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    predictions,
	})
}

// AnalysisBP BP神经网络预测接口
// 输入的数据格式：
//
//	{
//		"data": [120.5, 132.1, 149.7, 165.3, 178.8, 190.2],
//		"forecast_steps": 5,
//		"hidden_layers": 100,
//		"max_iter": 1000,
//		"learning_rate_init": 0.01
//	}
func (ctrl *AnalysisController) AnalysisBP(c *gin.Context) {
	// 接收输入数据
	var inputData map[string]interface{}
	if err := c.ShouldBindJSON(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	// 将输入数据转换为 JSON 字符串
	jsonData, err := json.Marshal(inputData)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}

	// 调用 Python 脚本进行 BP 神经网络预测
	currentDir, err := os.Getwd()
	if err != nil {
		log.Println("Error getting current directory:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}
	scriptPath := filepath.Join(currentDir, "controller", "analysis", "BP.py")

	cmd := exec.Command("python", scriptPath, string(jsonData))
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error executing Python script: %s, output: %s", err, string(output))
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}

	// 解析 Python 脚本返回的数据
	var predictions []float64
	if err := json.Unmarshal(output, &predictions); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error parsing Python response"})
		return
	}

	// 返回预测结果
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    predictions,
	})
}

// AnalysisOverview 总览功能接口，计算描述性统计信息
// 输入的数据格式：
//
//	{
//	   "data": [120.5, 132.1, 149.7, 165.3, 178.8, 190.2]
//	}
func (ctrl *AnalysisController) AnalysisOverview(c *gin.Context) {
	// 接收输入数据
	var inputData map[string]interface{}
	if err := c.ShouldBindJSON(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	// 将输入数据转换为 JSON 字符串
	jsonData, err := json.Marshal(inputData)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}

	// 调用 Python 脚本进行描述性统计计算
	currentDir, err := os.Getwd()
	if err != nil {
		log.Println("Error getting current directory:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}
	scriptPath := filepath.Join(currentDir, "controller", "analysis", "overall.py")

	cmd := exec.Command("python", scriptPath, string(jsonData))
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error executing Python script: %s, output: %s", err, string(output))
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}

	// 解析 Python 脚本返回的数据
	var stats map[string]interface{}
	if err := json.Unmarshal(output, &stats); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error parsing Python response"})
		return
	}

	// 返回统计结果
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    stats,
	})
}
func (ctrl *AnalysisController) AnalysisSVM(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "test",
	})
}
func (ctrl *AnalysisController) AnalysisRandomForest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "test",
	})
}
