package api

import (
	"log"
	"net/http"
)

func Serve(port string) {
	http.HandleFunc("POST /exchange-rate", requestHandler)
	log.Printf("Starting on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
