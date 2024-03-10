package api

import (
	"log"
	"net/http"
)

func Serve(port string) {
	http.HandleFunc("POST /exchange-rate", requestHandler)
	log.Fatal(http.ListenAndServe(port, nil))
}
