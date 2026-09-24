package main

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"jackmitchellfordyce.com/internal/blog"
	"jackmitchellfordyce.com/ui"
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

func (app *application) blog(w http.ResponseWriter, r *http.Request) {

	ts, err := template.ParseFS(ui.Files, "html/pages/blog.tmpl")
	if err != nil {
		app.errorLog.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data := templateData{
		Posts: blog.GetAllPosts(),
	}

	err = ts.Execute(w, data)
	if err != nil {
		app.errorLog.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (app *application) til(w http.ResponseWriter, r *http.Request) {
	posts, err := app.tils.get(r.Context())
	if err != nil && len(posts) == 0 {
		app.errorLog.Print(err)
		http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}
	if err != nil {
		app.errorLog.Print(err)
	}

	ts, err := template.ParseFS(ui.Files, "html/pages/til.tmpl")
	if err != nil {
		app.errorLog.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if err := ts.Execute(w, templateData{TILs: posts}); err != nil {
		app.errorLog.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (app *application) blogPost(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	slug := params.ByName("slug")

	post, found := blog.GetPostBySlug(slug)
	if !found {
		http.NotFound(w, r)
		return
	}

	ts, err := template.New("blog-post.tmpl").Funcs(template.FuncMap{
		"safeHTML": func(s string) template.HTML {
			return template.HTML(s)
		},
	}).ParseFS(ui.Files, "html/pages/blog-post.tmpl")
	if err != nil {
		app.errorLog.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data := templateData{
		Post: post,
	}

	err = ts.Execute(w, data)
	if err != nil {
		app.errorLog.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
