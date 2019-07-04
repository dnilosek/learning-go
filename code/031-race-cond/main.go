package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

/// FOR DEMO PURPOSES
const maxProcs = 8

func init() {
	runtime.GOMAXPROCS(maxProcs)
}

///

func main() {
	counter := 0
	var counter_atom int64
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			// time.Sleep(time.Second * 2)
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()

			atomic.AddInt64(&counter_atom, 1)
			fmt.Println("Counter\t", atomic.LoadInt64(&counter_atom))
			wg.Done()
		}()
		fmt.Println("Num goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Printf("Counter val: %v\n", counter)
}
