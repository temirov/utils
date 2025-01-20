package pointers

func FromFloat(f float64) *float64 {
	return &f
}

func fromString(s string) *string {
	return &s
}

func fromInteger(i int) *int {
	return &i
}

func fromBoolean(b bool) *bool {
	return &b
}
