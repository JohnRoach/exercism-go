package diffsquares

//SquareOfSums returns square of sums
func SquareOfSums(number int) int {
	var sum int
	for index := 1; index <= number; index++ {
		sum += index
	}
	return sum * sum
}

//SumOfSquares returns sum of squares
func SumOfSquares(number int) int {
	var sum int
	for index := 1; index <= number; index++ {
		sum += index * index
	}
	return sum
}

//Difference between square of sum and sum of squares
func Difference(number int) int {
	var squareSum int
	var sum int
	for index := 1; index <= number; index++ {
		squareSum += index * index
		sum += index
	}
	return (sum * sum) - squareSum
}
