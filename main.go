package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/miku272/RSS_Aggregator/internal/database"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	myPort := os.Getenv("PORT")
	if myPort == "" {
		log.Fatal("Port is not set")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not set")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	myRouter := chi.NewRouter()

	myRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/error", handlerError)
	v1Router.Get("/users", apiCfg.handlerGetUser)

	v1Router.Post("/users", apiCfg.handlerCreateUser)

	myRouter.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: myRouter,
		Addr:    ":" + myPort,
	}

	log.Printf("Starting server on port %s", myPort)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
