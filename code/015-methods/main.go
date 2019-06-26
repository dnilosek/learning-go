package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// func (r reciever) identifier(parameters) (return(s)) { <code> }

func (s secretAgent) speak() {
	fmt.Println("I am", s.last, ",", s.first, s.last)
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last)
}

// Anyone who is attached to a method speak() is also a human
type human interface {
	speak()
}

func bar(h human) {
	fmt.Println("I was passed into bar", h)
}

// Value can be of more than one type
func main() {
	p1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	p2 := secretAgent{
		person: person{
			first: "Miss",
			last:  "Moneypenny",
		},
		ltk: false,
	}
	p3 := person{
		first: "Dr.",
		last:  "Yes",
	}

	p1.speak()
	p2.speak()

	bar(p1)
	bar(p2)
	bar(p3)

	// Anonymous func
	anon := func() {
		fmt.Println("Anon func")
	}
	anon()

	func() {
		fmt.Println("Anon func")
	}()
}
