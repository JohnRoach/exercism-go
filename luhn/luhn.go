package luhn

import (
	"math"
	"strconv"
	"strings"
)

//Valid checks if number is valid per luhns algorithm
func Valid(numbers string) bool {
	var isValid bool
	var total float64
	var totalSecondDigit float64
	var numberPlacement int
	numbers = strings.Replace(numbers, " ", "", -1)
	if len(numbers) > 1 {
		for index := len(numbers) - 1; index >= 0; index-- {
			number, error := strconv.Atoi(string(numbers[index]))
			if error == nil {
				if numberPlacement%2 != 0 {
					if number*2 > 9 {
						totalSecondDigit += float64(number)*2 - 9
					} else {
						totalSecondDigit += float64(number) * 2
					}
				} else {
					total += float64(number)
				}
			} else {
				return isValid
			}
			numberPlacement++
		}
		total += totalSecondDigit
		if math.Mod(total, 10) == 0 {
			isValid = true
		}
	}
	return isValid
}
