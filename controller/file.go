package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Heath000/fzuSE2024/model"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// FileController 处理文件相关的操作
type FileController struct{}

// GetFileList 获取用户的文件列表
func (f *FileController) GetFileList(c *gin.Context) {
	// 从 JWT 中提取 claims 并获取用户 ID
	claims := jwt.ExtractClaims(c)
	idValue, ok := claims["ID"].(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "User ID not found in JWT",
		})
		return
	}

	// 将提取的 ID 转换为 uint 类型
	userID := uint(idValue)

	// 使用提取的 userID 获取用户的文件列表
	var fileModel model.File
	files, err := fileModel.GetFileListByUserId(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to retrieve files",
		})
		return
	}

	// 返回文件列表
	c.JSON(http.StatusOK, gin.H{
		"files": files,
	})
}

// GetFile 获取单个文件信息
func (f *FileController) GetFile(c *gin.Context) {
	// 从请求中获取文件ID
	fileIDStr := c.Param("file_id")
	fileID, err := strconv.ParseUint(fileIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid file ID format",
		})
		return
	}

	// 从 JWT 中提取 claims 并获取用户 ID
	claims := jwt.ExtractClaims(c)
	idValue, ok := claims["ID"].(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "User ID not found in JWT",
		})
		return
	}
	userID := uint(idValue)

	// 根据 file_id 和 userID 获取该文件的信息
	var fileModel model.File
	file, err := fileModel.GetFileByIDAndUserID(uint(fileID), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "File not found",
		})
		return
	}

	// 真正地获取服务器器里的文件
	c.JSON(http.StatusOK, gin.H{
		"file": file,
	})
}

// DeleteFile 删除文件
func (f *FileController) DeleteFile(c *gin.Context) {
	// 从请求中获取文件ID
	fileIDStr := c.Param("file_id")
	fileID, err := strconv.ParseUint(fileIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid file ID format",
		})
		return
	}

	// 从 JWT 中提取 claims 并获取用户 ID
	claims := jwt.ExtractClaims(c)
	idValue, ok := claims["ID"].(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "User ID not found in JWT",
		})
		return
	}
	userID := uint(idValue)

	// 删除文件，使用 fileID 和 userID 双重验证
	var fileModel model.File
	err = fileModel.DeleteFileByIDAndUserID(uint(fileID), userID)
	if err != nil {
		if err == model.ErrDataNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": "File not found or access denied",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "Failed to delete file",
			})
		}
		return
	}

	// 真正地删除服务器里面存的文件***********代码写在这里

	c.JSON(http.StatusOK, gin.H{
		"message": "File deleted successfully",
	})
}

// UploadFile 上传文件
func (f *FileController) UploadFile(c *gin.Context) {
	// 从 JWT 中提取 claims 并获取用户 ID
	claims := jwt.ExtractClaims(c)
	idValue, ok := claims["ID"].(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "User ID not found in JWT",
		})
		return
	}
	userID := uint(idValue)

	// 获取上传的文件
	file, _ := c.FormFile("file")
	if file == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "No file uploaded",
		})
		return
	}

	// 将文件保存到指定路径
	// ********************** 代码写在这里

	// 在数据库中保存文件记录
	fileInfo := model.File{
		UserID:     userID,
		Filename:   file.Filename,
		UploadTime: time.Now(), // 假设 UploadTime 自动设置
	}
	if err := fileInfo.PostFileInfo(userID, file.Filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to save file information in database",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully",
	})
}
