package hamming

import "fmt"

const testVersion = 5

// Distance returns the hamming distance between two strings
func Distance(a, b string) (int, error) {
	aLength := len(a)
	bLength := len(b)

	if aLength != bLength {
		return -1, fmt.Errorf("Strings %s and %s are not of equal length", a, b)
	}

	distance := 0

	for index := range a {
		if a[index] != b[index] {
			distance = distance + 1
		}
	}

	return distance, nil

}
