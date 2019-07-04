package main

import "fmt"

func main() {

	c := make(chan int, 1)
	c_sendonly := make(chan<- int, 1)
	c_recieveonly := make(<-chan int, 1)

	c <- 42
	c_sendonly <- 42
	// c_recieveonly <- 42 -- cant do this because recieve only

	fmt.Println(<-c)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", c_sendonly)
	fmt.Printf("%T\n", c_recieveonly)

	//fmt.Println(<-c_sendonly) -- cant do this because send only

	//Cannot assign specific direction to bidirection
	// var c_wontwork chan int = c_sendonly

	// Can assign bidirection to directional
	var c_assign_s chan<- int = c
	var c_assign_r <-chan int = c

	fmt.Printf("%T\n", c_assign_s)
	fmt.Printf("%T\n", c_assign_r)
}
