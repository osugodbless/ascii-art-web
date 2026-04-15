package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type Config struct {
	Text   string `json:"text"`
	Banner string `json:"banner"`
}

type AsciiResponse struct {
	Result string
}

type Application struct {
}

func (app *Application) HandleHomepage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, tmpl)
}

func (app *Application) HandleAscii(w http.ResponseWriter, r *http.Request) {
	// A pointer to Config struct
	c := &Config{}

	// Decode incoming JSON from the request body into Config struct
	err := json.NewDecoder(r.Body).Decode(c)

	// Check for bad JSON error
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Read the banner file
	bannerLines, err := ReadBanner("./" + c.Banner + ".txt")
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	// Initialize a pointer to the response struct
	res := &AsciiResponse{}

	res.Result = GenerateASCII(c, bannerLines)

	app.HandleHomepage(w, r)

}

func ReadBanner(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	contentStr := strings.ReplaceAll(string(data), "\r", "")

	// Split the file content by newline characters
	content := strings.Split(contentStr, "\n")
	return content, nil // return file content
}

// GenerateASCII handles standard and colored printing (using ANSI color codes).
func GenerateASCII(config *Config, bannerLines []string) string {
	result := ""
	words := strings.Split(config.Text, "\\n")

	for _, word := range words {
		if word == "" {
			result += "\n"
			continue
		}

		// Print the 8 lines for each character
		for i := 0; i < 8; i++ {
			for _, char := range word {
				// skip non-printable ASCII characters
				if char < 32 || char > 126 {
					continue
				}

				// Find the exact line in the banner file to print
				lineIndex := i + (int(char-32) * 9) + 1

				result += bannerLines[lineIndex]

			}
			// Add a newline after each ASCII art row
			result += "\n"
		}
	}

	return result
}
