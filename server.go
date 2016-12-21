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
	server := http.Server{
		Addr: "0.0.0.0:8090",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
