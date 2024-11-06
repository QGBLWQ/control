package main

import (
	"flag"
	"log"
	"path/filepath"
	"time"
	"github.com/Heath000/fzuSE2024/config"
	"github.com/Heath000/fzuSE2024/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	addr := flag.String("addr", config.Server.Addr, "Address to listen and serve")
	flag.Parse()

	if config.Server.Mode == gin.ReleaseMode {
		gin.DisableConsoleColor()
	}

	app := gin.Default()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5500"}, // 前端的地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // 身份验证需要 Authorization 头
		ExposeHeaders:    []string{"Content-Length", "Authorization"},         // 允许前端获取这些响应头
		AllowCredentials: true,                                                // 如果需要携带 cookies 或其他凭证，需要设置为 true
		MaxAge:           12 * time.Hour,                                      // 缓存预检请求结果的时间
	}))
	app.Static("/images", filepath.Join(config.Server.StaticDir, "img"))
	app.StaticFile("/favicon.ico", filepath.Join(config.Server.StaticDir, "img/favicon.ico"))
	app.LoadHTMLGlob(config.Server.ViewDir + "/*")
	app.MaxMultipartMemory = config.Server.MaxMultipartMemory << 20

	router.Route(app)

	// Listen and Serve
	if err := app.Run(*addr); err != nil {
		log.Fatal(err.Error())
	}
}
