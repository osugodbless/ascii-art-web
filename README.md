<div align="center">

![Logo](/assets/asii_logo.png)

# ASCII ART WEB

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=for-the-badge)](https://github.com/RichardLitt/standard-readme)
[![Go version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go)](https://golang.org/dl/)
[![MIT License](https://img.shields.io/badge/License-MIT-brightgreen?style=for-the-badge)](./LICENSE)
[![GitHub](https://img.shields.io/badge/GitHub-181717?style=for-the-badge&logo=github&logoColor=white)](https://github.com/osugodbless)

A web-based application written in Go that takes a string as input and displays it in a graphic representation using ASCII characters. This project transitions the classic command-line ASCII art generator I previously built (see it [here](https://github.com/osugodbless/ascii-art)) into a fully functional web server with a graphical user interface and server-side rendering.
</div>

## Table of Contents

- [ASCII ART WEB](#ascii-art-web)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Requirements](#requirements)
  - [Usage: How to Run](#usage-how-to-run)
    - [Local Development](#local-development)
  - [Examples](#examples)
  - [Implementation Details: Algorithm](#implementation-details-algorithm)
  - [Author(s)](#authors)
  - [License](#license)


## Features
- **Web Interface:** Easy-to-use HTML form for inputting text and selecting styles.
- **Multiple Banner Styles:** Supports `standard`, `shadow`, and `thinkertoy` ASCII fonts.
- **Extensive Character Support:** Handles uppercase, lowercase, numbers, spaces, and special characters.
- **Multiline Support:** Accurately processes `\n` literal inputs to create stacked, multiline ASCII art.
- **Graceful Error Handling:** Robust server-side validation and proper HTTP status code management (e.g., 400 Bad Request, 500 Internal Server Error).

## Requirements
- Go 1.22+ (Required for `GET /path` routing syntax)
- `github.com/joho/godotenv` package
- 
- An `.env` file specifying server configuration.

## Usage: How to Run

### Local Development
1. Clone the repository and navigate to the project root.
2. Create an `.env` file in the root directory:
   ```sh
   HOST=<your-host-here>   // e.g localhost or 127.0.0.1
   PORT=<your-port-here>   // e.g 8080
   ```
3. Install dependencies:
   ```sh
   go mod tidy
   ```
4. Start the server:
   ```sh
   go run .
   ```
5. Open your web browser and navigate to `http://<your-host>:<your-port>/home` e.g., `http://127.0.0.1:8080/home`

## Examples
Once the server is running, interact with the application via the web interface at `/home`:

- **Single Word**: Type any word e.g., Hello, select "Standard" radio button, and click the `Generate ASCII Art` button.

- **Newlines**: Type `Hello\nWorld` to render the words on separate stacked lines.

- **One or Multiple blank lines**: Type `Hello\n\nWorld` to render the words with a blank line between them.

## Implementation Details: Algorithm
The core logic for rendering the ASCII art is located in handler/handler.go within the GenerateASCII function. The algorithm operates on a horizontal, row-by-row string concatenation strategy rather than a character-by-character approach. Below is the step-by-step approach:

1. **Input Processing:** The raw input string is split into an array of words using the `\n `delimiter. Empty strings denote intentional blank lines.
2. **Line-by-Line Rendering:** Because each ASCII character is exactly 8 lines tall in the banner file, the algorithm utilizes an outer loop that runs exactly 8 times.
3. **Character Mapping:** For every iteration of the 8-line loop, an inner loop traverses each character of the current word.
   - Non-printable characters (ASCII decimal values < 32 or > 126) are ignored.
   - The exact starting line of a character in the text file is calculated using its decimal ASCII value and a specific offset formula: 
  
        `i + (int(char-32) * 9) + 1` 
  
        (Where i is the current row 0-7, 32 is the offset for the first printable character (Space), and 9 represents the 8 lines of character height plus 1 blank separation line).
1. **String Building:** The specific horizontal slice for each character is appended to the result string. Once all characters in the word have had their $i^{th}$ line appended, a newline character (`\n`) is added, and the loop moves down to render the next horizontal slice.

## Author(s)
[@osugodbless](https://github.com/osugodbless)

## License
MIT © 2026 [Osu Godbless](https://github.com/osugodbless)