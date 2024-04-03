package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/sidsharma96/go-aggregator/handlers"
	"github.com/sidsharma96/go-aggregator/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not found in the environment!")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("Database URL is not found in the environment!")
	}

	conn, dbErr := sql.Open("postgres", dbURL)
	if dbErr != nil {
		log.Fatal("Cannot connect to database - ", dbErr)
	}

	apiCfg := handlers.ApiConfig {
		DB: database.New(conn), 
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

	v1Router.Get("/cool", handlers.JsonHandler)
	v1Router.Get("/error", handlers.ErrorHandler)
	
	v1Router.Get("/users", apiCfg.MiddlewareAuth(apiCfg.HandleAuthorizedUser))
	v1Router.Get("/users/all", apiCfg.HandleGetUsers)
	v1Router.Post("/users", apiCfg.HandleCreateUser)

	v1Router.Get("/feeds", apiCfg.HandlerGetAllFeeds)
	v1Router.Post("/feeds", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeed))

	v1Router.Get("/feed_follows", apiCfg.MiddlewareAuth(apiCfg.HandlerGetFeedFollows))
	v1Router.Post("/feed_follows", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeedFollows))
	v1Router.Delete("/feed_follows/{feedFollowID}", apiCfg.MiddlewareAuth(apiCfg.HandlerDeleteFeedFollows))
	

	router.Mount("/v1", v1Router)
	
	log.Printf("Server running on port %v", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port is", port)
}