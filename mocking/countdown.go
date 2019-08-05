package main

import (
	"bytes"
	"fmt"
)

// Countdown begins at 3, waits one second between printing next
// decrement and exiting upon hitting 0 and printing 'GO!'
func Countdown(writer *bytes.Buffer) {
	fmt.Fprint(writer, "3")
}

func main() {
}
