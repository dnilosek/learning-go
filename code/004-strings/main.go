package main

import "fmt"

func main() {
	// String is a slice of bytes
	str := "Hello"
	fmt.Println(str)
	fmt.Printf("%T\n", str)

	// Convert string into slice of bytes
	bStr := []byte(str)
	fmt.Println(bStr)
	fmt.Printf("%T\n", bStr)

	for i := 0; i < len(str); i++ {
		fmt.Printf("%#U \n", str[i])
	}

	for i, v := range str {
		fmt.Printf("%v %#U \n", i, v)
	}
}
