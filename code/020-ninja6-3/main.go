package main

import "fmt"

func main() {

	func() {
		fmt.Println("Anon func")
	}()
	f := foo

	ff := f()
	ff()
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", ff)

	funcfoo(ff)

	incrementer := increment()
	i := 0
	for i < 10 {
		i = incrementer()
		fmt.Println(i)
	}
}

func foo() func() {
	return func() {
		fmt.Println("funky")
	}
}

func funcfoo(x func()) {
	x()
}

func increment() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
