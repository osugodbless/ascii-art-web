package printer_test

import (
	"asciiartweb/printer"
	"reflect"
	"testing"
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
		result, _ := printer.ReadBanner(tt.input)

		if !reflect.DeepEqual(result, tt.expected) {
			t.Error("Content does not match with expected")
		}
	}

}

func TestGenerateASCII(t *testing.T) {
	type Parameters struct {
		config      printer.Config
		bannerLines []string
	}
	type testCase struct {
		name     string
		input    Parameters
		expected string
	}
	result, err := printer.ReadBanner("../standard.txt")

	if err != nil {
		t.Errorf("ReadFile failed: %v", err)
	}

	tests := []testCase{
		{
			name: "test for substr in string",
			input: Parameters{
				config:      printer.Config{ColorCode: "\x1b[31m", Substring: "B", Text: "Blue", Banner: "standard"},
				bannerLines: result,
			},
			expected: `[31m ____   [0m _                 
[31m|  _ \  [0m| |                
[31m| |_) | [0m| |  _   _    ___  
[31m|  _ <  [0m| | | | | |  / _ \ 
[31m| |_) | [0m| | | |_| | |  __/ 
[31m|____/  [0m|_|  \__,_|  \___| 
[31m        [0m                   
[31m        [0m                   
`,
		},
		{
			name: "test for empty substr",
			input: Parameters{
				config:      printer.Config{ColorCode: "\x1b[31m", Substring: "", Text: "Blue", Banner: "standard"},
				bannerLines: result,
			},
			expected: `[31m ____   [0m[31m _  [0m[31m        [0m[31m       [0m
[31m|  _ \  [0m[31m| | [0m[31m        [0m[31m       [0m
[31m| |_) | [0m[31m| | [0m[31m _   _  [0m[31m  ___  [0m
[31m|  _ <  [0m[31m| | [0m[31m| | | | [0m[31m / _ \ [0m
[31m| |_) | [0m[31m| | [0m[31m| |_| | [0m[31m|  __/ [0m
[31m|____/  [0m[31m|_| [0m[31m \__,_| [0m[31m \___| [0m
[31m        [0m[31m    [0m[31m        [0m[31m       [0m
[31m        [0m[31m    [0m[31m        [0m[31m       [0m
`,
		},
		{
			name: "test for empty slice",
			input: Parameters{
				config:      printer.Config{ColorCode: "\x1b[31m", Substring: "B", Text: "", Banner: "standard"},
				bannerLines: result,
			},
			expected: `
`,
		},
	}

	for _, tt := range tests {
		s := printer.GenerateASCII(tt.input.config, tt.input.bannerLines)
		if s != tt.expected {
			t.Errorf("GenerateASCII function returned\n%q\nexpected -->\n%q", s, tt.expected)
		}
	}
}

func TestTargetIndices(t *testing.T) {
	type Parameters struct {
		str    string
		substr string
	}
	type testCase struct {
		name     string
		input    Parameters
		expected map[int]bool
	}

	tests := []testCase{
		{
			name: "Check that the correct indices are return when substr is present",
			input: Parameters{
				str:    "hello",
				substr: "ll",
			},
			expected: map[int]bool{
				2: true,
				3: true,
			},
		},
		{
			name: "Check that an empty map is return when substr is absent",
			input: Parameters{
				str:    "hello",
				substr: "low",
			},
			expected: map[int]bool{},
		},
		{
			name: "Check that an empty substring",
			input: Parameters{
				str:    "hello",
				substr: "",
			},
			expected: map[int]bool{},
		},
	}

	for _, tt := range tests {
		s := printer.TargetIndices(tt.input.str, tt.input.substr)
		if !reflect.DeepEqual(s, tt.expected) {
			t.Errorf("TargetIndices = %v\nExpected --> %v\n", s, tt.expected)
		}
	}
}
