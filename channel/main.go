package main

import (
	"fmt"
	"runtime"
)

// passing chan to function
func SayHelloTo(name string, messages chan string) {
	messages <- fmt.Sprintf("Hello, %s", name)
}

func PrintMessage(message chan string) {
	fmt.Println(<-message)
}

func main() {
	runtime.GOMAXPROCS(2)

	messages := make(chan string)

	names := []string{"Benzion", "Benjamin", "Deborah"}

	for _, name := range names {
		go SayHelloTo(name, messages)
	}

	for i := 0; i < len(names); i++ {
		PrintMessage(messages)
	}
}
