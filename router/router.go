package router

import (
	"github.com/Heath000/fzuSE2024/controller"
	"github.com/Heath000/fzuSE2024/middleware"
	"github.com/Heath000/fzuSE2024/model"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// Route makes the routing
func Route(app *gin.Engine) {
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
	}

	dataProcessingController := new(controller.DataProcessingController)
	data := app.Group("/data")
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
		query.GET("/region", queryController.QueryRegion)
		query.GET("/category", queryController.QueryCategory)
		query.GET("/basic_table", queryController.QueryBasicTable)
	}

}
