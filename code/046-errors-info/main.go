package main

import (
	"errors"
	"fmt"
)

var DivideZero = errors.New("Invalid operation: divide by zero")

func main() {
	_, err := divide(4, 2)
	if err != nil {
		fmt.Println(err)
	}

	_, err = divide(4, 0)
	if err != nil {
		fmt.Println(err)
	}
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, DivideZero
	}
	return a / b, nil
}
