package main

import (
	"asciiartweb/handler"
	"fmt"
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

	// Initialize application struct
	app := &handler.Application{}

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

	// if len(os.Args) < 2 {
	// 	fmt.Println("Without a flag, this program needs only one argument.\nYou can specify a custom banner.txt file to use for printing. However this is optional as the program uses the standard banner file as the default banner for printing. See 'Usage' below.\n\nUsage: go run . [STRING] [BANNER]\n\nEX: go run . \"something\" <bannerfilename>.txt")
	// 	os.Exit(1)
	// }

	// // Parse arguments into a struct
	// config, err := printer.ParseArgs(os.Args[1:])
	// if err != nil {
	// 	printer.ColorError()
	// }

	// // Read the banner file
	// bannerLines, err := printer.ReadBanner(config.Banner)
	// if err != nil {
	// 	fmt.Printf("Error: unable to read the banner file '%s'\n", config.Banner)
	// 	os.Exit(1)
	// }

	// // Generate and print the output
	// result := printer.GenerateASCII(config, bannerLines)

	// if config.OutputFilename != "" {
	// 	err := os.WriteFile(config.OutputFilename, []byte(result), 0644)
	// 	if err != nil {
	// 		fmt.Printf("Error writing file to %q\n", config.OutputFilename)
	// 	}
	// } else {
	// 	fmt.Print(result)
	// }
}
