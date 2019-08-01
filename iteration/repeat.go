package iteration

// Repeater accepts character in form of string, repeats character
// five times.
func Repeater(char string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated = repeated + char
	}
	return repeated
}
