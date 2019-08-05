package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

// Countdown begins at 3, waits one second between printing next
// decrement and exiting upon hitting 0 and printing 'GO!'
func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintf(out, "%v\n", i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprint(out, finalWord)

}

func main() {
	Countdown(os.Stdout)
}
