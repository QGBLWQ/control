package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

type DataProcessingController struct{}

// DataStandalize 数据标准化接口
// 输入的数据格式为：
//
//	{
//	    "data": [12.5, 18.7, 11.4, 19.1, 15.3]
//	}
//
// 输出数据：
//
//	{
//	    "message": "success",
//	    "data": [0.25, 1.0, 0.0, 1.25, 0.75]
//	}
func (ctrl *DataProcessingController) DataStandalize(c *gin.Context) {
	// 获取用户传递的 JSON 数据
	var inputData map[string]interface{}
	if err := c.ShouldBindJSON(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	// 提取输入数据中的 "data" 字段
	data, ok := inputData["data"].([]interface{})
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data format"})
		return
	}

	// 将数据转成浮动数组
	var floatData []float64
	for _, v := range data {
		if num, ok := v.(float64); ok {
			floatData = append(floatData, num)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid number in data"})
			return
		}
	}

	// 将浮点数据转换为 JSON 字符串
	jsonData, err := json.Marshal(floatData)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}

	// 将数据传递给 Python 脚本进行标准化处理

	// 获取当前工作目录并拼接脚本路径
	currentDir, err := os.Getwd()
	if err != nil {
		log.Println("Error getting current directory:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}
	scriptPath := filepath.Join(currentDir, "controller", "DataProcessPyScripts", "standardize.py")

	println(string(jsonData))
	cmd := exec.Command("python", scriptPath, string(jsonData))

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error executing Python script: %s, output: %s", err, string(output))
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}

	// 解析 Python 脚本返回的数据
	var standardizedData []float64
	if err := json.Unmarshal(output, &standardizedData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error parsing Python response"})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    standardizedData,
	})
}

// DataOutliersHandle 异常值处理
// 输入 JSON 格式示例：
//
//	{
//	    "data": [12.5, 18.7, 11.4, 200.1, 15.3]
//	}
//
// 输出 JSON 格式示例：
//
//	{
//	    "message": "success",
//	    "data": [12.5, 18.7, 11.4, null, 15.3]  // 假设200.1被识别为异常值并处理
//	}
func (ctrl *DataProcessingController) DataOutliersHandle(c *gin.Context) {
	// 获取用户传递的 JSON 数据
	var inputData map[string]interface{}
	if err := c.ShouldBindJSON(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	// 提取输入数据中的 "data" 字段
	data, ok := inputData["data"].([]interface{})
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data format"})
		return
	}

	// 将数据转换为浮动数组
	var floatData []float64
	for _, v := range data {
		if num, ok := v.(float64); ok {
			floatData = append(floatData, num)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid number in data"})
			return
		}
	}

	// 将浮动数据转换为 JSON 字符串
	jsonData, err := json.Marshal(floatData)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}

	// 获取当前工作目录并拼接脚本路径
	currentDir, err := os.Getwd()
	if err != nil {
		log.Println("Error getting current directory:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}
	scriptPath := filepath.Join(currentDir, "controller", "DataProcessPyScripts", "handle_outliers.py")

	println(string(jsonData))
	cmd := exec.Command("python", scriptPath, string(jsonData))

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error executing Python script: %s, output: %s", err, string(output))
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}

	// 解析 Python 脚本返回的数据
	var processedData []interface{}
	if err := json.Unmarshal(output, &processedData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error parsing Python response"})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    processedData,
	})
}

