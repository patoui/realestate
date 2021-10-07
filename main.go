package main

import (
    handler "github.com/patoui/realestate/handler/http"
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World!"))
    })
    r.Get("/home", handler.Home)
    http.ListenAndServe(":3000", r)
}
