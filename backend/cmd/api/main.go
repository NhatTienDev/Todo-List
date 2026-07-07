package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"
	
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	todoService "github.com/nhattiendev/todo-list/internal/todo/service"
	todoHandler "github.com/nhattiendev/todo-list/internal/todo/handler"
	todoRepository "github.com/nhattiendev/todo-list/internal/todo/repository"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, reading from system env")
	}

	bePort := os.Getenv("BE_PORT")
	dbURL := os.Getenv("DB_URL")
	if bePort == "" || dbURL == "" {
		log.Fatal("Error: Missing required environment variables (BE_PORT, DB_URL)")
	}

	// Initialize PostgreSQL connection
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error: Failed to connect to DB configuration: %v", err)
	}
	
	if err = db.Ping(); err != nil {
		log.Fatalf("Error: Failed to ping DB, please check password and DB status: %v", err)
	}
	log.Println("Successfully connected to PostgreSQL")
	defer db.Close()

	tRepository := todoRepository.NewTodoRepository(db)
	tService := todoService.NewTodoService(tRepository)
	tHandler := todoHandler.NewTodoHandler(tService)

	// General router configuration
	r := chi.NewRouter()

	// CORS middleware configuration
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Health check endpoint
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	
	tHandler.RegisterTodoRoutes(r)

	server := &http.Server{
		Addr:    ":" + bePort,
		Handler: r,
	}

	log.Printf("Starting server on port %s...", bePort)

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Error: Failed to start server: %v", err)
	}
}