package pangram

import "unicode"

const testVersion = 1

// mappedAlphabet maps the alphabet to false values for default
func mappedAlphabet() map[rune]bool {
	makeMap := make(map[rune]bool, 28)
	for i := 'A'; i <= 'Z'; i++ {
		//Set all mapped values as false first
		makeMap[i] = false
	}
	return makeMap
}

// IsPangram checks if a given sentence is a pangram or not.
func IsPangram(sentence string) bool {
	mappedAlphabet := mappedAlphabet()
	for _, character := range []rune(sentence) {
		if unicode.IsLetter(character) {
			formatedLetter := unicode.ToUpper(character)
			if _, goodLetter := mappedAlphabet[formatedLetter]; goodLetter {
				mappedAlphabet[formatedLetter] = true
			}
		}
	}

	// check if one of the mapped values still doesn't have true
	for _, value := range mappedAlphabet {
		if !value {
			return false
		}
	}
	return true
}
