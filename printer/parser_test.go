package printer_test

import (
	"asciiartweb/printer"
	"reflect"
	"testing"
)

func TestParseArgs(t *testing.T) {

	type testCase struct {
		name     string
		input    []string
		expected printer.Config
	}

	tests := []testCase{
		{
			name:  "Check for the presence of color and output flag",
			input: []string{"--color=red", "--output=banner.txt", "B", "Blue", "standard"},
			expected: printer.Config{
				ColorCode:      "\x1b[31m",
				OutputFilename: "banner.txt",
				Substring:      "B",
				Text:           "Blue",
				Banner:         "standard.txt",
			},
		},
		{
			name:  "Check for when only color flag is present",
			input: []string{"--color=red", "B", "Blue"},
			expected: printer.Config{
				ColorCode:      "\x1b[31m",
				OutputFilename: "",
				Substring:      "B",
				Text:           "Blue",
				Banner:         "standard.txt",
			},
		},
		{
			name:  "No flag present, but banner present",
			input: []string{"Blue", "standard"},
			expected: printer.Config{
				ColorCode:      "",
				OutputFilename: "",
				Substring:      "",
				Text:           "Blue",
				Banner:         "standard.txt",
			},
		},
		{
			name:  "No flag present, no banner present",
			input: []string{"Blue"},
			expected: printer.Config{
				ColorCode:      "",
				OutputFilename: "",
				Substring:      "",
				Text:           "Blue",
				Banner:         "standard.txt",
			},
		},
	}

	for _, tt := range tests {
		s, _ := printer.ParseArgs(tt.input)
		if !reflect.DeepEqual(s, tt.expected) {
			t.Errorf("ParseArgs = %v\nExpected --> %v\n", s, tt.expected)
		}
	}
}
