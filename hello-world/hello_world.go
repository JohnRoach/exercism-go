package greeting

const testVersion = 3

// HelloWorld needs a comment documenting it as package does.
func HelloWorld(what string) string {
	var resultant string
	if what != "" {
		resultant = what
	} else {
		resultant = "World"
	}
	return "Hello, " + resultant + "!"
}
