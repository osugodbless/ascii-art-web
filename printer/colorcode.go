package printer

func GetColorCode(input string) string {
	colorMap := map[string]string{
		"red":     "\x1b[31m",
		"green":   "\033[32m",
		"blue":    "\033[34m",
		"yellow":  "\033[33m",
		"magenta": "\033[35m",
		"cyan":    "\033[36m",
		"orange":  "\033[38;2;255;165;0m",
	}
	return colorMap[input]
}
