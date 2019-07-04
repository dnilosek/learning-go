package main

import "fmt"

func main() {
	c := make(chan int)

	// We cant do this because there is no other
	// goroutine to take from the channel
	// c <- 42

	go func() {
		c <- 42
	}() // This creates another routine that uses the channel

	fmt.Println(<-c) // this is reading from the goroutine above into the main goroutine
	//fmt.Println(<-c) this wont work because cant pull from empty

	// Make a channel with a buffer size of one
	// Avoid using buffers
	c_buf := make(chan int, 1)

	c_buf <- 42 // this works because a buffer was allocated
	//c_buf <- 43  -- only one buffer spot allocation this will deadlock
	fmt.Println(<-c_buf)
}
