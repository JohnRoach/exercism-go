// Package clock is the expected package
package clock

import (
	"fmt"
	"math/big"
)

// The test version
const testVersion = 4

// Clock struct is defined with an hour and minute
type Clock struct {
	hour, minute int
}

// New returns a newClock
func New(hour int, minute int) Clock {
	var hours, minutes int
	var c Clock
	hours, minutes = ConvertMinutes(minute)
	c.hour = ConvertHours(hour + hours)
	c.minute = minutes
	return c
}

// String returns a string of the Clock
func (c Clock) String() string {
	var hour, hours, minutes int

	hours, minutes = ConvertMinutes(c.minute)
	hour = ConvertHours(c.hour + hours)

	return fmt.Sprintf("%0.2d:%0.2d", hour, minutes)
}

// ConvertHours converts any hour over 24 or below 0 to a standard hour
func ConvertHours(hours int) int {
	return ((hours % 24) + 24) % 24
}

// ConvertMinutes returns hour and minutes for a given total number of minutes
func ConvertMinutes(minutes int) (hours, minute int) {
	var bigHours, bigMinutes *big.Int
	bigHours, bigMinutes = new(big.Int).DivMod(big.NewInt(int64(minutes)),
		big.NewInt(60), big.NewInt(60))
	return int(bigHours.Int64()), int(bigMinutes.Int64())
}

// Add returns a clock with added minutes
func (c Clock) Add(minutes int) Clock {
	var addHour, addMinute int

	addHour, addMinute = ConvertMinutes(minutes + c.minute)

	c.hour = ConvertHours(c.hour + addHour)
	c.minute = addMinute

	return c
}
