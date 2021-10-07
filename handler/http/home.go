package handler_http

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	respondwithJSON(w, http.StatusOK, map[string]string{"message": "Hello world!"})
}
