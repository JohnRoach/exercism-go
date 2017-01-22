package raindrops

import (
	"bytes"
	"strconv"
)

const testVersion = 2

// Convert converts numbers to a string
func Convert(convertMe int) string {
	var buffer bytes.Buffer

	if (convertMe % 3) == 0 {
		buffer.WriteString("Pling")
	}

	if (convertMe % 5) == 0 {
		buffer.WriteString("Plang")
	}

	if (convertMe % 7) == 0 {
		buffer.WriteString("Plong")
	}

	if buffer.Len() == 0 {
		buffer.WriteString(strconv.Itoa(convertMe))
	}

	return buffer.String()
}