// DataMissingValuesHandle 缺失值处理
// 输入 JSON 格式示例：
//
//	{
//	    "data": [12.5, null, 11.4, null, 15.3]
//	}
//
// 输出 JSON 格式示例：
//
//	{
//	    "message": "success",
//	    "data": [12.5, 13.0667, 11.4, 13.0667, 15.3]  // 假设缺失值用均值填充
//	}
func (ctrl *DataProcessingController) DataMissingValuesHandle(c *gin.Context) {
	// 获取用户传递的 JSON 数据
	var inputData map[string]interface{}
	if err := c.ShouldBindJSON(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	// 提取输入数据中的 "data" 字段
	data, ok := inputData["data"].([]interface{})
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data format"})
		return
	}

	// 将数据转换为浮动数组（将 null 转为 NaN）
	var floatData []interface{}
	for _, v := range data {
		if v == nil {
			floatData = append(floatData, nil) // 保留 nil 值作为缺失值
		} else if num, ok := v.(float64); ok {
			floatData = append(floatData, num)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid number in data"})
			return
		}
	}

	// 将浮点数据转换为 JSON 字符串
	jsonData, err := json.Marshal(floatData)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}

	// 获取当前工作目录并拼接脚本路径
	currentDir, err := os.Getwd()
	if err != nil {
		log.Println("Error getting current directory:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}
	scriptPath := filepath.Join(currentDir, "controller", "DataProcessPyScripts", "handle_missing_values.py")

	println(string(jsonData))
	cmd := exec.Command("python", scriptPath, string(jsonData))

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error executing Python script: %s, output: %s", err, string(output))
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}

	// 解析 Python 脚本返回的数据
	var processedData []interface{}
	if err := json.Unmarshal(output, &processedData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error parsing Python response"})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    processedData,
	})
}

// DataFeatureCorrelation 特征相关性计算
// 输入 JSON 格式示例：
//
//	{
//	    "data": [
//	        [1.0, 2.0, 3.0],
//	        [2.0, 3.0, 4.0],
//	        [3.0, 4.0, 5.0]
//	    ]
//	}
//
// 输出 JSON 格式示例：
//
//	{
//	    "message": "success",
//	    "data": [
//	        [1.0, 0.99, 0.95],
//	        [0.99, 1.0, 0.97],
//	        [0.95, 0.97, 1.0]
//	    ]
//	}
func (ctrl *DataProcessingController) DataFeatureCorrelation(c *gin.Context) {
	// 获取用户传递的 JSON 数据
	var inputData map[string]interface{}
	if err := c.ShouldBindJSON(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	// 提取输入数据中的 "data" 字段
	data, ok := inputData["data"].([]interface{})
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data format"})
		return
	}

	// 将数据转换为二维浮点数组
	var floatData [][]float64
	for _, row := range data {
		rowSlice, ok := row.([]interface{})
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data format"})
			return
		}
		var floatRow []float64
		for _, v := range rowSlice {
			if num, ok := v.(float64); ok {
				floatRow = append(floatRow, num)
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid number in data"})
				return
			}
		}
		floatData = append(floatData, floatRow)
	}

	// 将二维浮点数组转换为 JSON 字符串
	jsonData, err := json.Marshal(floatData)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}

	// 获取当前工作目录并拼接脚本路径
	currentDir, err := os.Getwd()
	if err != nil {
		log.Println("Error getting current directory:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}
	scriptPath := filepath.Join(currentDir, "controller", "DataProcessPyScripts", "calculate_correlation.py")

	println(string(jsonData))
	cmd := exec.Command("python", scriptPath, string(jsonData))

	// 执行 Python 脚本计算特征相关系数
	//cmd := exec.Command("python", "C:\\Users\\19872\\Desktop\\ruangong\\fzuSE2024-main\\controller\\DataProcessPyScripts\\calculate_correlation.py", string(jsonData))
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error executing Python script: %s, output: %s", err, string(output))
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}

	// 解析 Python 脚本返回的数据
	var correlationMatrix [][]float64
	if err := json.Unmarshal(output, &correlationMatrix); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error parsing Python response"})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    correlationMatrix,
	})
}

