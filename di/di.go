package di

import (
	"bytes"
	"fmt"
)

// Greet takes a name, string, and returns a greeting
func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
