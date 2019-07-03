package main

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	age   int
}

func (p person) String() string { // format for printing
	return fmt.Sprintf("%s: %d", p.first, p.age)
}

// Custom sort interface requires Len(), Swap(), Less()
type ByAge []person

func (p ByAge) Len() int           { return len(p) }
func (p ByAge) Less(i, j int) bool { return p[i].age < p[j].age }
func (p ByAge) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	fmt.Println(xi)
	fmt.Println(xs)

	sort.Ints(xi)
	fmt.Println(xi)

	sort.Strings(xs)
	fmt.Println(xs)

	p1 := person{"James", 32}
	p2 := person{"Moneypenny", 27}
	p3 := person{"Q", 64}
	p4 := person{"M", 56}

	people := []person{p1, p2, p3, p4}
	fmt.Println(people)

	sort.Sort(ByAge(people))
	fmt.Println(people)

}
