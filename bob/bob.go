package bob // package name must match the package name in bob_test.go
import "strings"

const testVersion = 2

// Hey will return an answer to preceding conversation
func Hey(converse string) string {
	cleanConverse := strings.TrimSpace(converse)
	if cleanConverse == "" {
		return "Fine. Be that way!"
	}
	if strings.ToUpper(cleanConverse) == cleanConverse && strings.ToLower(cleanConverse) != cleanConverse {
		return "Whoa, chill out!"
	}
	if strings.HasSuffix(cleanConverse, "?") {
		return "Sure."
	}
	return "Whatever."
}
