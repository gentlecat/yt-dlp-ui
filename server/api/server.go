package api

import (
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	requestHandlers "go.roman.zone/yt-dlp-ui/server/api/handlers"
)

func CreateServer(staticContent fs.FS) *http.Server {
	r := mux.NewRouter().StrictSlash(true)

	templates, err := template.ParseFS(staticContent, "templates/*.html")
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}

	r.HandleFunc("/", requestHandlers.NewCreationInterfaceHandler(templates).Handle).Methods(http.MethodGet)
	r.HandleFunc("/", requestHandlers.CreationRequestHandler).Methods(http.MethodPost)

	r.PathPrefix("/static/").Handler(http.FileServer(http.FS(staticContent)))

	return &http.Server{
		Handler: handlers.LoggingHandler(os.Stdout, r),
		Addr:    "0.0.0.0:8080",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
