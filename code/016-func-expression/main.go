package main

import "fmt"

func main() {

	// Func is first class
	f := func() {
		fmt.Println("func expression")
	}
	f()

	s1 := foo()
	fmt.Println(s1())

	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(is_even, ii...)
	fmt.Println(s)
	s = sum(is_odd, ii...)
	fmt.Println(s)
	s = sum(func(x int) bool { return true }, ii...)
	fmt.Println(s)
}

func foo() func() int {
	return func() int {
		return 451
	}
}

func sum(f func(x int) bool, x ...int) int {
	total := 0
	for _, v := range x {
		if f(v) {
			total += v
		}
	}
	return total
}

func is_even(x int) bool {
	return x%2 == 0
}

func is_odd(x int) bool {
	return x%2 != 0
}
