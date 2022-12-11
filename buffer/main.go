package main

import (
	"fmt"
	"time"
)

func SequentialBuffer() {
	c := make(chan int, 3)

	c <- 1
	c <- 2
	c <- 3

	for i := 0; i < cap(c); i++ {
		fmt.Println(<-c)
	}
}

func BlockingBuffer() {
	c := make(chan int, 1)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Printf("Send %d\n", i+1)
			c <- i + 1
		}
	}()

	for i := 0; i < 6; i++ {
		fmt.Printf("Receives %d\n", <-c)
	}
}

func RangeBuffer() {
	c := make(chan int, 1)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Printf("Send %d\n", i+1)
			c <- i + 1
		}
		close(c)
	}()

	for item := range c {
		fmt.Printf("Receives %d\n", item)
	}
}

func RangeSequential() {
	c := make(chan int, 5)

	for i := 0; i < 5; i++ {
		fmt.Printf("Send %d\n", i+1)
		c <- i + 1
	}

	close(c)

	for item := range c {
		fmt.Printf("Receives %d\n", item)
	}
}

func SelectBuffer() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func() {
		for i := 0; i < 10; i++ {
			c1 <- fmt.Sprintf("Every 500 milliseconds")
			time.Sleep(500 * time.Millisecond)
		}
		close(c1)
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c2 <- fmt.Sprintf("Every 1 second")
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 10; i++ {
		select {
		case msg := <-c1:
			fmt.Println(msg)
		case msg := <-c2:
			fmt.Println(msg)
		}
	}
}

func main() {
	SelectBuffer()
}
