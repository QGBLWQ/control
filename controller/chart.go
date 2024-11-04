package controller

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-echarts/go-echarts/v2/opts"
	"net/http"
	"os"
	"os/exec"
)

type ChartController struct{}

// GetIndex home page
type PieChartParams struct {
	Title string         `json:"title"`
	Data  []opts.PieData `json:"data"`
}

func (ctrl *ChartController) ChartPie(c *gin.Context) {
	var params PieChartParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert params to JSON string
	paramsJSON, err := json.Marshal(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Call the Python script
	cmd := exec.Command("python", "D:\\code\\se\\fzuSE2024\\controller\\draw_pie_chart.py", string(paramsJSON))
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": stderr.String()})
		return
	}

	// Read the generated PNG file
	pngData, err := os.ReadFile("pie_chart.png")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Encode PNG to Base64
	base64Image := base64.StdEncoding.EncodeToString(pngData)

	// Return Base64 image as JSON
	c.JSON(http.StatusOK, gin.H{
		"image": base64Image,
	})
}

func (ctrl *ChartController) ChartLine(c *gin.Context) {
	var params struct {
		Title string `json:"title"`
		Data  []struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"data"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert params to JSON string
	paramsJSON, err := json.Marshal(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Call the Python script
	cmd := exec.Command("python", "D:\\code\\se\\fzuSE2024\\controller\\draw_line_chart.py", string(paramsJSON))
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": stderr.String()})
		return
	}

	// Read the generated PNG file
	pngData, err := os.ReadFile("line_chart.png")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Encode PNG to Base64
	base64Image := base64.StdEncoding.EncodeToString(pngData)

	// Return Base64 image as JSON
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"image":   base64Image,
	})
}
func (ctrl *ChartController) ChartBar(c *gin.Context) { //柱状图图生成
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "test",
	})
}
func (ctrl *ChartController) ChartLineBarMixed(c *gin.Context) { //折柱混合图生成
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "test",
	})
}
