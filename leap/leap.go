// Package leap is expected
package leap

// testVersion should match the targetTestVersion in the test file.
const testVersion = 2

// IsLeapYear checks if the year is leap or not.
// It returns true for leap years
func IsLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	}
	return false
}
