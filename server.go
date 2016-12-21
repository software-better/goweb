package main // import "github.com/software-better/goweb"

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	t.Execute(w, "Hello World!")
}

func main() {
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
