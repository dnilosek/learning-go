package main

import "fmt"

// Global var
var global = 42

// Zero value is assigned when not initalized
var globalIntUninit int

var interpString string = "I am an inepreted string"

// Raw string literal will keep everthing
var rawString string = `I am a 


raw string`

func main() {

	//Declare and assign at the same time
	var x = 42
	fmt.Println("I am not short", x)

	// Short Decl. Op (cannot be global)
	// Decalre and assign at the same time
	y := 42
	fmt.Println("I am short", y)

	fmt.Println("I am global", global)

	fmt.Println("I am not initalized", globalIntUninit)
	globalIntUninit = 42
	fmt.Println("I am initalized", globalIntUninit)

	// Get type
	fmt.Printf("%T\n", globalIntUninit)
	fmt.Printf("%v\n", globalIntUninit)
	fmt.Printf("%#v\n", globalIntUninit)
	fmt.Printf("%b\n", globalIntUninit)
	fmt.Printf("%x\n", globalIntUninit)
	fmt.Printf("%#x\n", globalIntUninit)
	fmt.Printf("%#o\n", globalIntUninit)

	var zero_int int
	fmt.Println("Zero int: ", zero_int)

	var zero_str string
	fmt.Println("Zero string: ", zero_str)

	var zero_flt float64
	fmt.Println("Zero float64: ", zero_flt)

	var zero_img complex128
	fmt.Println("Zero complex128: ", zero_img)

	var zero_bool bool
	fmt.Println("Zero bool: ", zero_bool)

	var zero_ptr_int *int
	fmt.Println("Zero pointer: ", zero_ptr_int)
}
