package main

import (
	"embed"
	"io/fs"
	"log"

	"go.roman.zone/yt-dlp-ui/server/api"
)

//go:embed frontend/templates/*.html frontend/static/*
var frontendContent embed.FS

func main() {
	actualFrontendContent, _ := fs.Sub(frontendContent, "frontend")

	log.Println("Starting server on http://localhost:8080...")
	err := api.CreateServer(actualFrontendContent).ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
