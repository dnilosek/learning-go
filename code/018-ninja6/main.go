package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("My name is %v %v and I am %v years old\n", p.first, p.last, p.age)
}

func main() {
	defer fmt.Println(foo())
	fmt.Println(bar())

	x := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(x...))
	fmt.Println(sumSlice(x))

	p := person{
		first: "Dave",
		last:  "Nilosek",
		age:   32,
	}
	p.speak()
}

func foo() int {
	return 42
}
func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
func sumSlice(x []int) int {
	return sum(x...)
}

func bar() (int, string) {
	return 42, "bar"
}
