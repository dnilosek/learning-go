package main

import "fmt"

func main() {

	c := make(chan int)

	// Send
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c) // range loop in other routine pulls until channel is closed
	}()

	// receieve
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}
