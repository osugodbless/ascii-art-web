package handler_test

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/osugodbless/ascii-art-web/handler"
)

func TestHandleHomepage(t *testing.T) {
	type testCase struct {
		name       string
		method     string
		statusCode int
	}

	tests := []testCase{
		{name: "Valid GET request", method: http.MethodGet, statusCode: http.StatusOK},
		{name: "Invalid POST request", method: http.MethodPost, statusCode: http.StatusMethodNotAllowed},
	}

	tmpl := template.Must(template.ParseFiles("../testdata/home.html"))

	for _, tt := range tests {
		t.Run(tt.name, func(*testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(tt.method, "/home", nil)

			app := &handler.Application{
				Template: tmpl,
			}

			app.HandleHomepage(w, r)

			if w.Code != tt.statusCode {
				fmt.Println(tt.name)
				t.Errorf("Handler returned wrong status code: %v expected --> %v", w.Code, tt.statusCode)
			}
		})

	}
}

func TestHandleAscii(t *testing.T) {
	type testCase struct {
		name       string
		method     string
		reqBody    io.Reader
		statusCode int
	}

	tests := []testCase{
		{name: "Invalid GET request", method: http.MethodGet, statusCode: http.StatusMethodNotAllowed, reqBody: nil},
		{name: "POST request without form data", method: http.MethodPost, statusCode: http.StatusNotFound, reqBody: nil},
		{name: "Valid POST request", method: http.MethodPost, statusCode: http.StatusOK, reqBody: strings.NewReader("text=H&banner=standard")},
	}

	tmpl := template.Must(template.ParseFiles("../testdata/home.html"))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(tt.method, "/ascii-art", tt.reqBody)

			r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			app := &handler.Application{
				Template: tmpl,
			}

			app.HandleAscii(w, r)

			if w.Code != tt.statusCode {
				fmt.Println(tt.name)
				t.Errorf("Handler returned wrong status code: %v expected --> %v", w.Code, tt.statusCode)
			}
		})
	}
}
func TestReadBanner(t *testing.T) {
	type testCase struct {
		name     string
		filename string
		expected []string
	}

	tests := []testCase{
		{
			name:     "Single string",
			filename: "test",
			expected: []string{" _    _  ", "| |  | | ", "| |__| | ", "|  __  | ", "| |  | | ", "|_|  |_| ", "         "},
		},
		{
			name:     "Wrong file name",
			filename: "tes",
			expected: nil,
		},
	}

	for _, tt := range tests {
		result, _ := handler.ReadBanner(tt.filename)

		if !reflect.DeepEqual(result, tt.expected) {
			t.Error("Content does not match with expected")
		}
	}

}

func TestGenerateASCII(t *testing.T) {
	type Parameters struct {
		config      *handler.Config
		bannerLines []string
	}
	type testCase struct {
		name     string
		input    Parameters
		expected string
	}
	result, err := handler.ReadBanner("standard")

	if err != nil {
		t.Errorf("ReadFile failed: %v", err)
	}

	tests := []testCase{
		{
			name: "test for string with standard banner file",
			input: Parameters{
				config:      &handler.Config{Text: "Hello"},
				bannerLines: result,
			},
			expected: ` _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
`,
		},
		{
			name: "test for empty string",
			input: Parameters{
				config:      &handler.Config{Text: ""},
				bannerLines: result,
			},
			expected: `
`,
		},
	}

	for _, tt := range tests {
		s := handler.GenerateASCII(tt.input.config, tt.input.bannerLines)
		if s != tt.expected {
			t.Errorf("GenerateASCII function returned\n%q\nexpected -->\n%q", s, tt.expected)
		}
	}
}
