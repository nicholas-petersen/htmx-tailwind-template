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
	mux.HandleFunc("/winelist", winelist)

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
		"./web/template/partials/table.html",
		"./web/template/views/home.html",
	}

	data := struct {
		Header []string
		Rows   []wine
	}{
		Header: []string{"Country", "Name", "Price", "Year"},
		Rows:   fetchWines(),
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func winelist(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Header []string
		Rows   []wine
	}{
		Header: []string{"Country", "Name", "Price", "Year"},
		Rows:   fetchWines(),
	}

	tmpl, _ := template.ParseFiles("./web/template/partials/table.html")

	tmpl.Execute(w, data)
}
