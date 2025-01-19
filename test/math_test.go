package test

import (
	"github.com/temirov/utils/math"
	"testing"
)

// TestFormatNumber is a table-driven test for the FormatNumber function.
func TestFormatNumber(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected string
	}{
		{"Whole number", 4.0, "4"},
		{"Whole number with trailing zeros", 5.000, "5"},
		{"Decimal number", 4.5, "4.5"},
		{"Decimal number with trailing zeros", 4.500, "4.5"},
		{"Decimal number with multiple decimal places", 4.657, "4.657"},
		{"Negative whole number", -3.0, "-3"},
		{"Negative decimal number", -3.14, "-3.14"},
		{"Zero", 0.0, "0"},
		{"Negative zero", -0.0, "0"}, // Go treats -0.0 as 0.0 in string formatting
		{"Large whole number", 123456789.0, "123456789"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := math.FormatNumber(tt.input)
			if result != tt.expected {
				t.Errorf("FormatNumber(%v) = %v; expected %v", tt.input, result, tt.expected)
			}
		})
	}
}
