package handler

import (
	"embed"
	"html/template"
	"net/http"
	"strings"
)

//go:embed *.txt
var BannerFiles embed.FS

type Config struct {
	Text       string `json:"text"`
	BannerName string `json:"banner"`
}

type AsciiPageData struct {
	Result string
}

type Application struct {
	Template *template.Template
}

func (app *Application) HandleHomepage(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Execute template and handle error if any
	err := app.Template.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (app *Application) HandleAscii(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

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
		Text:       text,
		BannerName: banner,
	}

	// Read the banner file
	bannerLines, errB := ReadBanner(c.BannerName)
	if errB != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	// Initialize a pointer to the response struct
	res := &AsciiPageData{}

	// Generate the ascii art
	res.Result = GenerateASCII(c, bannerLines)

	if res.Result != "" {
		err := app.Template.Execute(w, res)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}

func ReadBanner(filename string) ([]string, error) {
	data, err := BannerFiles.ReadFile(filename + ".txt")
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
