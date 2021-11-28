package handler

import (
	"encoding/json"
	"net/http"
)

// respondWithHTML write json response format
func respondWithHTML(w http.ResponseWriter, code int, content []byte) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(code)
	w.Write(content)
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
