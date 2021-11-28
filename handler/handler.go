package handler

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/patoui/realestate/db"
)

var dbInstance db.Database

func NewHandler(db db.Database) http.Handler {
	router := chi.NewRouter()
	dbInstance = db

	router.MethodNotAllowed(methodNotAllowedHandler)
	router.NotFound(notFoundHandler)

	// API routes
	router.Group(func(r chi.Router) {
		router.Route("/api/items", items)
		router.Route("/api/listings", listings)
	})

	router.Route("/", home)

	// Serve static files/assets
	fileServer := http.FileServer(http.Dir("./static/client/"))
	router.Handle("/static/*", http.StripPrefix("/static/client", fileServer))

	return router
}

func clientHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(405)
	render.Render(w, r, ErrMethodNotAllowed)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	contentType := w.Header().Get("Accepts")

	if strings.Contains(contentType, "json") == true {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		render.Render(w, r, ErrNotFound)
	} else {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(404)
		http.ServeFile(w, r, "static/index.html")
	}
}
