package main

import (
	"fmt"
	"runtime"
	"sync"
)

/// FOR DEMO PURPOSES
const maxProcs = 1

func init() {
	runtime.GOMAXPROCS(maxProcs)
}

///

var wg sync.WaitGroup

func main() {
	fmt.Println("CPUs\t\t", maxProcs)
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	fmt.Printf("\n")

	wg.Add(1) // wait for one thing
	go foo()
	bar()

	fmt.Printf("\n")
	fmt.Println("CPUs\t\t", maxProcs)
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Wait()

	fmt.Printf("\n")
	fmt.Println("CPUs\t\t", maxProcs)
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
