package iteration

// Repeater accepts character in form of string, repeats character
// five times.
func Repeater(char string, num int) string {
	var repeated string
	for i := 0; i < num; i++ {
		repeated += char
	}
	return repeated
}
