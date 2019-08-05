package main

import (
	"fmt"
	"io"
	"os"
)

// Countdown begins at 3, waits one second between printing next
// decrement and exiting upon hitting 0 and printing 'GO!'
func Countdown(out io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintf(out, "%v\n", i)
	}
	fmt.Fprint(out, "Go!")

}

func main() {
	Countdown(os.Stdout)
}
