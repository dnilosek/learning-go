package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	// send
	go send(c)

	// receieve
	receive(c)

	fmt.Println("about to exit")
}

// send
func send(c chan<- int) {
	c <- 42
}

// receive
func receive(c <-chan int) {
	fmt.Println(<-c)
}
