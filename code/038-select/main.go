package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// Send
	go send(eve, odd, quit)

	receieve(eve, odd, quit)
}
func receieve(eve, odd, quit <-chan int) {
	for {
		select {
		// Comma okay idiom to check if recieve successful
		case v, ok := <-eve:
			fmt.Println("From even channel:", v, ok)
		case v, ok := <-odd:
			fmt.Println("From odd channel:", v, ok)
		case v, ok := <-quit:
			if !ok {
				fmt.Println("From the quit channel", v, ok)
				return
			} else {
				fmt.Println("From the quit channel", v)
			}
		}
	}
}
func send(eve, odd, quit chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			eve <- i
		} else {
			odd <- i
		}
	}
	quit <- 0
	close(quit)
}
