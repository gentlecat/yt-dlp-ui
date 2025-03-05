package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func CreationRequestHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse submitted form.", http.StatusInternalServerError)
		return
	}

	submittedURL, err := url.ParseRequestURI(r.Form.Get("url"))
	if submittedURL == nil || err != nil {
		http.Error(w, "Invalid URL provided.", http.StatusBadRequest)
		return
	}

	path := r.Form.Get("path")

	go download(submittedURL, path)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Status string
	}{
		Status: "OK",
	})
}
