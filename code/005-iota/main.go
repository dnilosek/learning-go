package main

import "fmt"

const (
	_     = iota
	kb_bs = 1 << (iota * 10)
	mb_bs = 1 << (iota * 10)
	gb_bs = 1 << (iota * 10)
)

func main() {
	x := 2
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

	fmt.Printf("%d\t\t\t%b\n", kb_bs, kb_bs)
	fmt.Printf("%d\t\t\t%b\n", mb_bs, mb_bs)
	fmt.Printf("%d\t\t%b\n", gb_bs, gb_bs)

}
