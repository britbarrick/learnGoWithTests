package main

import "fmt"

// Hello passes in welcome string.
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
