package main

import "fmt"

type hotdog int

func main() {
	a := 42
	var b hotdog = 43
	fmt.Printf("a is %T with val %v\n", a, a)
	fmt.Printf("b is %T with val %v\n", b, b)

	// Won't comple because hotdog is a different type than int
	// a = b

	// Assign via conversion
	a = int(b)
	fmt.Printf("a is %T with val %v\n", a, a)
}
