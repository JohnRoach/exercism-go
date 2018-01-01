package isogram

import (
	"regexp"
	"strings"
)

//IsIsogram will check if string is an isogram or not
func IsIsogram(word string) bool {
	cleanWord := strings.ToLower(word)
	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString
	var isIsogram = true
	for index, character := range cleanWord {
		for _, subCharacter := range cleanWord[index+1 : len(cleanWord)] {
			if subCharacter == character && isAlpha(string(character)) {
				isIsogram = false
				break
			}
		}
	}
	return isIsogram
}
