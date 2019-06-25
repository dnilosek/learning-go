package main

import "fmt"

func main() {
	foo(1, 2, 3, 4, 5, 6, 7, 8, 9)

	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	foo(xi...) // variadic parameter expansion
	foo()      // variadic can also be empty
}

/// func (r reciever) identifier(parameter(s)) return(s) {code}
func foo(x ...int) { // Variadic parameter has to be last parameter
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for _, val := range x {
		sum += val
	}
	fmt.Println(sum)
}
