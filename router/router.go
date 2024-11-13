package router

import (
	"fmt"

	"github.com/Heath000/fzuSE2024/controller"
	"github.com/Heath000/fzuSE2024/middleware"
	"github.com/Heath000/fzuSE2024/model"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// Route makes the routing
func Route(app *gin.Engine) {
	fmt.Println("Routes are registered")
	indexController := new(controller.IndexController)
	app.GET(
		"/", indexController.GetIndex,
	)

	auth := app.Group("/auth")
	authMiddleware := middleware.Auth()
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", func(c *gin.Context) {
			claims := jwt.ExtractClaims(c)
			user, _ := c.Get("email")
			c.JSON(200, gin.H{
				"email": claims["email"],
				"name":  user.(*model.User).Name,
				"text":  "Hello World.",
			})
		})
	}

	userController := new(controller.UserController)
	app.GET(
		"/user/:id", userController.GetUser,
	).GET(
		"/signup", userController.SignupForm,
	).POST(
		"/signup", userController.Signup,
	).GET(
		"/login", userController.LoginForm,
	).POST(
		"/login", authMiddleware.LoginHandler,
	)

	api := app.Group("/api")
	{
		api.GET("/version", indexController.GetVersion)
	}

	chartController := new(controller.ChartController)
	chart := app.Group("/chart")
	{
		chart.POST("/pie", chartController.ChartPie)
		chart.POST("/bar", chartController.ChartBar)
		chart.POST("/line", chartController.ChartLine)
		chart.POST("/linebarmixed", chartController.ChartLineBarMixed)
	}

	analysisController := new(controller.AnalysisController)
	analysis := app.Group("/analysis")
	{
		analysis.POST("/overall", analysisController.AnalysisOverview)
		analysis.POST("/linear_regression", analysisController.AnalysisLinearRegress)
		analysis.POST("/grey_predict", analysisController.AnalysisGreyPredict)
		analysis.POST("/ARIMA", analysisController.AnalysisARIMA)
		analysis.POST("/BP", analysisController.AnalysisBP)
		analysis.POST("/SVM", analysisController.AnalysisSVM)
		analysis.POST("/RandomForest", analysisController.AnalysisRandomForest)
	}

	dataProcessingController := new(controller.DataProcessingController)
	data := app.Group("/dataProcessing")
	{
		data.POST("/standardize", dataProcessingController.DataStandalize)
		data.POST("/outliers", dataProcessingController.DataOutliersHandle)
		data.POST("/missing", dataProcessingController.DataMissingValuesHandle)
		data.POST("/feature/variance", dataProcessingController.DataFeatureVariance)
		data.POST("/feature/correlation", dataProcessingController.DataFeatureCorrelation)
		data.POST("/feature/chi_square", dataProcessingController.DataFeatureChiSquare)
	}

	queryController := new(controller.QueryController)
	query := app.Group("/query")
	{
		query.POST("/region", queryController.QueryRegions)
		query.POST("/top_category", queryController.QueryTopCategories)
		query.POST("/sub_category", queryController.QuerySubCategories)
		query.POST("/available_year", queryController.QueryAvailableYears)
		query.POST("/data", queryController.QueryData)

	}
	llmController := new(controller.LlmController)
	llm := app.Group("/llm")
	{
		llm.POST("/report", llmController.GetReport)
	}

	fileController := new(controller.FileController)
	file := app.Group("/file")
	file.Use(authMiddleware.MiddlewareFunc()) // 所有文件相关的路由都需要鉴权
	{
		file.GET("/get_file_list", fileController.GetFileList)          // 获取文件列表
		file.GET("/get_file/:file_id", fileController.GetFile)          // 获取单个文件
		file.DELETE("/delete_file/:file_id", fileController.DeleteFile) // 删除文件
		file.POST("/upload_file", fileController.UploadFile)            // 上传文件
	}

}
