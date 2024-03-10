package main

import (
	"exchangerate/api"
	"github.com/joho/godotenv"
	"os"
)

func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		return ":8000"
	}

	return ":" + port
}

func main() {
	godotenv.Load()
	api.Serve(getPort())
}
