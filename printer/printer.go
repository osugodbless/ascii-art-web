package printer

import (
	"fmt"
	"os"
	"strings"
)

// ReadBanner reads the banner file and returns it as a slice of lines
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
func GenerateASCII(config Config, bannerLines []string) string {
	result := ""
	words := strings.Split(config.Text, "\\n")
	colorReset := "\033[0m"

	for _, word := range words {
		if word == "" {
			result += "\n"
			continue
		}

		// map containing all index positions to be colored (if any)
		stringIndexToColor := TargetIndices(word, config.Substring)

		// Print the 8 lines for each character
		for i := 0; i < 8; i++ {
			for j, char := range word {
				// skip non-printable ASCII characters
				if char < 32 || char > 126 {
					continue
				}

				// Find the exact line in the banner file to print
				lineIndex := i + (int(char-32) * 9) + 1

				// check if struct contains a color code
				if config.ColorCode != "" {
					if config.Substring == "" {
						result += config.ColorCode + bannerLines[lineIndex] + colorReset // Color the whole word
						continue
					}
					if stringIndexToColor[j] {
						result += config.ColorCode + bannerLines[lineIndex] + colorReset // Color just the matching substring letters
						continue
					}

					result += bannerLines[lineIndex]

				} else {
					result += bannerLines[lineIndex]
				}

			}
			// Add a newline after each ASCII art row
			result += "\n"
		}
	}

	return result
}

func TargetIndices(s string, substr string) map[int]bool {
	stringIndexToColor := make(map[int]bool)
	delStringLength := 0

	if substr == "" {
		return stringIndexToColor
	}

	for {
		idx := strings.Index(s, substr)

		if idx == -1 {
			break
		}
		for i := 0; i < len(substr); i++ {
			stringIndexToColor[delStringLength+idx+i] = true
		}
		end := idx + len(substr)
		s = s[end:]
		delStringLength += end
	}

	return stringIndexToColor
}

func ColorError() {
	fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> \"something\"")
	os.Exit(0)
}

func OutputError() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<filename.txt> \"something\" standard")
	os.Exit(0)
}
