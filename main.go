package main

import (
	"log"

	"go.roman.zone/yt-dlp-ui/server/api"
)

func main() {
	log.Println("Starting server on http://localhost:8080...")
	err := api.CreateServer().ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
