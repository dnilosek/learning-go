package main

import "fmt"

func main() {
	c := make(chan int)

	go gen(c)

	for i := 0; i < 100; i++ {
		fmt.Println(<-c)
	}
}

func gen(c chan int) {

	const numRoutines = 10
	const numVal = 10

	for i := 0; i < numRoutines; i++ {
		go func() {
			for k := 0; k < numVal; k++ {
				c <- k
			}
		}()
	}
}
