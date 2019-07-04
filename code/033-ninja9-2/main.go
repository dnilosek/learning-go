package main

import "fmt"

type person struct {
	first string
	last  string
}

func (p *person) speak() {
	fmt.Printf("Hi, my name is %v %v\n", p.first, p.last)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
	}
	p1.speak() // autocompiles to (&p1).speak()
	(&p1).speak()

	//saySomething(p1) -- person doesnt implement human interface
	saySomething(&p1) // *person does
}
