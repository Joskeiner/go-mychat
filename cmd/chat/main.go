package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/joskeiner/go-myChat/internal/config"
	"github.com/joskeiner/go-myChat/internal/server"
)

func main() {
	run()
}

// this function will execute the server and the dependecies
func run() {

	router := http.NewServeMux()
	adr, basePath := config.LoandingDeps()

	server := server.NewServer(adr, router)

	tmpl := template.Must(template.ParseFiles(basePath + "/internal/static/views/index.html"))
	// server static files
	fs := http.FileServer(http.Dir(basePath + "/internal/static/views"))
	router.Handle("/files/", http.StripPrefix("/files/", fs))

	// router test
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, " test to my server")
	})
	// router to static file
	router.HandleFunc("/static", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/html; charset=utf-8")
		tmpl.Execute(w, nil)
	})

	server.Start()
}
