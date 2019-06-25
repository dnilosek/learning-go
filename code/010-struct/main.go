package main

import "fmt"

// Composite/Aggregated/Complex data structure
type person struct {
	first string
	last  string
}

// Embedded (anonymous) field in struct
type secretAgent struct {
	person // Anonymous (unqulified type name) acts as field name
	ltk    bool
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
	}
	sa1 := secretAgent{
		person: p1,
		ltk:    true,
	}
	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
	}

	fmt.Println(p1)
	fmt.Println(sa1)
	fmt.Println(p2)

	fmt.Println(sa1.first)
	fmt.Println(sa1.person.first) // Dont need to namespace unless for name collision

	// Anonymous struct
	a1 := struct {
		person
		happy bool
	}{
		person: p1,
		happy:  false,
	}

	fmt.Println(a1)
}
