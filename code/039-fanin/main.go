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

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanIn := make(chan int)

	go send(even, odd)

	go receive(even, odd, fanIn)

	for v := range fanIn {
		fmt.Println("Pulling off of fanin", v)
	}

	fmt.Println("exiting...")
}

func send(even, odd chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
			fmt.Println("Adding even")
		} else {
			odd <- i
			fmt.Println("Adding odd")
		}
	}
	close(even)
	close(odd)
}

func receive(even, odd <-chan int, fanIn chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range even {
			fanIn <- v
			fmt.Println("Pulling from even to fanIn")
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanIn <- v
			fmt.Println("Pulling from odd to fanIn")
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanIn)
}
