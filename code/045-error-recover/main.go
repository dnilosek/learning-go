package main

import "fmt"

func main() {
	foo()
	fmt.Println("Exit normal")
}

func foo() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in foo", r)
		}
	}()

	fmt.Println("Calling counter")
	counter(0)
	fmt.Println("Return normally from foo")

}

func counter(i int) {
	if i > 3 {
		fmt.Println("Panic!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in counter", i)
	fmt.Println("Printing in counter", i)
	counter(i + 1)
}
