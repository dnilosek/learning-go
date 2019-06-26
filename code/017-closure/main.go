package main

import "fmt"

func main() {
	inc := incrementor()
	inc2 := incrementor()
	for i := 0; i < 10; i++ {
		fmt.Println(inc())
	}
	fmt.Println(inc2())
}

func incrementor() func() int {
	// stays in scope of returned function
	var x int
	return func() int {
		x++
		return x
	}
}
