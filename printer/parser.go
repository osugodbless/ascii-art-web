package printer

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Config holds all the correctly formatted command-line arguments
type Config struct {
	OutputFilename string
	ColorCode      string
	Substring      string
	Text           string
	Banner         string
}

// ParseArgs reads the raw arguments and uses the data passed to build a struct.
// If it correctly builds the struct without error, it returns the struct.
func ParseArgs(args []string) (Config, error) {
	var config Config
	config.Banner = "standard.txt" // Set default banner

	var positional []string

	for _, arg := range args {

		if strings.HasPrefix(arg, "-color") || strings.HasPrefix(arg, "-colour") || arg == "--color" {
			ColorError()
		}

		if strings.HasPrefix(arg, "-output") || arg == "--output" {
			OutputError()
		}

		if strings.HasPrefix(arg, "--output=") {
			filename := strings.TrimPrefix(arg, "--output=")

			if !strings.HasSuffix(filename, ".txt") {
				OutputError()
			}

			config.OutputFilename = filename
		} else if strings.HasPrefix(arg, "--color=") {
			colorName := strings.TrimPrefix(arg, "--color=")
			config.ColorCode = GetColorCode(strings.ToLower(colorName))

			if config.ColorCode == "" {
				fmt.Println("Color not supported")
				os.Exit(1)
			}
		} else {
			positional = append(positional, arg)
		}
	}

	// Route the remaining positional arguments safely
	if len(positional) == 0 {
		return config, errors.New("Invalid Input, need at least one argument to turn into ASCII")
	}

	if len(positional) == 1 {
		config.Text = positional[0]
	}

	if len(positional) > 2 && config.ColorCode == "" && config.OutputFilename == "" {
		fmt.Println("Without a flag, this program needs only one argument.\nYou can specify a custom banner.txt file to use for printing. However this is optional as the program uses the standard banner file as the default banner for printing. See 'Usage' below.\n\nUsage: go run . [STRING] [BANNER]\n\nEX: go run . \"something\" <bannerfilename>.txt")
		os.Exit(1)
	}

	if len(positional) > 2 && (config.OutputFilename != "" && config.ColorCode == "") {
		OutputError()
	}

	if len(positional) == 2 && config.ColorCode == "" && (config.OutputFilename == "" || config.OutputFilename != "") {
		config.Text = positional[0]
		config.Banner = positional[1] + ".txt"
	}

	if len(positional) == 2 && config.ColorCode != "" && config.OutputFilename != "" {
		config.Text = positional[0]
		config.Banner = positional[1] + ".txt"
	}

	if len(positional) == 2 && config.ColorCode != "" && config.OutputFilename == "" {
		config.Substring = positional[0]
		config.Text = positional[1]
	}

	if len(positional) == 3 && config.ColorCode != "" {
		config.Substring = positional[0]
		config.Text = positional[1]
		config.Banner = positional[2] + ".txt"
	}

	if len(positional) > 3 && (config.ColorCode != "" || config.OutputFilename != "") {
		fmt.Println("Too many arguments")
		os.Exit(1)
	}

	return config, nil
}
