package handler

import (
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

	res := &AsciiResponse{}

	err = tmpl.Execute(w, res)
}

func (app *Application) HandleAscii(w http.ResponseWriter, r *http.Request) {

	// Parse the HTML form data
	err := r.ParseForm()

	// Check for errors
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Get the "text" and "banner" values
	text := r.FormValue("text")
	banner := r.FormValue("banner")

	// A pointer to Config struct
	c := &Config{
		Text:   text,
		Banner: banner,
	}

	// Read the banner file
	bannerLines, errB := ReadBanner("./" + c.Banner + ".txt")
	if errB != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	// Initialize a pointer to the response struct
	res := &AsciiResponse{}

	res.Result = GenerateASCII(c, bannerLines)

	if res.Result != "" {
		app.HandleHomepage(w, r)
		w.WriteHeader(http.StatusOK)

	}

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
