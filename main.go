package main

import (
	"os"
	"log"
	"./src"
	"net/http"
	"github.com/gorilla/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == ""{
		log.Fatal("PORT has not been set")
	}

	router := src.NewRouter()
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	log.Fatal(http.ListenAndServe(":" + port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
