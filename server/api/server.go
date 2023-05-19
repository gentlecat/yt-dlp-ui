package api

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	requestHandlers "go.roman.zone/yt-dlp-ui/server/api/handlers"
)

func CreateServer() *http.Server {
	router := makeRouter()
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	return &http.Server{
		Handler: loggedRouter,
		Addr:    "0.0.0.0:8080",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}

func makeRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", requestHandlers.CreationInterfaceHandler).Methods(http.MethodGet)
	r.HandleFunc("/", requestHandlers.CreationRequestHandler).Methods(http.MethodPost)

	// Static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("frontend/static"))))

	return r
}
