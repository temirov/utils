// Package pointers provides helper functions to obtain pointers to basic
// primitive values. These utilities simplify working with optional or pointer
// based APIs.
package pointers

// FromFloat returns a pointer to the provided float64 value.
func FromFloat(f float64) *float64 {
	return &f
}

// fromString returns a pointer to the provided string. It is unexported as it
// is primarily used internally for testing.
func fromString(s string) *string {
	return &s
}

// fromInteger returns a pointer to the provided int. It is kept unexported
// because only specific packages require it.
func fromInteger(i int) *int {
	return &i
}

// fromBoolean returns a pointer to the provided bool value.
func fromBoolean(b bool) *bool {
	return &b
}
