package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	c "NothingsBland.com/nothingsbland/config"
	"NothingsBland.com/nothingsbland/controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
)

func todayFilename() string {
	today := time.Now().Format("Jan 02 2006")
	return today + ".txt"
}

func newLogFile() *os.File {
	filename := todayFilename()
	// open an output file, this will append to the today's file if server restarted.
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}

func notFoundHandler(ctx iris.Context) {
	ctx.Redirect("/", iris.StatusTemporaryRedirect)
}

func main() {
	configFile := flag.String("config", "app.yaml", "NothingsBland configuration file")

	// Get app.yaml from command line flag
	cfg := c.ParseConfig(*configFile)

	app := iris.New()
	app.Logger().SetLevel(cfg.Server.LogLevel)

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// File Logging
	// f := newLogFile()
	// defer f.Close()
	// app.Logger().SetOutput(newLogFile())

	// Environment
	// app.Favicon("./public/images/favicon.ico")
	// app.StaticWeb("/assets", "./public/assets")
	app.OnErrorCode(iris.StatusNotFound, notFoundHandler)

	// Controllers
	mvc.New(app).Handle(new(controllers.MainController))

	// Run
	app.Run(iris.Addr(fmt.Sprintf(":%v", cfg.Server.Port)),
		iris.WithoutServerError(iris.ErrServerClosed))
	// LetsEncrypt
	// app.Run(iris.AutoTLS(":443", "example.com", "mail@example.com"))
}
