package main

import (
	"flag"
	"log"
	"path/filepath"

	"github.com/Heath000/fzuSE2024/config"
	"github.com/Heath000/fzuSE2024/router"
	"github.com/gin-gonic/gin"
)

func main() {

	addr := flag.String("addr", config.Server.Addr, "Address to listen and serve")
	flag.Parse()

	if config.Server.Mode == gin.ReleaseMode {
		gin.DisableConsoleColor()
	}

	app := gin.Default()

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
