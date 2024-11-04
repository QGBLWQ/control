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
	"path/filepath"
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
	cwd, err := os.Getwd()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	scriptPath := filepath.Join(cwd, "controller", "draw_pie_chart.py")
	cmd := exec.Command("python", scriptPath, string(paramsJSON))
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
	cwd, err := os.Getwd()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	scriptPath := filepath.Join(cwd, "controller", "draw_line_chart.py")
	cmd := exec.Command("python", scriptPath, string(paramsJSON))
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
func (ctrl *ChartController) ChartBar(c *gin.Context) {
	var params struct {
		Title string `json:"title"`
		Data  []struct {
			Label string  `json:"label"`
			Value float64 `json:"value"`
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
	cwd, err := os.Getwd()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	scriptPath := filepath.Join(cwd, "controller", "draw_bar_chart.py")
	cmd := exec.Command("python", scriptPath, string(paramsJSON))
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": stderr.String()})
		return
	}

	// Read the generated PNG file
	pngData, err := os.ReadFile("bar_chart.png")
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
func (ctrl *ChartController) ChartLineBarMixed(c *gin.Context) {
	var params struct {
		Title    string `json:"title"`
		LineData []struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"line_data"`
		BarData []struct {
			Label string  `json:"label"`
			Value float64 `json:"value"`
		} `json:"bar_data"`
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
	cwd, err := os.Getwd()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	scriptPath := filepath.Join(cwd, "controller", "draw_line_bar_mixed_chart.py")
	cmd := exec.Command("python", scriptPath, string(paramsJSON))
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": stderr.String()})
		return
	}

	// Read the generated PNG file
	pngData, err := os.ReadFile("line_bar_mixed_chart.png")
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