// DataFeatureVariance 特征方差计算
// 输入 JSON 格式示例：
//
//	{
//	    "data": [
//	        [1.0, 2.0, 3.0],
//	        [2.0, 3.0, 4.0],
//	        [3.0, 4.0, 5.0]
//	    ]
//	}
//
// 输出 JSON 格式示例：
//
//	{
//	    "message": "success",
//	    "data": [0.6667, 0.6667, 0.6667]  // 每列的方差
//	}
func (ctrl *DataProcessingController) DataFeatureVariance(c *gin.Context) {
	// 获取用户传递的 JSON 数据
	var inputData map[string]interface{}
	if err := c.ShouldBindJSON(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	// 提取输入数据中的 "data" 字段
	data, ok := inputData["data"].([]interface{})
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data format"})
		return
	}

	// 将数据转换为二维浮点数组
	var floatData [][]float64
	for _, row := range data {
		rowSlice, ok := row.([]interface{})
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data format"})
			return
		}
		var floatRow []float64
		for _, v := range rowSlice {
			if num, ok := v.(float64); ok {
				floatRow = append(floatRow, num)
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid number in data"})
				return
			}
		}
		floatData = append(floatData, floatRow)
	}

	// 将二维浮点数组转换为 JSON 字符串
	jsonData, err := json.Marshal(floatData)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}

	// 获取当前工作目录并拼接脚本路径
	currentDir, err := os.Getwd()
	if err != nil {
		log.Println("Error getting current directory:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}
	scriptPath := filepath.Join(currentDir, "controller", "DataProcessPyScripts", "calculate_variance.py")

	println(string(jsonData))
	cmd := exec.Command("python", scriptPath, string(jsonData))

	// 执行 Python 脚本计算方差
	//cmd := exec.Command("python", "C:\\Users\\19872\\Desktop\\ruangong\\fzuSE2024-main\\controller\\DataProcessPyScripts\\calculate_variance.py", string(jsonData))
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error executing Python script: %s, output: %s", err, string(output))
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}

	// 解析 Python 脚本返回的数据
	var variances []float64
	if err := json.Unmarshal(output, &variances); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error parsing Python response"})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    variances,
	})
}

// DataFeatureChiSquare 特征卡方检验
// 输入 JSON 格式示例：
//
//	{
//	    "data": [
//	        [1, 0, 1],
//	        [0, 1, 1],
//	        [1, 0, 0],
//	        [0, 1, 1]
//	    ]
//	}
//
// 输出 JSON 格式示例：
//
//	{
//	    "message": "success",
//	    "data": [
//	        [0, 5.0, 3.2],
//	        [5.0, 0, 4.1],
//	        [3.2, 4.1, 0]
//	    ]
//	}
func (ctrl *DataProcessingController) DataFeatureChiSquare(c *gin.Context) {
	// 获取用户传递的 JSON 数据
	var inputData map[string]interface{}
	if err := c.ShouldBindJSON(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	// 提取输入数据中的 "data" 字段
	data, ok := inputData["data"].([]interface{})
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data format"})
		return
	}

	// 将数据转换为二维浮点数组
	var intData [][]int
	for _, row := range data {
		rowSlice, ok := row.([]interface{})
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data format"})
			return
		}
		var intRow []int
		for _, v := range rowSlice {
			if num, ok := v.(float64); ok {
				intRow = append(intRow, int(num))
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid number in data"})
				return
			}
		}
		intData = append(intData, intRow)
	}

	// 将二维整数数组转换为 JSON 字符串
	jsonData, err := json.Marshal(intData)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}
	// 获取当前工作目录并拼接脚本路径
	currentDir, err := os.Getwd()
	if err != nil {
		log.Println("Error getting current directory:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}
	scriptPath := filepath.Join(currentDir, "controller", "DataProcessPyScripts", "calculate_chi_square.py")

	println(string(jsonData))
	cmd := exec.Command("python", scriptPath, string(jsonData))

	//cmd := exec.Command("python", "C:\\Users\\19872\\Desktop\\ruangong\\fzuSE2024-main\\controller\\DataProcessPyScripts\\calculate_chi_square.py", string(jsonData))
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error executing Python script: %s, output: %s", err, string(output))
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error processing data"})
		return
	}

	// 解析 Python 脚本返回的数据
	var chiSquareMatrix [][]float64
	if err := json.Unmarshal(output, &chiSquareMatrix); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error parsing Python response"})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    chiSquareMatrix,
	})
}

//还有缺失值处理	特征筛选-相关系数法、方差选择法、卡方检验法
