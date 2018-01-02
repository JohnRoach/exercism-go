package reverse

//String reverses a string
func String(word string) string {
	var r = []rune(word)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
