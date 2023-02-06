package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"jackmitchellfordyce.com/ui"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	//pages, err := fs.Glob(ui.Files, "html/pages/*tmpl")
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
	}
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	//fileServer := http.FileServer(http.Dir("./ui/static/"))

	fileServer := http.FileServer(http.FS(ui.Files))
	mux.Handle("/static/*filepath", fileServer)

	//mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	infoLog.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	errorLog.Fatal(err)
}
