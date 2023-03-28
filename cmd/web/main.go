package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"jackmitchellfordyce.com/ui"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
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

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Healthy"))
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	totalCounter.Inc()
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ts, err := template.ParseFS(ui.Files, "html/pages/index.tmpl")
	if err != nil {
		app.errorLog.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.errorLog.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	successCounter.Inc()
}

func main() {

	router := httprouter.New()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	fileServer := http.FileServer(http.FS(ui.Files))
	router.Handler(http.MethodGet, "/static/*filepath", fileServer)

	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/health", health)

	prometheus.MustRegister(totalCounter)
	prometheus.MustRegister(successCounter)
	router.Handler(http.MethodGet, "/metrics", promhttp.Handler())

	infoLog.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", router)
	errorLog.Fatal(err)
}
