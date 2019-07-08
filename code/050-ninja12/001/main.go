package main

import (
	"fmt"
	"github.eagleview.com/david.nilosek/learning-go/code/050-ninja12/001/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}
