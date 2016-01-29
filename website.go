package main

import (
	"fmt"
	"net/http"
)

func AddHtml(s string) string {
	return fmt.Sprintf("<html><body>%s</body></html>", s)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	p := &Page{Title: "Title", Body: []byte("Hi there, I'm trying stuff! Check out my projects! <a href=\"projects/\">Projects</a>")}
	fmt.Fprintf(w, AddHtml(string(p.Body)))
	
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, AddHtml("Hey, I'm showing projects!"))
}

func main() {
    http.HandleFunc("/", indexHandler)
	 http.HandleFunc("/projects/", projectsHandler)
	 http.HandleFunc("/photography/", projectsHandler)
    http.ListenAndServe(":8080", nil)
}