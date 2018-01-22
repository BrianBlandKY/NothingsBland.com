package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	c "NothingsBland.com/web/config"
	"NothingsBland.com/web/controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/cache"
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
	configFile := flag.String("config", "", "NothingsBland configuration file")
	cfg := c.BuildConfig(*configFile)

	app := iris.New()
	app.Logger().SetLevel(cfg.Server.LogLevel)

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	if cfg.Server.EnableCaching {
		cacheHandler := cache.Handler(24 * time.Hour)
		app.Use(cacheHandler)
	}

	// File Logging
	if cfg.Server.EnableLogging {
		f := newLogFile()
		defer f.Close()
		app.Logger().SetOutput(newLogFile())
	}

	// Environment
	app.OnErrorCode(iris.StatusNotFound, notFoundHandler)
	app.Favicon("./public/images/favicon.ico")
	app.StaticWeb("/assets", "./public/assets")

	// Templates
	tmpl := iris.HTML("./templates", ".html").Layout("layout.html")
	tmpl.Reload(cfg.Server.RebuildTemplates)
	app.RegisterView(tmpl)

	// Controllers
	mvc.New(app).Handle(new(controllers.MainController))

	// Run
	app.Run(iris.Addr(fmt.Sprintf(":%v", cfg.Server.Port)),
		iris.WithoutServerError(iris.ErrServerClosed))
	// LetsEncrypt
	// app.Run(iris.AutoTLS(":443", "example.com", "mail@example.com"))
}
