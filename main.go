package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port not found!")
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    	AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    	ExposedHeaders: []string{"Link"},
    	AllowCredentials: false,
		MaxAge: 300, // Maximum value not ignored by any of major browsers
	}))

	srv := &http.Server{
		Handler: router,
		Addr: ":" + port,
	}

	v1Router := chi.NewRouter()
	v1Router.Get("/cool", jsonHandler)
	v1Router.Get("/error", errorHandler)

	router.Mount("/v1", v1Router)
	
	log.Printf("Server running on port %v", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port is", port)
}