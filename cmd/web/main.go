package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	port := 3001
	staticFiles := http.FileServer(http.Dir("./web/static"))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", staticFiles))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/wine", wine)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	log.Printf("starting server on :%d", port)
	err := srv.ListenAndServe()
	log.Fatalf("Failed to ListenAndServe: %d", err)
}

func home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./web/template/base.html",
		"./web/template/partials/nav.html",
		"./web/template/views/home.html",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	err = tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func wine(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./web/template/base.html",
		"./web/template/partials/nav.html",
		"./web/template/views/wine.html",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	err = tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
