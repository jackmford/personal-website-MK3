package main

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"jackmitchellfordyce.com/internal/models"
	"jackmitchellfordyce.com/ui"
	"io/fs"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

type templateData struct {
	Posts models.BlogPosts
	Post  *models.BlogPost
}

var successCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "home_success_request_count",
		Help: "No of requests handled successfully by home handler",
	},
)

var totalCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "home_total_request_count",
		Help: "No of request handled by home handler",
	},
)

func main() {
	router := httprouter.New()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	// Create sub-filesystems for different types of static files
	cssFS, err := fs.Sub(ui.Files, "static/css")
	if err != nil {
		errorLog.Fatal(err)
	}

	imgFS, err := fs.Sub(ui.Files, "static/img")
	if err != nil {
		errorLog.Fatal(err)
	}

	// Serve CSS files
	cssServer := http.FileServer(http.FS(cssFS))
	router.Handler(http.MethodGet, "/static/css/*filepath", http.StripPrefix("/static/css/", cssServer))

	// Serve image files
	imgServer := http.FileServer(http.FS(imgFS))
	router.Handler(http.MethodGet, "/static/img/*filepath", http.StripPrefix("/static/img/", imgServer))

	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/blog", app.blog)
	router.HandlerFunc(http.MethodGet, "/blog/:slug", app.blogPost)
	router.HandlerFunc(http.MethodGet, "/health", health)

	prometheus.MustRegister(totalCounter)
	prometheus.MustRegister(successCounter)
	router.Handler(http.MethodGet, "/metrics", promhttp.Handler())

	infoLog.Print("Starting server on :4000")
	err = http.ListenAndServe(":4000", router)
	errorLog.Fatal(err)
}
