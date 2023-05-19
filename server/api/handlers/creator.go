package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"

	"go.roman.zone/yt-dlp-ui/server/api/handlers/templates"
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

func CreationInterfaceHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplates(w, struct{}{}, "frontend/templates/index.html")
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
}

func download(url *url.URL, path string) {
	cmd := exec.Command(
		"yt-dlp",
		"-P", path,
		url.String(),
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Println("Error: ", err)
	}
}
