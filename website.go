package main

import (
	"html/template"
	"net/http"
	"fmt"
)

func fail(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func redirectTo(w http.ResponseWriter, url string, r *http.Request) {
	http.Redirect(w, r, "www.google.de", http.StatusFound)
}

func renderPage(w http.ResponseWriter, templateName string, data interface{}, r *http.Request) {
	if err := templates[templateName].ExecuteTemplate(w, "base", data); err != nil {
		fail(w, err)
	}
}

var fixedTemplates = template.Must(template.ParseGlob("templates/fixed/*"))
var templates = make(map[string]*template.Template)

func main() {
	addPageTemplates("index")
	addPageTemplates("projects")

	http.Handle("/res/", http.StripPrefix("/res/", http.FileServer(http.Dir("res"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { renderPage(w, "index", &Page{}, r) })
	http.HandleFunc("/projects/", func(w http.ResponseWriter, r *http.Request) { renderPage(w, "projects", nil, r) })
	http.HandleFunc("/photography/", func(w http.ResponseWriter, r *http.Request) { renderPage(w, "projects", nil, r) })
	http.ListenAndServe(":8080", nil)
}

func addPageTemplates(name string) {
	if element, err := fixedTemplates.Clone(); err != nil {
		fmt.Println(err.Error())
	} else {
		element.ParseFiles("templates/"+name+".html")
		templates[name] = element
	}
}