package math

import (
	"crypto/rand"
	"fmt"
	"math"
	"strings"
)

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func FormatNumber(num float64) string {
	if num == math.Trunc(num) {
		return fmt.Sprintf("%.0f", num) // Whole number: no decimal places
	}
	// Convert to string with a large precision to avoid scientific notation
	str := fmt.Sprintf("%.15f", num)
	str = strings.TrimRight(str, "0") // Remove trailing zeros
	str = strings.TrimRight(str, ".") // Remove trailing dot if no decimals left
	return str
}

// ChanceOf returns true with the given probability (0.0 to 1.0)
func ChanceOf(probability float64) bool {
	if probability <= 0 {
		return false
	}
	if probability >= 1 {
		return true
	}

	// Generate random byte (0-255)
	randomBytes := make([]byte, 1)
	// Fill the buffer with random bytes
	_, err := rand.Read(randomBytes)
	if err != nil {
		fmt.Println("Error reading random bytes:", err)
		return false
	}

	// Convert probability to 0-255 range
	threshold := uint8(probability * 255)

	return randomBytes[0] < threshold
}
