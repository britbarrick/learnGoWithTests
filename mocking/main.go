package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

const (
	finalWord      = "Go!"
	countdownStart = 3
)

// Countdown begins at 3, waits one second between printing next
// decrement and exiting upon hitting 0 and printing 'GO!'
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintf(out, "%v\n", i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)

}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
