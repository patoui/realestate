package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/patoui/realestate/db"
	"github.com/patoui/realestate/models"
)

var listingIDKey = "listingID"

func listings(router chi.Router) {
	router.Get("/", getAllListings)
	router.Post("/", createListing)

	router.Route("/{listingId}", func(router chi.Router) {
		router.Use(ListingContext)
		router.Get("/", getListing)
		router.Put("/", updateListing)
		router.Delete("/", deleteListing)
	})
}

func ListingContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		listingId := chi.URLParam(r, "listingId")
		if listingId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("listing ID is required")))
			return
		}
		id, err := strconv.Atoi(listingId)
		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid listing ID")))
		}
		ctx := context.WithValue(r.Context(), listingIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllListings(w http.ResponseWriter, r *http.Request) {
	listings, err := dbInstance.GetAllListings()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, listings); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

func createListing(w http.ResponseWriter, r *http.Request) {
	listing := &models.Listing{}
	if err := render.Bind(r, listing); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	if err := dbInstance.AddListing(listing); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, listing); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func getListing(w http.ResponseWriter, r *http.Request) {
	listingID := r.Context().Value(listingIDKey).(int)
	listing, err := dbInstance.GetListingById(listingID)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &listing); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func deleteListing(w http.ResponseWriter, r *http.Request) {
	listingId := r.Context().Value(listingIDKey).(int)
	err := dbInstance.DeleteListing(listingId)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
}

func updateListing(w http.ResponseWriter, r *http.Request) {
	listingId := r.Context().Value(listingIDKey).(int)
	listingData := models.Listing{}
	if err := render.Bind(r, &listingData); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	listing, err := dbInstance.UpdateListing(listingId, listingData)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &listing); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
