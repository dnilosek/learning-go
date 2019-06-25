package main

import "fmt"

func main() {
	foo()
	bar("bar")

	s1 := woo("bond")
	fmt.Println(s1)

	x, y := mouse("Ian", "Flemming")
	fmt.Println(x)
	fmt.Println(y)
}

//function def
// func (r receiver) identifier(parameters) (return(s)) { ... }
func foo() {
	fmt.Println("Foo you")
}

// Everything is pass by value
func bar(s string) {
	fmt.Println("Hello", s)
}

func woo(st string) string {
	return fmt.Sprint("Hello from woo, ", st)
}

func mouse(fn string, ln string) (string, bool) {
	s := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := false
	return s, b
}
