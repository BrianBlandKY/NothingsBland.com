package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	c "NothingsBland.com/web/config"
	"github.com/julienschmidt/httprouter"
)

// NothingsBlandServer -
type NothingsBlandServer struct {
	cfg    c.Config
	logger *log.Logger
	router *httprouter.Router
}

func (s *NothingsBlandServer) newLogFile() *os.File {
	filename := time.Now().Format("Jan 02 2006")
	// open an output file, this will append to the today's file if server restarted.
	logFileName := fmt.Sprintf("%s/%s", s.cfg.Server.LogDirectory, filename)
	f, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}

func (s *NothingsBlandServer) log(format string, data ...interface{}) {
	s.logger.Printf(format, data...)
}

// Setup -
func (s *NothingsBlandServer) Setup() {
	// Loggiing
	if s.cfg.Server.Environment == "dev" {
		s.logger = log.New(os.Stderr, time.Now().Format("2006-01-02 15:04:05")+" - NothingsBlandServer - ", 0)
	} else {
		s.logger = log.New(s.newLogFile(), time.Now().Format("2006-01-02 15:04:05")+" - NothingsBlandServer - ", 0)
	}

	s.router = httprouter.New()

	// Index Handler
	// Returns a single HTML page
	indexHandler := func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		t, _ := template.ParseFiles(fmt.Sprintf("%s/index.html", s.cfg.Server.AssetsDirectory))
		t.Execute(w, nil)
	}

	// Not Found Handler
	// Logs and redirects all not found errors to the index page.
	notFoundHandler := func(w http.ResponseWriter, r *http.Request) {
		s.log("Not Found %v \n", r.URL)
		http.Redirect(w, r, "/", 301)
	}

	// Panic Handler
	// Logs panic errors and redirects to the root
	panicHandler := func(w http.ResponseWriter, r *http.Request, data interface{}) {
		s.log("Error %v \n", data)
		http.Redirect(w, r, "/", 301)
	}

	s.router.NotFound = notFoundHandler
	s.router.PanicHandler = panicHandler
	s.router.GET("/", indexHandler)
	s.router.ServeFiles("/assets/*filepath", http.Dir(s.cfg.Server.AssetsDirectory))
}

// Run -
func (s *NothingsBlandServer) Run() {
	log.Printf("Server running @ :%s \n", s.cfg.Server.Port)
	s.log("Server running @ :%s \n", s.cfg.Server.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", s.cfg.Server.Port), s.router))
}

func main() {
	configFile := flag.String("config", "app.dev.yaml", "NothingsBland configuration file")
	flag.Parse()
	cfg := c.GetConfig(*configFile)

	server := NothingsBlandServer{
		cfg: cfg,
	}
	server.Setup()
	server.Run()
}
