package iteration

// Repeater accepts character in form of string and a number, repeats character
// the number of times chosen.
func Repeater(char string, num int) string {
	var repeated string
	for i := 0; i < num; i++ {
		repeated += char
	}
	return repeated
}
