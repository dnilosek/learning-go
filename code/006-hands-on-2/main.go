package main

import "fmt"

func PrintIt(num int) {
	fmt.Printf("%d\n", num)
	fmt.Printf("%o\n", num)
	fmt.Printf("%b\n", num)
}

const (
	this_year         = 2019 + iota
	next_year         = 2019 + iota
	nextnext_year     = 2019 + iota
	nextnextnext_year = 2019 + iota
)

func main() {
	num := 42
	PrintIt(num)

	num = num << 1
	PrintIt(num)

	raw_string := `I'm a "raw" string literal`
	fmt.Println(raw_string)

	fmt.Printf("The next four years are: %d , %d, %d , %d", this_year, next_year, nextnext_year, nextnextnext_year)
}
