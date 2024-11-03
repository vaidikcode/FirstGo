package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	//reading the env file and loading it
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//it gets the value of environment variable from env file its more like key value thing
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("$PORT must be set")
	}
	//start a router
	router := chi.NewRouter()
	//config for router
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	//giving the router handler to server
	srv := &http.Server{Handler: router, Addr: ":" + portString}
	//Serving the server
	err2 := srv.ListenAndServe()
	if err2 != nil {
		log.Fatal(err2)
	}
}
