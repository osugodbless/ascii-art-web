package main

import (
	"asciiartweb/handler"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get server configuration
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	// Parse template. If successful, returns template, otherwise shut down server initialization
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	// Inject parsed template into Application struct
	app := &handler.Application{
		Template: tmpl,
	}

	// Initialize custom multiplexer and register handlers to their corresponding pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("GET /home", app.HandleHomepage)
	mux.HandleFunc("POST /ascii-art", app.HandleAscii)

	// Set up custom server configurations
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", host, port),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      mux,
	}

	// Start server
	fmt.Printf("Server is running on %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
