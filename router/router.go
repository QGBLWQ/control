package router

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/Heath000/fzuSE2024/controller"
	"github.com/Heath000/fzuSE2024/middleware"
	"github.com/Heath000/fzuSE2024/model"
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
    app.POST("/chart/pie", chartController.ChartPie)
    app.POST("/chart/bar", chartController.ChartBar)
    app.POST("/chart/line", chartController.ChartLine)
	app.POST("/chart/linebarmixed", chartController.ChartLineBarMixed)

	analysisController := new(controller.AnalysisController)
    app.POST("/analysis/overall", analysisController.AnalysisOverview)
    app.POST("/analysis/linear_regression", analysisController.AnalysisLinearRegress)
	app.POST("/analysis/grey_predict", analysisController.AnalysisGreyPredict)
	app.POST("/analysis/ARIMA", analysisController.AnalysisARIMA)
	app.POST("/analysis/BP", analysisController.AnalysisBP)

	dataProcessingController := new(controller.DataProcessingController)
    app.POST("/data/standardize", dataProcessingController.DataStandalize)
    app.POST("/data/outliers", dataProcessingController.DataOutliersHandle)
	app.POST("/data/outliers", dataProcessingController.DataMissingValuesHandle)
	app.POST("/data/outliers", dataProcessingController.DataFeatureVariance)
	app.POST("/data/outliers", dataProcessingController.DataFeatureCorrelation)
	app.POST("/data/outliers", dataProcessingController.DataFeatureChiSquare)

	queryController := new(controller.QueryController)
    app.GET("/query/Region", queryController.QueryRegion)
    app.GET("/query/Category", queryController.QueryCategory)
	app.GET("/query/BasicTable", queryController.QueryBasicTable)
}
