package handler_test

import (
	"reflect"
	"testing"

	"github.com/osugodbless/ascii-art-web/handler"
)

func TestReadBanner(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected []string
	}
	tests := []testCase{
		{
			name:     "Single string",
			input:    "../testdata/test.txt",
			expected: []string{" _    _  ", "| |  | | ", "| |__| | ", "|  __  | ", "| |  | | ", "|_|  |_| ", "         "},
		},
		{
			name:     "Wrong file name",
			input:    "../testdata/tes.txt",
			expected: nil,
		},
	}

	for _, tt := range tests {
		result, _ := handler.ReadBanner(tt.input)

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
	result, err := handler.ReadBanner("../standard.txt")

	if err != nil {
		t.Errorf("ReadFile failed: %v", err)
	}

	tests := []testCase{
		{
			name: "test for string with standard banner file",
			input: Parameters{
				config:      &handler.Config{Text: "Hello", Banner: "standard"},
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
				config:      &handler.Config{Text: "", Banner: "standard"},
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
