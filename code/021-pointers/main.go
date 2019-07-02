package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func changeMe(p *person) {
	(*p).firstName = "Miss"   // one way to dereference struct << probably more clear
	p.lastName = "Moneypenny" // this way works too if not a method
}

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	p := person{
		firstName: "James",
		lastName:  "Bond",
	}
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}
