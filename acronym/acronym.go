package acronym

import (
	"bytes"
	"regexp"
	"strings"
)

const testVersion = 2

// Abbreviate create abbreviations for a given string
func Abbreviate(abbreviateMe string) string {
	var buffer bytes.Buffer

	re := regexp.MustCompile(`\b[A-Z]{1}|\b\w|[A-Z]([a-z])`)
	splitme := re.FindAllString(abbreviateMe, -1)

	for index := range splitme {
		buffer.WriteString(strings.ToUpper(string([]rune(splitme[index])[0])))
	}

	return buffer.String()
}
