package utils

import "strings"

// Normalize whitespace in Markdown output to avoid formatting mismatches
func Normalize(input string) string {
	lines := strings.Split(input, "\n")
	var cleaned []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" { // Skip empty lines
			cleaned = append(cleaned, trimmed)
		}
	}
	return strings.Join(cleaned, "\n")
}
