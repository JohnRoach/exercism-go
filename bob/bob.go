package bob // package name must match the package name in bob_test.go
import (
	"regexp"
	"strings"
)

const testVersion = 2

// Hey will return an answer to preceding conversation
func Hey(converse string) string {
	cleanConverse := strings.Trim(converse, " \n\r\t\v\u00a0\u2002")

	maybeYell, _ := regexp.MatchString(`\b[A-Z][^a-z].*\b`, cleanConverse)

	if cleanConverse == "" {
		return "Fine. Be that way!"
	} else if (cleanConverse == strings.ToUpper(cleanConverse)) && maybeYell {
		return "Whoa, chill out!"
	} else if isQuestion, _ := regexp.MatchString(`\?$`, cleanConverse); isQuestion {
		// the trick above made this to be a little faster 429954 ns/op vs 385752 ns/op
		return "Sure."
	}
	return "Whatever."
}
