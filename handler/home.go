package handler

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/go-chi/chi/v5"
)

func home(router chi.Router) {
	router.Get("/", index)
}

func index(w http.ResponseWriter, r *http.Request) {
	fp, err := filepath.Abs("templates/index.html")

	if err != nil {
		log.Fatalf("Could not determine file path: %v\n", err)
	}

	t, err := template.ParseFiles(fp)

	if err != nil {
		log.Fatalf("Unable to parse files: %v\n", err)
	}

	var a struct{}

	err = t.Execute(w, a)
	if err != nil {
		log.Fatalf("Unable to execute template: %v\n", err)
	}
}
